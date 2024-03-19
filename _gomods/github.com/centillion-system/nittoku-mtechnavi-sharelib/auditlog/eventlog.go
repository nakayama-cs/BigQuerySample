package auditlog

import (
	"context"
	"encoding/json"
	pb "mtechnavi/sharelib/auditlog/protobuf"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	EventName_TenantCreated  = "TENANT_CREATED"
	EventName_RecordCreated  = "RECORD_CREATED"
	EventName_RecordUpdated  = "RECORD_UPDATED"
	EventName_RecordDeleted  = "RECORD_DELETED"
	EventName_RecordShared   = "RECORD_SHARED"
	EventName_RecordUnshared = "RECORD_UNSHARED"
)

// EventTriggerで受け取ったリクエストから各種パラメータを取得するためのインターフェース
// "mtechnavi/notification/protobuf"のEventメッセージのデータにアクセスするために使用する
// 例）HandleTriggerEvent関数
//
// Sharelibから"mtechnavi/notification/protobuf"への依存が発生するのを避けるため、
// 本インターフェースを経由してアクセスさせる
type EventRequest interface {
	GetEventId() string
	GetEventName() string
	GetTenantId() string
	GetUserId() string
	GetParameters() *anypb.Any
}

func HandleTriggerEvent(ctx context.Context, req EventRequest,
	eventMap map[string]func(context.Context, EventRequest, *pb.GenericEventParameter) error) error {

	logger := ctxlogrus.Extract(ctx)

	fn, ok := eventMap[req.GetEventName()]
	if !ok {
		logger.Warnf("not supported event eventId:%s,eventName:%s", req.GetEventId(), req.GetEventName())
		// 対応していないEventはリトライする必要がないのでエラー無しとして正常終了させる
		return nil
	}

	parameters := &pb.GenericEventParameter{}
	if err := anypb.UnmarshalTo(req.GetParameters(), parameters, proto.UnmarshalOptions{}); err != nil {
		logger.WithError(err).Error("failed to unmarshal parameter")
		// リトライしても復旧できないので、エラーログを残して、エラー無しとして正常終了させる
		return nil
	}

	return fn(ctx, req, parameters)
}

type EventParameterTenantCreated struct {
	Data *pb.GenericEventParameter
}

type EventParameterRecordCreated struct {
	Data *pb.GenericEventParameter
}

type EventParameterRecordUpdated struct {
	Data *pb.GenericEventParameter
}
type EventParameterRecordDeleted struct {
	Data *pb.GenericEventParameter
}

type EventParameterRecordShared struct {
	Data *pb.GenericEventParameter
}

type EventParameterRecordUnshared struct {
	Data *pb.GenericEventParameter
}

func setGenericEventParameter(parameter **pb.GenericEventParameter, key string, val string) {
	if *parameter == nil {
		*parameter = &pb.GenericEventParameter{}
	}
	if (*parameter).Data == nil {
		(*parameter).Data = map[string]string{}
	}

	(*parameter).Data[key] = val
}

func getGenericEventParameter(parameter *pb.GenericEventParameter, key string) string {
	if parameter == nil || parameter.Data == nil {
		return ""
	}

	return parameter.Data[key]
}

func (p *EventParameterTenantCreated) SetCreatedTenantId(val string) {
	setGenericEventParameter(&p.Data, "createdTenantId", val)
}

func (p *EventParameterTenantCreated) GetCreatedTenantId() string {
	return getGenericEventParameter(p.Data, "createdTenantId")
}

func (p *EventParameterRecordCreated) SetTypeName(val string) {
	setGenericEventParameter(&p.Data, "typeName", val)
}

func (p *EventParameterRecordCreated) GetTypeName() string {
	return getGenericEventParameter(p.Data, "typeName")
}

func (p *EventParameterRecordCreated) SetRecordId(val string) {
	setGenericEventParameter(&p.Data, "recordId", val)
}

func (p *EventParameterRecordCreated) GetRecordId() string {
	return getGenericEventParameter(p.Data, "recordId")
}

func (p *EventParameterRecordUpdated) SetTypeName(val string) {
	setGenericEventParameter(&p.Data, "typeName", val)
}

func (p *EventParameterRecordUpdated) GetTypeName() string {
	return getGenericEventParameter(p.Data, "typeName")
}

func (p *EventParameterRecordUpdated) SetRecordId(val string) {
	setGenericEventParameter(&p.Data, "recordId", val)
}

func (p *EventParameterRecordUpdated) GetRecordId() string {
	return getGenericEventParameter(p.Data, "recordId")
}

func (p *EventParameterRecordDeleted) SetTypeName(val string) {
	setGenericEventParameter(&p.Data, "typeName", val)
}

func (p *EventParameterRecordDeleted) GetTypeName() string {
	return getGenericEventParameter(p.Data, "typeName")
}

func (p *EventParameterRecordDeleted) SetRecordId(val string) {
	setGenericEventParameter(&p.Data, "recordId", val)
}

func (p *EventParameterRecordDeleted) GetRecordId() string {
	return getGenericEventParameter(p.Data, "recordId")
}

func (p *EventParameterRecordShared) SetSharedByTenantId(val string) {
	setGenericEventParameter(&p.Data, "sharedByTenantId", val)
}

func (p *EventParameterRecordShared) GetSharedByTenantId() string {
	return getGenericEventParameter(p.Data, "sharedByTenantId")
}

func (p *EventParameterRecordShared) SetSharedToTenantId(val string) {
	setGenericEventParameter(&p.Data, "sharedToTenantId", val)
}

func (p *EventParameterRecordShared) GetSharedToTenantId() string {
	return getGenericEventParameter(p.Data, "sharedToTenantId")
}

