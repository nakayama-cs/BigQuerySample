package auditlog

import (
	"context"
	pb "mtechnavi/sharelib/auditlog/protobuf"

	"github.com/google/uuid"
)

type Email interface {
	GetTemplateId() string
	GetSubject() string
	GetTo() string
	GetCc() string
	GetBcc() string
}

func PrintSendgridRequestOKLog(ctx context.Context, tenantId, userId string, email Email) {
	writeSendgridLog(ctx, tenantId, userId, "SEND_MAIL", pb.MessageCode_SENDGRID_REQUEST_OK, email)
}

func PrintSendgridRequestNGLog(ctx context.Context, tenantId, userId string, email Email) {
	writeSendgridLog(ctx, tenantId, userId, "SEND_MAIL", pb.MessageCode_SENDGRID_REQUEST_NG, email)
}

func writeSendgridLog(ctx context.Context, tenantId, userId, auditSubcategory string, message pb.MessageCode, email Email) {
	writeLog(ctx, logRequest{
		EventID:          uuid.New().String(),
		TenantID:         tenantId,
		UserID:           userId,
		AuditCategory:    pb.AuditCategory_INTERNAL.String(),
		AuditSubcategory: auditSubcategory,
		Message:          message,
		Parameters: map[string]interface{}{
			"templateId": email.GetTemplateId(),
			"subject":    email.GetSubject(),
			"to":         email.GetTo(),
			"cc":         email.GetCc(),
			"bcc":        email.GetBcc(),
		},
	})
}
