package auditlog

import (
	"context"
	"encoding/json"

	pb "mtechnavi/sharelib/auditlog/protobuf"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
)

const (
	fieldEventId          = "eventId"
	fieldTenantId         = "tenantId"
	fieldUserId           = "userId"
	fieldRecordId         = "recordId"
	fieldTable            = "table"
	fieldColumns          = "columns"
	fieldMessage          = "message"
	fieldAuditCategory    = "auditCategory"
	fieldAuditSubCategory = "auditSubCategory"
	fieldParameters       = "parameters"
)

type logRequest struct {
	EventID          string
	TenantID         string
	UserID           string
	RecordID         string
	Table            string
	Columns          string
	AuditCategory    string
	AuditSubcategory string
	Message          pb.MessageCode
	Parameters       interface{}
}

func writeLog(ctx context.Context, req logRequest) {
	if req.AuditCategory == "" || req.AuditSubcategory == "" || req.Message == 0 {
		panic("req.AuditCategory, req.AuditSubcategory, req.Message is mandatory")
	}
	// TODO: Extractできなかったときに`nullLogger`が返されてしまうため、デフォルトのlogger初期化処理を追加する
	logger := ctxlogrus.Extract(ctx)
	for _, pair := range [][]string{
		{fieldTenantId, req.TenantID},
		{fieldUserId, req.UserID},
	} {
		name, value := pair[0], pair[1]
		if value == "" {
			value = "-"
		}
		logger = logger.WithField(name, value)
	}
	for _, pair := range [][]string{
		{fieldEventId, req.EventID},
		{fieldRecordId, req.RecordID},
		{fieldTable, req.Table},
		{fieldColumns, req.Columns},
		{fieldAuditCategory, req.AuditCategory},
		{fieldAuditSubCategory, req.AuditSubcategory},
	} {
		name, value := pair[0], pair[1]
		if value == "" {
			continue
		}
		logger = logger.WithField(name, value)
	}

	if req.Parameters != nil {
		switch v := req.Parameters.(type) {
		case json.RawMessage:
			logger = logger.WithField(fieldParameters, v)
		case map[string]interface{}:
			if b, err := json.Marshal(v); err != nil {
				logrus.WithError(err).Error("writeLog: fail to marshal req.Parameters")
			} else {
				logger = logger.WithField(fieldParameters, string(b))
			}
		default:
			panic("invalid type")
		}

	}
	logger.Info(req.Message)
}