func (p *EventParameterRecordShared) SetTypeName(val string) {
	setGenericEventParameter(&p.Data, "typeName", val)
}

func (p *EventParameterRecordShared) GetTypeName() string {
	return getGenericEventParameter(p.Data, "typeName")
}

func (p *EventParameterRecordShared) SetRecordId(val string) {
	setGenericEventParameter(&p.Data, "recordId", val)
}

func (p *EventParameterRecordShared) GetRecordId() string {
	return getGenericEventParameter(p.Data, "recordId")
}

func (p *EventParameterRecordUnshared) SetUnsharedByTenantId(val string) {
	setGenericEventParameter(&p.Data, "unsharedByTenantId", val)
}

func (p *EventParameterRecordUnshared) GetUnsharedByTenantId() string {
	return getGenericEventParameter(p.Data, "unsharedByTenantId")
}

func (p *EventParameterRecordUnshared) SetUnsharedToTenantId(val string) {
	setGenericEventParameter(&p.Data, "unsharedToTenantId", val)
}

func (p *EventParameterRecordUnshared) GetUnsharedToTenantId() string {
	return getGenericEventParameter(p.Data, "unsharedToTenantId")
}

func (p *EventParameterRecordUnshared) SetTypeName(val string) {
	setGenericEventParameter(&p.Data, "typeName", val)
}

func (p *EventParameterRecordUnshared) GetTypeName() string {
	return getGenericEventParameter(p.Data, "typeName")
}

func (p *EventParameterRecordUnshared) SetRecordId(val string) {
	setGenericEventParameter(&p.Data, "recordId", val)
}

func (p *EventParameterRecordUnshared) GetRecordId() string {
	return getGenericEventParameter(p.Data, "recordId")
}

func UnaryServerEventLogger(fn func(ctx context.Context, info *grpc.UnaryServerInfo, req, res interface{}, err error)) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		res, err := handler(ctx, req)
		fn(ctx, info, req, res, err)
		return res, err
	}
}

func PrintEventTenantCreatedLog(ctx context.Context, tenantId, userId string, createdTenantId string) {
	p := &EventParameterTenantCreated{}
	p.SetCreatedTenantId(createdTenantId)

	writeEventLog(ctx, tenantId, userId, EventName_TenantCreated, pb.MessageCode_EVENT_TENANT_CREATED, p.Data)
}

func PrintRecordSharedLog(ctx context.Context, tenantId, userId string, sharedByTenantId string, sharedToTenantId string, typeName string, recordID string) {
	p := &EventParameterRecordShared{}
	p.SetSharedByTenantId(sharedByTenantId)
	p.SetSharedToTenantId(sharedToTenantId)
	p.SetTypeName(typeName)
	p.SetRecordId(recordID)

	writeEventLog(ctx, tenantId, userId, EventName_RecordShared, pb.MessageCode_EVENT_RECORD_SHARED, p.Data)
}

func PrintRecordUnsharedLog(ctx context.Context, tenantId, userId string, unsharedByTenantId string, unsharedToTenantId string, typeName string, recordID string) {
	p := &EventParameterRecordUnshared{}
	p.SetUnsharedByTenantId(unsharedByTenantId)
	p.SetUnsharedToTenantId(unsharedToTenantId)
	p.SetTypeName(typeName)
	p.SetRecordId(recordID)

	writeEventLog(ctx, tenantId, userId, EventName_RecordUnshared, pb.MessageCode_EVENT_RECORD_UNSHARED, p.Data)
}

func PrintRecordCreatedLog(ctx context.Context, tenantId, userId string, typeName string, recordID string) {
	p := &EventParameterRecordCreated{}
	p.SetTypeName(typeName)
	p.SetRecordId(recordID)

	writeEventLog(ctx, tenantId, userId, EventName_RecordCreated, pb.MessageCode_EVENT_RECORD_CREATED, p.Data)
}

func PrintRecordUpdatedLog(ctx context.Context, tenantId, userId string, typeName string, recordID string) {
	p := &EventParameterRecordUpdated{}
	p.SetTypeName(typeName)
	p.SetRecordId(recordID)

	writeEventLog(ctx, tenantId, userId, EventName_RecordUpdated, pb.MessageCode_EVENT_RECORD_UPDATED, p.Data)
}

func PrintRecordDeletedLog(ctx context.Context, tenantId, userId string, typeName string, recordID string) {
	p := &EventParameterRecordDeleted{}
	p.SetTypeName(typeName)
	p.SetRecordId(recordID)

	writeEventLog(ctx, tenantId, userId, EventName_RecordDeleted, pb.MessageCode_EVENT_RECORD_DELETED, p.Data)
}

func writeEventLog(ctx context.Context, tenantId, userId, auditSubcategory string, message pb.MessageCode, parameters *pb.GenericEventParameter) {
	logger := ctxlogrus.Extract(ctx)

	parameterData := &anypb.Any{}
	if err := anypb.MarshalFrom(parameterData, parameters, proto.MarshalOptions{}); err != nil {
		logger.WithError(err).Error("failed to marshal anypb")
		return
	}

	var jsonMessage json.RawMessage
	if v, err := protojson.Marshal(parameterData); err != nil {
		logger.WithError(err).Error("failed to protojson marshal")
		return
	} else {
		jsonMessage = v
	}

	if !json.Valid(jsonMessage) {
		logger.Error("invaid json message")
		return
	}

	writeLog(ctx, logRequest{
		EventID:          uuid.New().String(),
		TenantID:         tenantId,
		UserID:           userId,
		AuditCategory:    pb.AuditCategory_EVENT.String(),
		AuditSubcategory: auditSubcategory,
		Message:          message,
		Parameters:       jsonMessage,
	})
}
