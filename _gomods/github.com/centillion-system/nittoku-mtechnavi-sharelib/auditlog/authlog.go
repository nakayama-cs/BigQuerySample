package auditlog

import (
	"context"
	pb "mtechnavi/sharelib/auditlog/protobuf"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func UnaryServerPermissionLogger(fn func(context.Context) LogParameters) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		res, err := handler(ctx, req)
		params := fn(ctx)
		switch status.Code(err) {
		case codes.Unauthenticated:
			PrintPermissionNGLog(ctx, params.TenantId, params.UserId, "") // TODO: 権限等の要件決まり次第修正
		}
		return res, err
	}
}

// 認証・認可ログ
func PrintLoginOKLog(ctx context.Context, tenantId, userId string) {
	writeLoginLog(ctx, tenantId, userId, "LOGIN", pb.MessageCode_AUTH_LOGIN_OK)
}

func PrintLoginNGLog(ctx context.Context, tenantId, userId string) {
	writeLoginLog(ctx, tenantId, userId, "LOGIN", pb.MessageCode_AUTH_LOGIN_NG)
}

func PrintLogoutLog(ctx context.Context, tenantId, userId string) {
	writeLoginLog(ctx, tenantId, userId, "LOGOUT", pb.MessageCode_AUTH_LOGOUT)
}

func writeLoginLog(ctx context.Context, tenantId, userId, auditSubcategory string, message pb.MessageCode) {
	var (
		remoteAddr string
		userAgent  string
	)
	if pr, ok := peer.FromContext(ctx); ok {
		remoteAddr = pr.Addr.String()
	}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ua, ok := md["user-agent"]; ok {
			userAgent = ua[0]
		}
	}
	writeLog(ctx, logRequest{
		EventID:          uuid.New().String(),
		TenantID:         tenantId,
		UserID:           userId,
		AuditCategory:    pb.AuditCategory_AUTH.String(),
		AuditSubcategory: auditSubcategory,
		Message:          message,
		Parameters: map[string]interface{}{
			"remoteAddress": remoteAddr,
			"userAgent":     userAgent,
		},
	})
}

func PrintPermissionNGLog(ctx context.Context, tenantId, userId, requiredPermission string) {
	writePermissionLog(ctx, tenantId, userId, "PERMISSION_DENY", pb.MessageCode_AUTH_PERMISSION_NG, requiredPermission)
}

func writePermissionLog(ctx context.Context, tenantId, userId, auditSubcategory string, message pb.MessageCode, requiredPermission string) {
	writeLog(ctx, logRequest{
		TenantID:         tenantId,
		UserID:           userId,
		AuditCategory:    pb.AuditCategory_AUTH.String(),
		AuditSubcategory: auditSubcategory,
		Message:          message,
		Parameters: map[string]interface{}{
			"requiredPermission": requiredPermission,
		},
	})
}
