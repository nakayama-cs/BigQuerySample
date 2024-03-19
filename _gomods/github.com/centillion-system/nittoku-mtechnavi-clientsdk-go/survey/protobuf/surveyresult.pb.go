// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: surveyresult.proto

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

// 依頼受信データ
//
// Implementation Note:
//
//	依頼明細データのshared時に生成されるデータ
type SurveyReception struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 依頼受信ID
	// 依頼明細ID survey_request_id と同様
	SurveyReceptionId string `protobuf:"bytes,1,opt,name=survey_reception_id,json=surveyReceptionId,proto3" json:"survey_reception_id,omitempty"`
	// システム通知先
	SystemNotificationUsers []*protobuf.UserReference `protobuf:"bytes,9,rep,name=system_notification_users,json=systemNotificationUsers,proto3" json:"system_notification_users,omitempty"`
	// readonly 依頼受信日
	//
	// Implementation Note:
	//
	//	依頼明細データの公開情報.shared_at
	ReceiptedAt int64 `protobuf:"varint,4,opt,name=receipted_at,json=receiptedAt,proto3" json:"receipted_at,omitempty"`
	CreatedAt   int64 `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   int64 `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   int64 `protobuf:"varint,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *SurveyReception) Reset() {
	*x = SurveyReception{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveyresult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyReception) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyReception) ProtoMessage() {}

func (x *SurveyReception) ProtoReflect() protoreflect.Message {
	mi := &file_surveyresult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyReception.ProtoReflect.Descriptor instead.
func (*SurveyReception) Descriptor() ([]byte, []int) {
	return file_surveyresult_proto_rawDescGZIP(), []int{0}
}

func (x *SurveyReception) GetSurveyReceptionId() string {
	if x != nil {
		return x.SurveyReceptionId
	}
	return ""
}

func (x *SurveyReception) GetSystemNotificationUsers() []*protobuf.UserReference {
	if x != nil {
		return x.SystemNotificationUsers
	}
	return nil
}

func (x *SurveyReception) GetReceiptedAt() int64 {
	if x != nil {
		return x.ReceiptedAt
	}
	return 0
}

func (x *SurveyReception) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SurveyReception) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *SurveyReception) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

// 依頼回答データ
type SurveyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 依頼回答ID
	SurveyResultId string `protobuf:"bytes,1,opt,name=survey_result_id,json=surveyResultId,proto3" json:"survey_result_id,omitempty"`
	// 依頼明細ID
	SurveyRequestId string `protobuf:"bytes,2,opt,name=survey_request_id,json=surveyRequestId,proto3" json:"survey_request_id,omitempty"`
	// readonly 回答送信日時
	SendedAt int64 `protobuf:"varint,3,opt,name=sended_at,json=sendedAt,proto3" json:"sended_at,omitempty"`
	// 部分回答
	PartialReplyed bool `protobuf:"varint,4,opt,name=partial_replyed,json=partialReplyed,proto3" json:"partial_replyed,omitempty"`
	// 辞退
	Declined bool `protobuf:"varint,5,opt,name=declined,proto3" json:"declined,omitempty"`
	// コメント
	Comment string `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	// 回答ファイルリスト
	ReplyAttachments []*protobuf.Attachment `protobuf:"bytes,7,rep,name=reply_attachments,json=replyAttachments,proto3" json:"reply_attachments,omitempty"`
	// readonly 共有情報
	SharedProperties *protobuf.EmbeddedSharedProperties `protobuf:"bytes,8,opt,name=shared_properties,json=sharedProperties,proto3" json:"shared_properties,omitempty"`
	CreatedAt        int64                              `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        int64                              `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt        int64                              `protobuf:"varint,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *SurveyResult) Reset() {
	*x = SurveyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveyresult_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyResult) ProtoMessage() {}

func (x *SurveyResult) ProtoReflect() protoreflect.Message {
	mi := &file_surveyresult_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyResult.ProtoReflect.Descriptor instead.
func (*SurveyResult) Descriptor() ([]byte, []int) {
	return file_surveyresult_proto_rawDescGZIP(), []int{1}
}

func (x *SurveyResult) GetSurveyResultId() string {
	if x != nil {
		return x.SurveyResultId
	}
	return ""
}

func (x *SurveyResult) GetSurveyRequestId() string {
	if x != nil {
		return x.SurveyRequestId
	}
	return ""
}

func (x *SurveyResult) GetSendedAt() int64 {
	if x != nil {
		return x.SendedAt
	}
	return 0
}

func (x *SurveyResult) GetPartialReplyed() bool {
	if x != nil {
		return x.PartialReplyed
	}
	return false
}

func (x *SurveyResult) GetDeclined() bool {
	if x != nil {
		return x.Declined
	}
	return false
}

func (x *SurveyResult) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *SurveyResult) GetReplyAttachments() []*protobuf.Attachment {
	if x != nil {
		return x.ReplyAttachments
	}
	return nil
}

func (x *SurveyResult) GetSharedProperties() *protobuf.EmbeddedSharedProperties {
	if x != nil {
		return x.SharedProperties
	}
	return nil
}

func (x *SurveyResult) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SurveyResult) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *SurveyResult) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

// 修正依頼データ
type SurveyResultChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 修正依頼ID
	SurveyResultChangeRequestId string `protobuf:"bytes,1,opt,name=survey_result_change_request_id,json=surveyResultChangeRequestId,proto3" json:"survey_result_change_request_id,omitempty"`
	// 依頼明細ID
	SurveyRequestId string `protobuf:"bytes,2,opt,name=survey_request_id,json=surveyRequestId,proto3" json:"survey_request_id,omitempty"`
	// 依頼回答ID
	SurveyResultId string `protobuf:"bytes,3,opt,name=survey_result_id,json=surveyResultId,proto3" json:"survey_result_id,omitempty"`
	// 修正依頼日時
	RequestedAt int64 `protobuf:"varint,4,opt,name=requested_at,json=requestedAt,proto3" json:"requested_at,omitempty"`
	// コメント
	Comment string `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	// readonly 共有情報
	SharedProperties *protobuf.EmbeddedSharedProperties `protobuf:"bytes,6,opt,name=shared_properties,json=sharedProperties,proto3" json:"shared_properties,omitempty"`
	CreatedAt        int64                              `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        int64                              `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt        int64                              `protobuf:"varint,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *SurveyResultChangeRequest) Reset() {
	*x = SurveyResultChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveyresult_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyResultChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyResultChangeRequest) ProtoMessage() {}

func (x *SurveyResultChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_surveyresult_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyResultChangeRequest.ProtoReflect.Descriptor instead.
func (*SurveyResultChangeRequest) Descriptor() ([]byte, []int) {
	return file_surveyresult_proto_rawDescGZIP(), []int{2}
}

func (x *SurveyResultChangeRequest) GetSurveyResultChangeRequestId() string {
	if x != nil {
		return x.SurveyResultChangeRequestId
	}
	return ""
}

func (x *SurveyResultChangeRequest) GetSurveyRequestId() string {
	if x != nil {
		return x.SurveyRequestId
	}
	return ""
}

func (x *SurveyResultChangeRequest) GetSurveyResultId() string {
	if x != nil {
		return x.SurveyResultId
	}
	return ""
}

func (x *SurveyResultChangeRequest) GetRequestedAt() int64 {
	if x != nil {
		return x.RequestedAt
	}
	return 0
}

func (x *SurveyResultChangeRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *SurveyResultChangeRequest) GetSharedProperties() *protobuf.EmbeddedSharedProperties {
	if x != nil {
		return x.SharedProperties
	}
	return nil
}

func (x *SurveyResultChangeRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SurveyResultChangeRequest) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *SurveyResultChangeRequest) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

// 回答ファイル出力済みデータ
type ExportedSurveyResultReplyAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExportedSurveyResultReplyAttachmentId string `protobuf:"bytes,1,opt,name=exported_survey_result_reply_attachment_id,json=exportedSurveyResultReplyAttachmentId,proto3" json:"exported_survey_result_reply_attachment_id,omitempty"`
	// 依頼回答ID
	SurveyResultId string `protobuf:"bytes,2,opt,name=survey_result_id,json=surveyResultId,proto3" json:"survey_result_id,omitempty"`
	// 依頼明細ID
	SurveyRequestId string `protobuf:"bytes,3,opt,name=survey_request_id,json=surveyRequestId,proto3" json:"survey_request_id,omitempty"`
	// created_by
	CreatedBy *protobuf.EmbeddedUser `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt int64                  `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64                  `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt int64                  `protobuf:"varint,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *ExportedSurveyResultReplyAttachment) Reset() {
	*x = ExportedSurveyResultReplyAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveyresult_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportedSurveyResultReplyAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportedSurveyResultReplyAttachment) ProtoMessage() {}

func (x *ExportedSurveyResultReplyAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_surveyresult_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportedSurveyResultReplyAttachment.ProtoReflect.Descriptor instead.
func (*ExportedSurveyResultReplyAttachment) Descriptor() ([]byte, []int) {
	return file_surveyresult_proto_rawDescGZIP(), []int{3}
}

func (x *ExportedSurveyResultReplyAttachment) GetExportedSurveyResultReplyAttachmentId() string {
	if x != nil {
		return x.ExportedSurveyResultReplyAttachmentId
	}
	return ""
}

func (x *ExportedSurveyResultReplyAttachment) GetSurveyResultId() string {
	if x != nil {
		return x.SurveyResultId
	}
	return ""
}

func (x *ExportedSurveyResultReplyAttachment) GetSurveyRequestId() string {
	if x != nil {
		return x.SurveyRequestId
	}
	return ""
}

func (x *ExportedSurveyResultReplyAttachment) GetCreatedBy() *protobuf.EmbeddedUser {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *ExportedSurveyResultReplyAttachment) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ExportedSurveyResultReplyAttachment) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ExportedSurveyResultReplyAttachment) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_surveyresult_proto protoreflect.FileDescriptor

var file_surveyresult_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x1a, 0x14, 0x6d, 0x74, 0x6e, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x6d, 0x74, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x02, 0x0a, 0x0f, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52,
	0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x63,
	0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x53, 0x0a, 0x19, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x17, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x14, 0xd2, 0xb8, 0x02, 0x10, 0x0a, 0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22,
	0x9a, 0x04, 0x0a, 0x0c, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x11, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0xd2, 0xb8, 0x02,
	0x15, 0x0a, 0x13, 0x0a, 0x11, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x52, 0x0f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02,
	0x08, 0x01, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2f, 0x0a, 0x0f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x65, 0x64, 0x12, 0x22, 0x0a,
	0x08, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x08, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65,
	0x64, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x10, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4f,
	0x0a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x10, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf3, 0x03, 0x0a,
	0x19, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x1f, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x1b, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x4b, 0x0a, 0x11, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xc2, 0xb8, 0x02,
	0x02, 0x08, 0x01, 0xd2, 0xb8, 0x02, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x52, 0x0f, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x48, 0x0a,
	0x10, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0xd2,
	0xb8, 0x02, 0x14, 0x0a, 0x12, 0x0a, 0x10, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x52, 0x0e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2,
	0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x10, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xab, 0x03, 0x0a, 0x23, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x59, 0x0a, 0x2a, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x25,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x10, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1e, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0xd2, 0xb8, 0x02, 0x14, 0x0a, 0x12, 0x0a, 0x10, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x52,
	0x0e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12,
	0x4b, 0x0a, 0x11, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xc2, 0xb8, 0x02, 0x02,
	0x08, 0x01, 0xd2, 0xb8, 0x02, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x52, 0x0f, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_surveyresult_proto_rawDescOnce sync.Once
	file_surveyresult_proto_rawDescData = file_surveyresult_proto_rawDesc
)

func file_surveyresult_proto_rawDescGZIP() []byte {
	file_surveyresult_proto_rawDescOnce.Do(func() {
		file_surveyresult_proto_rawDescData = protoimpl.X.CompressGZIP(file_surveyresult_proto_rawDescData)
	})
	return file_surveyresult_proto_rawDescData
}

var file_surveyresult_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_surveyresult_proto_goTypes = []interface{}{
	(*SurveyReception)(nil),                     // 0: mtechnavi.api.survey.SurveyReception
	(*SurveyResult)(nil),                        // 1: mtechnavi.api.survey.SurveyResult
	(*SurveyResultChangeRequest)(nil),           // 2: mtechnavi.api.survey.SurveyResultChangeRequest
	(*ExportedSurveyResultReplyAttachment)(nil), // 3: mtechnavi.api.survey.ExportedSurveyResultReplyAttachment
	(*protobuf.UserReference)(nil),              // 4: sharelib.UserReference
	(*protobuf.Attachment)(nil),                 // 5: sharelib.Attachment
	(*protobuf.EmbeddedSharedProperties)(nil),   // 6: sharelib.EmbeddedSharedProperties
	(*protobuf.EmbeddedUser)(nil),               // 7: sharelib.EmbeddedUser
}
var file_surveyresult_proto_depIdxs = []int32{
	4, // 0: mtechnavi.api.survey.SurveyReception.system_notification_users:type_name -> sharelib.UserReference
	5, // 1: mtechnavi.api.survey.SurveyResult.reply_attachments:type_name -> sharelib.Attachment
	6, // 2: mtechnavi.api.survey.SurveyResult.shared_properties:type_name -> sharelib.EmbeddedSharedProperties
	6, // 3: mtechnavi.api.survey.SurveyResultChangeRequest.shared_properties:type_name -> sharelib.EmbeddedSharedProperties
	7, // 4: mtechnavi.api.survey.ExportedSurveyResultReplyAttachment.created_by:type_name -> sharelib.EmbeddedUser
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_surveyresult_proto_init() }
func file_surveyresult_proto_init() {
	if File_surveyresult_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_surveyresult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyReception); i {
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
		file_surveyresult_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyResult); i {
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
		file_surveyresult_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyResultChangeRequest); i {
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
		file_surveyresult_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportedSurveyResultReplyAttachment); i {
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
			RawDescriptor: file_surveyresult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_surveyresult_proto_goTypes,
		DependencyIndexes: file_surveyresult_proto_depIdxs,
		MessageInfos:      file_surveyresult_proto_msgTypes,
	}.Build()
	File_surveyresult_proto = out.File
	file_surveyresult_proto_rawDesc = nil
	file_surveyresult_proto_goTypes = nil
	file_surveyresult_proto_depIdxs = nil
}
