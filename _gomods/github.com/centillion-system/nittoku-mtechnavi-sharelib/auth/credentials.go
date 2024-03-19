package auth

import (
	"context"
	"encoding/json"
	"errors"
	"mtechnavi/sharelib/cache"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"google.golang.org/api/iamcredentials/v1"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"google.golang.org/grpc/credentials"
)

const (
	// テナント内での特権管理者として利用する、固定のシステムユーザID
	SystemUserID = "4512cee7-5770-0140-b9e4-5b6af7f17cfb"

	// テナント内での特権管理グループとして利用する、固定のシステムユーザグループID
	SystemUserGroupID = "010b1406-948e-0f89-b505-b2a7ee41a618"

	// システムユーザID用のロール名
	SystemUserRoleName = "Role-System-User"

	// 内部トークンとして発行するJWT Claimsにおける、テナントIDキー
	ClaimTenantID = "https://m-technavi.com/claims/tenant-id"

	// JWTのIssuer
	JWTIssuer = "https://m-technavi.com"
)

// サービスごとの設定ではなく、システム基盤側の設定のため、環境変数を直接読む

var (
	// 内部トークンとして発行するJWTを署名するサービスアカウントのメールアドレス
	// XXX: 内部トークンを発行する場合、環境変数による設定が必須
	JWTSignerEmail = os.Getenv("MTECHNAVI_JWT_SIGNER_EMAIL")

	// for testing
	TimeNow = time.Now
)

type internalCredentials struct {
	TenantID string

	UserID string
}

func InternalCredentials(tenantID, userID string) credentials.PerRPCCredentials {
	return &internalCredentials{
		TenantID: tenantID,
		UserID:   userID,
	}
}

func (cred *internalCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	signedJWT, err := SignedJWT(ctx, cred.TenantID, cred.UserID)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"X-Authorization": "Internal " + signedJWT,
	}, nil
}

func (cred *internalCredentials) RequireTransportSecurity() bool {
	return true
}

type insecureCredentials struct {
	credentials.PerRPCCredentials

	Insecure bool
}

// for debugging
func WrapInsecureCredentials(creds credentials.PerRPCCredentials) credentials.PerRPCCredentials {
	return &insecureCredentials{
		PerRPCCredentials: creds,
		Insecure:          true,
	}
}

func (cred *insecureCredentials) RequireTransportSecurity() bool {
	return !cred.Insecure
}

var (
	iamcredentialsServiceMu sync.Mutex

	iamcredentialsService *iamcredentials.Service
)

type cachedSignedJWT struct {
	Token string

	ExpiresAt int64
}

func SignedJWT(ctx context.Context, tenantID, userID string) (string, error) {
	if JWTSignerEmail == "" {
		return "", errors.New("SignedJWT disabled")
	}
	now := TimeNow()
	lru := cache.CacheFromContext(ctx)
	var (
		cacheKey = "mtechnavi/sharelib/auth.cachedSignedJWT:" + tenantID + ":" + userID
		cached   cachedSignedJWT
	)
	if lru.Get(cacheKey, &cached) {
		// 失効前（多少の遊びを持たせて）なら再利用
		if now.Add(30*time.Second).Unix() < cached.ExpiresAt {
			return cached.Token, nil
		}
		lru.Delete(cacheKey)
	}

	// See https://www.rfc-editor.org/rfc/rfc7519.html#section-4.1
	expiresAt := now.Add(time.Hour).Unix()
	token, err := signJWT(ctx, map[string]any{
		"aud": JWTSignerEmail,
		"sub": userID,
		"exp": expiresAt,
		"iss": JWTIssuer,
		"iat": now.Unix(),
		// API実行対象となるテナントID
		ClaimTenantID: tenantID,
	})
	if err != nil {
		return "", err
	}
	lru.Set(cacheKey, &cachedSignedJWT{
		Token:     token,
		ExpiresAt: expiresAt,
	})
	return token, nil
}

func signJWT(ctx context.Context, claims map[string]any) (string, error) {
	payload, err := json.Marshal(claims)
	if err != nil {
		panic(err)
	}

	iamcredentialsServiceMu.Lock()
	defer iamcredentialsServiceMu.Unlock()

	if iamcredentialsService == nil {
		service, err := iamcredentials.NewService(ctx)
		if err != nil {
			return "", err
		}
		iamcredentialsService = service
	}

	signJwtRes, err := iamcredentialsService.Projects.ServiceAccounts.SignJwt(
		// expect "projects/-/serviceAccounts/{{serviceAccountEmail}}"
		path.Join("projects/-/serviceAccounts", JWTSignerEmail),
		&iamcredentials.SignJwtRequest{
			Payload: string(payload),
		},
	).Context(ctx).Do()
	if err != nil {
		return "", err
	}
	return signJwtRes.SignedJwt, nil
}

type serviceAccountJWKTransport struct {
	Base http.RoundTripper

	Email string
}

func (tp *serviceAccountJWKTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	u := *req.URL
	u.RawFragment = ""
	u.RawQuery = ""

	// googleSACertsURL  string = "https://www.googleapis.com/oauth2/v3/certs"
	// swap https://www.googleapis.com/service_accounts/v1/metadata/jwk/{ACCOUNT_EMAIL}
	switch u.String() {
	case "https://www.googleapis.com/oauth2/v3/certs":
		req.URL.Path = path.Join("/service_accounts/v1/metadata/jwk", tp.Email)
	}
	return tp.Base.RoundTrip(req)
}

func VerifySignedJWT(ctx context.Context, token string) (map[string]any, error) {
	if JWTSignerEmail == "" {
		return nil, errors.New("SignedJWT disabled")
	}

	// See https://cloud.google.com/iam/docs/reference/credentials/rest/v1/projects.serviceAccounts/signJwt#response-body
	// https://github.com/googleapis/google-api-go-client/issues/640
	lru := cache.CacheFromContext(ctx)
	hc := lru.NewHTTPClient(ctx, option.WithoutAuthentication())
	hc.Transport = &serviceAccountJWKTransport{
		Base:  hc.Transport,
		Email: JWTSignerEmail,
	}
	validator, err := idtoken.NewValidator(ctx, idtoken.WithHTTPClient(hc))
	if err != nil {
		return nil, err
	}
	tok, err := validator.Validate(ctx, token, JWTSignerEmail)
	if err != nil {
		return nil, err
	}
	if tok.Issuer != JWTIssuer {
		return nil, errors.New("invalid jwt issuer")
	}
	out := map[string]any{
		"iss": tok.Issuer,
		"aud": tok.Audience,
		"exp": tok.Expires,
		"iat": tok.IssuedAt,
		"sub": tok.Subject,
	}
	for k, v := range tok.Claims {
		out[k] = v
	}
	return out, nil
}
