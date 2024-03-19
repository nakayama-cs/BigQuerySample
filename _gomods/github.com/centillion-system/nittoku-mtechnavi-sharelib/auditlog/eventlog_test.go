package auditlog

import (
	"context"
	"fmt"
	pb "mtechnavi/sharelib/auditlog/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type DummyEventRequest struct {
	EventId    string
	EventName  string
	TenantId   string
	UserId     string
	Parameters *anypb.Any
}

func (d *DummyEventRequest) GetEventId() string {
	return d.EventId
}
func (d *DummyEventRequest) GetEventName() string {
	return d.EventName
}
func (d *DummyEventRequest) GetTenantId() string {
	return d.TenantId
}
func (d *DummyEventRequest) GetUserId() string {
	return d.UserId
}
func (d *DummyEventRequest) GetParameters() *anypb.Any {
	return d.Parameters
}

func TestEventParameterTenantCreated(t *testing.T) {
	p := &EventParameterTenantCreated{}

	assert.Equal(t, "", p.GetCreatedTenantId())

	p = &EventParameterTenantCreated{}

	p.SetCreatedTenantId("tenantId")
	assert.Equal(t, "tenantId", p.GetCreatedTenantId())
}

func TestEventParameterRecordCreated(t *testing.T) {
	p := &EventParameterRecordCreated{}

	assert.Equal(t, "", p.GetRecordId())
	assert.Equal(t, "", p.GetTypeName())

	p = &EventParameterRecordCreated{}

	p.SetRecordId("recordId")
	assert.Equal(t, "recordId", p.GetRecordId())
	p.SetTypeName("typeName")
	assert.Equal(t, "typeName", p.GetTypeName())
}

func TestEventParameterRecordUpdated(t *testing.T) {
	p := &EventParameterRecordUpdated{}

	assert.Equal(t, "", p.GetRecordId())
	assert.Equal(t, "", p.GetTypeName())

	p = &EventParameterRecordUpdated{}

	p.SetRecordId("recordId")
	assert.Equal(t, "recordId", p.GetRecordId())
	p.SetTypeName("typeName")
	assert.Equal(t, "typeName", p.GetTypeName())
}

func TestEventParameterRecordDeleted(t *testing.T) {
	p := &EventParameterRecordDeleted{}

	assert.Equal(t, "", p.GetRecordId())
	assert.Equal(t, "", p.GetTypeName())

	p = &EventParameterRecordDeleted{}

	p.SetRecordId("recordId")
	assert.Equal(t, "recordId", p.GetRecordId())
	p.SetTypeName("typeName")
	assert.Equal(t, "typeName", p.GetTypeName())
}

func TestEventParameterRecordShared(t *testing.T) {
	p := &EventParameterRecordShared{}

	assert.Equal(t, "", p.GetRecordId())
	assert.Equal(t, "", p.GetTypeName())
	assert.Equal(t, "", p.GetSharedByTenantId())
	assert.Equal(t, "", p.GetSharedToTenantId())

	p = &EventParameterRecordShared{}

	p.SetRecordId("recordId")
	assert.Equal(t, "recordId", p.GetRecordId())
	p.SetTypeName("typeName")
	assert.Equal(t, "typeName", p.GetTypeName())
	p.SetSharedByTenantId("sharedByTenantId")
	assert.Equal(t, "sharedByTenantId", p.GetSharedByTenantId())
	p.SetSharedToTenantId("sharedToTenantId")
	assert.Equal(t, "sharedToTenantId", p.GetSharedToTenantId())
}

func TestHandleTriggerEvent(t *testing.T) {

	//テストケース
	t.Run("Event呼び分け", func(t *testing.T) {
		ctx := context.Background()

		eventMap := map[string]func(context.Context, EventRequest, *pb.GenericEventParameter) error{
			EventName_RecordCreated: onRecordCreated,
			EventName_RecordShared:  onRecordShared,
		}

		p := &EventParameterTenantCreated{}
		p.SetCreatedTenantId("createdTenantId")

		parameterData := &anypb.Any{}
		if err := anypb.MarshalFrom(parameterData, p.Data, proto.MarshalOptions{}); err != nil {
			panic(err)
		}

		req := &DummyEventRequest{
			EventId:    "EventId",
			EventName:  "",
			TenantId:   "TenantId",
			UserId:     "UserId",
			Parameters: parameterData,
		}

		var err error
		req.EventName = EventName_RecordShared
		err = HandleTriggerEvent(ctx, req, eventMap)
		assert.Nil(t, err)

		req.EventName = EventName_RecordCreated
		err = HandleTriggerEvent(ctx, req, eventMap)
		assert.Nil(t, err)

		// Mapに存在しないイベントであってもErrorは発生しないのが正
		req.EventName = EventName_RecordUpdated
		err = HandleTriggerEvent(ctx, req, eventMap)
		assert.Nil(t, err)

		req.EventName = ""
		err = HandleTriggerEvent(ctx, req, eventMap)
		assert.Nil(t, err)
	})
}

func onRecordCreated(context.Context, EventRequest, *pb.GenericEventParameter) error {
	fmt.Println("call onRecordCreated")
	return nil
}

func onRecordShared(context.Context, EventRequest, *pb.GenericEventParameter) error {
	fmt.Println("call onRecordShared")
	return nil
}
