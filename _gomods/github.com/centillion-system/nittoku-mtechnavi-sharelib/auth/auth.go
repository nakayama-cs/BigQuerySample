package auth

import (
	"context"
	"crypto/rand"
	"math/big"
	"path"
	"sort"
	"strings"
	"sync"

	idppb "mtechnavi/idp/protobuf"
	"mtechnavi/sharelib/constants"
	"mtechnavi/sharelib/grpc/dialer"
	st "mtechnavi/sharelib/grpc/status"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Middleware struct {
	// 認証/認可システムAPIのURL
	IdentityServiceURL string

	// 認証処理をスキップさせるフルメソッド名
	SkipFullMethodNames []string

	initOnce sync.Once

	initErr error

	identityClient idppb.IdentityClient
}

func (m *Middleware) init(ctx context.Context) error {
	m.initOnce.Do(func() {
		if cc, err := dialer.DialContext(ctx, m.IdentityServiceURL); err != nil {
			m.initErr = err
			return
		} else {
			m.identityClient = idppb.NewIdentityClient(cc)
		}
	})
	return m.initErr
}

func (m *Middleware) shouldSkip(fullMethodName string) bool {
	for _, s := range m.SkipFullMethodNames {
		if s == fullMethodName {
			return true
		}
	}
	return false
}

func (m *Middleware) authenticate(ctx context.Context, fullMethodName string) (context.Context, error) {
	logger := ctxlogrus.Extract(ctx)
	if err := m.init(ctx); err != nil {
		logger.WithError(err).Error("failed to initialize")
		return ctx, status.Error(codes.Internal, "UNKNOWN")
	}

	if m.shouldSkip(fullMethodName) {
		logger.Debugf("skip %q", fullMethodName)
		return ctx, nil
	}

	md, _ := metadata.FromIncomingContext(ctx)
	// VerifyUserに処理を移譲するため、ブラウザから受け取ったヘッダなどを渡す
	// XXX: CloudRunで動作したとき、acquireCredentialsで自動付与されるAuthorizationヘッダを正しく反映するため、Authorizationヘッダを削除しておく
	// CloudRunでは、AuthorizationヘッダはIAM認証を通すために利用される
	// 一度IAM認証を通ると、CloudRunはAuthorizationヘッダ値を改変するため、ここで削除しないと改変されたAuthorizationヘッダが優先されて呼び出しがうまくいかない
	md.Delete("authorization")
	verifyUserCtx := metadata.NewOutgoingContext(ctx, md)

	// Identityサービス側に処理を移譲
	// 速度的な問題は、必要性が出たら考える
	verifyUserRes, err := m.identityClient.VerifyUser(verifyUserCtx, &emptypb.Empty{})
	switch status.Code(err) {
	case codes.OK:
	case codes.Unauthenticated:
		return ctx, status.Error(codes.Unauthenticated, "UNAUTHENTICATED")
	default:
		logger.WithError(err).Error("failed to verify user")
		return ctx, status.Error(codes.Internal, "UNKNOWN")
	}
	ctx = ContextWithInvoker(ctx, &Invoker{
		Tenant: verifyUserRes.Tenant,
		User:   verifyUserRes.User,
		Groups: verifyUserRes.Groups,
		Roles:  verifyUserRes.Roles,
	})

	// 認証が取れたため、後続の処理用に認証に必要なヘッダを付与
	for _, v := range md.Get("cookie") {
		ctx = metadata.AppendToOutgoingContext(ctx, "cookie", v)
	}
	for _, v := range md.Get("x-authorization") {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-authorization", v)
	}
	return ctx, nil
}

func (m *Middleware) authorize(ctx context.Context, fullMethodName string) (context.Context, error) {
	logger := ctxlogrus.Extract(ctx)
	if err := m.init(ctx); err != nil {
		logger.WithError(err).Error("failed to initialize")
		return ctx, status.Error(codes.Internal, "UNKNOWN")
	}

	if m.shouldSkip(fullMethodName) {
		logger.Debugf("skip %q", fullMethodName)
		return ctx, nil
	}
	invoker := InvokerFromContext(ctx)
	// ログインユーザの保有するロールから、パーミッションの一覧を作成
	var permissions []string
	for _, v := range invoker.Roles {
		permissions = append(permissions, v.Permissions...)
	}
	permissions = uniqStrings(permissions)
	// fullMethodNameを "{service}/{method}" という形式で検証
	fullMethodName = strings.TrimPrefix(fullMethodName, "/")
	for _, perm := range permissions {
		// XXX: まだパーミッションの形式が定義されていないため、暫定でGlobマッチ
		ok, err := path.Match(perm, fullMethodName)
		if err != nil {
			logger.WithError(err).Panicf("unable to match glob with pattern: %q", perm)
		}
		if ok {
			logger.Debugf("Authn allow: (fullMethodName, permission) = (%q, %q)", fullMethodName, perm)
			return ctx, nil
		}
	}
	logger.Debugf("Authn deny: (fullMethodName) = (%q)", fullMethodName)
	return ctx, st.Error(ctx, codes.PermissionDenied, constants.MessageName_E0000019)
}

func uniqStrings(l []string) []string {
	memo := make(map[string]struct{}, len(l))
	for _, v := range l {
		memo[v] = struct{}{}
	}
	out := make([]string, 0, len(memo))
	for v := range memo {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}

func (m *Middleware) StreamServerAuthenticator() grpc.StreamServerInterceptor {
	return streamServerInterceptor(m.authenticate)
}

func (m *Middleware) UnaryServerAuthenticator() grpc.UnaryServerInterceptor {
	return unaryServerInterceptor(m.authenticate)
}

func (m *Middleware) StreamServerAuthorizer() grpc.StreamServerInterceptor {
	return streamServerInterceptor(m.authorize)
}

func (m *Middleware) UnaryServerAuthorizer() grpc.UnaryServerInterceptor {
	return unaryServerInterceptor(m.authorize)
}

type authFunc func(ctx context.Context, fullMethodName string) (context.Context, error)

func unaryServerInterceptor(fn authFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, err := fn(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

func streamServerInterceptor(fn authFunc) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		newCtx, err := fn(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		wrapped := grpcmiddleware.WrapServerStream(stream)
		wrapped.WrappedContext = newCtx
		return handler(srv, wrapped)
	}
}

func ContainsAnyGroupID(userGroups []*idppb.UserGroup, targetUserGroupIDs []string) bool {

	targetUserGroupIDMap := map[string]struct{}{}
	for _, targetUserGroupID := range targetUserGroupIDs {
		targetUserGroupIDMap[targetUserGroupID] = struct{}{}
	}

	for _, userGroup := range userGroups {
		if _, ok := targetUserGroupIDMap[userGroup.UserGroupId]; ok {
			return true
		}
	}
	return false
}

func ContainsAnyGroupCode(userGroups []*idppb.UserGroup, targetUserGroupCodes []string) bool {

	targetUserGroupCodeMap := map[string]struct{}{}
	for _, targetUserGroupCode := range targetUserGroupCodes {
		targetUserGroupCodeMap[targetUserGroupCode] = struct{}{}
	}

	for _, userGroup := range userGroups {
		if _, ok := targetUserGroupCodeMap[userGroup.UserGroupCode]; ok {
			return true
		}
	}
	return false
}

func IsTenantAdmin(userGroups []*idppb.UserGroup) bool {
	for _, g := range userGroups {
		if g.UserGroupCode == "Group-Tenant-Admin" {
			return true
		}
	}
	return false
}

func IsTenantMemberAll(userGroups []*idppb.UserGroup) bool {
	for _, g := range userGroups {
		if g.UserGroupCode == "Group-Tenant-Member-All" {
			return true
		}
	}
	return false
}

func GeneratePassword() string {
	var password []byte
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$*?!@#%&><_"
	PasswordLength := 10
	// 英字、数字から1文字ずつ取得
	letter := getRandomChar("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	number := getRandomChar("0123456789")

	// 英字と数字を1文字ずつ使用パスワードの配列に追加
	password = append(password, letter, number)

	// 残りの8文字をcharsからランダムに選択
	for i := 2; i < PasswordLength; i++ {
		password = append(password, getRandomChar(chars))
	}

	// シャッフルして順番をランダム化
	for i := range password {
		j := randInt(len(password))
		password[i], password[j] = password[j], password[i]
	}

	return string(password)
}

func getRandomChar(s string) byte {
	index := randInt(len(s))
	return s[index]
}

func randInt(n int) int {
	max := big.NewInt(int64(n))
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return int(r.Int64())
}
