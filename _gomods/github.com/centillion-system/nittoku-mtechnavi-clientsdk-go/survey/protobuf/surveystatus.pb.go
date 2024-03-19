// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: surveystatus.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	protobuf "mtechnavi/sharelib/protobuf"
	_ "mtechnavi/sharelib/protobuf/mtn"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 終了データ
type Complete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 終了ID
	CompleteId string `protobuf:"bytes,1,opt,name=complete_id,json=completeId,proto3" json:"complete_id,omitempty"`
	// urn
	// "urn:mtechnavi.api.survey.SurveyRequest:{{survey_request_id}}"
	Urn string `protobuf:"bytes,2,opt,name=urn,proto3" json:"urn,omitempty"`
	// "mtechnavi.api.survey.SurveyRequest"
	TypeName string `protobuf:"bytes,9,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// "{{survey_request_id}}"
	RecordId string `protobuf:"bytes,10,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	// readonly 終了日時
	CompletedAt int64 `protobuf:"varint,8,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	// readonly 共有情報
	SharedProperties *protobuf.EmbeddedSharedProperties `protobuf:"bytes,4,opt,name=shared_properties,json=sharedProperties,proto3" json:"shared_properties,omitempty"`
	CreatedAt        int64                              `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        int64                              `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt        int64                              `protobuf:"varint,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Complete) Reset() {
	*x = Complete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveystatus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Complete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Complete) ProtoMessage() {}

func (x *Complete) ProtoReflect() protoreflect.Message {
	mi := &file_surveystatus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Complete.ProtoReflect.Descriptor instead.
func (*Complete) Descriptor() ([]byte, []int) {
	return file_surveystatus_proto_rawDescGZIP(), []int{0}
}

func (x *Complete) GetCompleteId() string {
	if x != nil {
		return x.CompleteId
	}
	return ""
}

func (x *Complete) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *Complete) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Complete) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *Complete) GetCompletedAt() int64 {
	if x != nil {
		return x.CompletedAt
	}
	return 0
}

func (x *Complete) GetSharedProperties() *protobuf.EmbeddedSharedProperties {
	if x != nil {
		return x.SharedProperties
	}
	return nil
}

func (x *Complete) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Complete) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Complete) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

// 破棄データ
type Discard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 破棄ID
	DiscardId string `protobuf:"bytes,1,opt,name=discard_id,json=discardId,proto3" json:"discard_id,omitempty"`
	// urn
	// "urn:mtechnavi.api.survey.SurveyRequest:{{survey_request_id}}"
	Urn string `protobuf:"bytes,2,opt,name=urn,proto3" json:"urn,omitempty"`
	// "mtechnavi.api.survey.SurveyRequest"
	TypeName string `protobuf:"bytes,9,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// "{{survey_request_id}}"
	RecordId string `protobuf:"bytes,10,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	// コメント
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	// readonly 破棄日時
	DiscardedAt int64 `protobuf:"varint,4,opt,name=discarded_at,json=discardedAt,proto3" json:"discarded_at,omitempty"`
	// readonly 共有情報
	SharedProperties *protobuf.EmbeddedSharedProperties `protobuf:"bytes,5,opt,name=shared_properties,json=sharedProperties,proto3" json:"shared_properties,omitempty"`
	CreatedAt        int64                              `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        int64                              `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt        int64                              `protobuf:"varint,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Discard) Reset() {
	*x = Discard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveystatus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discard) ProtoMessage() {}

func (x *Discard) ProtoReflect() protoreflect.Message {
	mi := &file_surveystatus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discard.ProtoReflect.Descriptor instead.
func (*Discard) Descriptor() ([]byte, []int) {
	return file_surveystatus_proto_rawDescGZIP(), []int{1}
}

func (x *Discard) GetDiscardId() string {
	if x != nil {
		return x.DiscardId
	}
	return ""
}

func (x *Discard) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *Discard) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Discard) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *Discard) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Discard) GetDiscardedAt() int64 {
	if x != nil {
		return x.DiscardedAt
	}
	return 0
}

func (x *Discard) GetSharedProperties() *protobuf.EmbeddedSharedProperties {
	if x != nil {
		return x.SharedProperties
	}
	return nil
}

func (x *Discard) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Discard) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Discard) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

// 開封済みデータ
type Opened struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 開封済みID
	OpenedId string `protobuf:"bytes,1,opt,name=opened_id,json=openedId,proto3" json:"opened_id,omitempty"`
	// urn
	// "urn:mtechnavi.api.survey.SurveyRequest:{{survey_request_id}}"
	Urn string `protobuf:"bytes,2,opt,name=urn,proto3" json:"urn,omitempty"`
	// "mtechnavi.api.survey.SurveyRequest"
	TypeName string `protobuf:"bytes,3,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// "{{survey_request_id}}"
	RecordId string `protobuf:"bytes,4,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	// 開封者情報
	OpenUserProcess *Process `protobuf:"bytes,5,opt,name=open_user_process,json=openUserProcess,proto3" json:"open_user_process,omitempty"`
	// readonly 共有情報
	SharedProperties *protobuf.EmbeddedSharedProperties `protobuf:"bytes,6,opt,name=shared_properties,json=sharedProperties,proto3" json:"shared_properties,omitempty"`
	CreatedAt        int64                              `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        int64                              `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt        int64                              `protobuf:"varint,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Opened) Reset() {
	*x = Opened{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveystatus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Opened) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Opened) ProtoMessage() {}

