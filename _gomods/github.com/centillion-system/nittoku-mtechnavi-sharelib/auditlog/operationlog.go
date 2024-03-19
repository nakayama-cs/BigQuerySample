package auditlog

import (
	"context"
	pb "mtechnavi/sharelib/auditlog/protobuf"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryServerOperationLogger(fn func(context.Context) LogParameters) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		res, err := handler(ctx, req)
		params := fn(ctx)
		PrintOperationLog(ctx, params.TenantId, params.UserId, info.FullMethod, req, err)
		return res, err
	}
}

// 操作ログ
func PrintOperationLog(ctx context.Context, tenantId, userId, fullMethodName string, req interface{}, err error) {
	writeOperationLog(ctx, tenantId, userId, "ACCESS_LOG", fullMethodName, req, err)
}

func writeOperationLog(ctx context.Context, tenantId, userId, auditSubcategory, fullMethodName string, req interface{}, err error) {
	// MEMO: grpcステータスを取得する。RESTの場合はステータスコードが`Unknown`となってしなうため、別途でREST用のログ関数を作成するなど検討する
	status, _ := status.FromError(err)
	writeLog(ctx, logRequest{
		EventID:          uuid.New().String(),
		TenantID:         tenantId,
		UserID:           userId,
		AuditCategory:    pb.AuditCategory_OPERATION.String(),
		AuditSubcategory: auditSubcategory,
		Message:          pb.MessageCode_OPERATED,
		Parameters: map[string]interface{}{
			"requestName": fullMethodName,
			"response": map[string]interface{}{
				"code":      status.Code().String(),
				"message":   status.Message(),
				"arguments": req,
			},
		},
	})
}