func (x *Opened) ProtoReflect() protoreflect.Message {
	mi := &file_surveystatus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Opened.ProtoReflect.Descriptor instead.
func (*Opened) Descriptor() ([]byte, []int) {
	return file_surveystatus_proto_rawDescGZIP(), []int{2}
}

func (x *Opened) GetOpenedId() string {
	if x != nil {
		return x.OpenedId
	}
	return ""
}

func (x *Opened) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *Opened) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Opened) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *Opened) GetOpenUserProcess() *Process {
	if x != nil {
		return x.OpenUserProcess
	}
	return nil
}

func (x *Opened) GetSharedProperties() *protobuf.EmbeddedSharedProperties {
	if x != nil {
		return x.SharedProperties
	}
	return nil
}

func (x *Opened) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Opened) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Opened) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_surveystatus_proto protoreflect.FileDescriptor

var file_surveystatus_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x1a, 0x14, 0x6d, 0x74, 0x6e, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x6d, 0x74, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x62, 0x61, 0x73, 0x69, 0x63, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x08, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0xd2, 0xb8, 0x02, 0x07,
	0x0a, 0x05, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x34, 0x0a, 0x09,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x17, 0xd2, 0xb8, 0x02, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x08, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2,
	0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x4f, 0x0a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65,
	0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x10, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0xa3, 0x03, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x63, 0x61,
	0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11,
	0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0xd2, 0xb8, 0x02, 0x07, 0x0a, 0x05, 0x0a, 0x03, 0x75, 0x72,
	0x6e, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xd2, 0xb8, 0x02, 0x0d, 0x0a,
	0x0b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0xb8, 0x02, 0x02,
	0x08, 0x01, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x09,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08,
	0x01, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4f,
	0x0a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x10, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa6, 0x03, 0x0a,
	0x06, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x65,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e,
	0x65, 0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x11, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0xd2, 0xb8, 0x02, 0x07, 0x0a, 0x05, 0x0a,
	0x03, 0x75, 0x72, 0x6e, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xd2, 0xb8,
	0x02, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xc2,
	0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x49, 0x64, 0x12, 0x51, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06,
	0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4f, 0x0a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x10, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61,
	0x76, 0x69, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_surveystatus_proto_rawDescOnce sync.Once
	file_surveystatus_proto_rawDescData = file_surveystatus_proto_rawDesc
)

func file_surveystatus_proto_rawDescGZIP() []byte {
	file_surveystatus_proto_rawDescOnce.Do(func() {
		file_surveystatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_surveystatus_proto_rawDescData)
	})
	return file_surveystatus_proto_rawDescData
}

var file_surveystatus_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_surveystatus_proto_goTypes = []interface{}{
	(*Complete)(nil), // 0: mtechnavi.api.survey.Complete
	(*Discard)(nil),  // 1: mtechnavi.api.survey.Discard
	(*Opened)(nil),   // 2: mtechnavi.api.survey.Opened
	(*protobuf.EmbeddedSharedProperties)(nil), // 3: sharelib.EmbeddedSharedProperties
	(*Process)(nil), // 4: mtechnavi.api.survey.Process
}
var file_surveystatus_proto_depIdxs = []int32{
	3, // 0: mtechnavi.api.survey.Complete.shared_properties:type_name -> sharelib.EmbeddedSharedProperties
	3, // 1: mtechnavi.api.survey.Discard.shared_properties:type_name -> sharelib.EmbeddedSharedProperties
	4, // 2: mtechnavi.api.survey.Opened.open_user_process:type_name -> mtechnavi.api.survey.Process
	3, // 3: mtechnavi.api.survey.Opened.shared_properties:type_name -> sharelib.EmbeddedSharedProperties
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_surveystatus_proto_init() }
func file_surveystatus_proto_init() {
	if File_surveystatus_proto != nil {
		return
	}
	file_surveybasictypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_surveystatus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Complete); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_surveystatus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Discard); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_surveystatus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Opened); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_surveystatus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_surveystatus_proto_goTypes,
		DependencyIndexes: file_surveystatus_proto_depIdxs,
		MessageInfos:      file_surveystatus_proto_msgTypes,
	}.Build()
	File_surveystatus_proto = out.File
	file_surveystatus_proto_rawDesc = nil
	file_surveystatus_proto_goTypes = nil
	file_surveystatus_proto_depIdxs = nil
}
