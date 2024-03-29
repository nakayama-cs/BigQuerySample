// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: surveybasictypes.proto

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

// 自動催促設定
type AutoReminderSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 自動催促
	AutoReminder bool `protobuf:"varint,1,opt,name=auto_reminder,json=autoReminder,proto3" json:"auto_reminder,omitempty"`
	// 初回サイト
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// 初回前後区分
	// 名称マスタにおける種別 A1000006 (初回前後区分) を参照
	OffsetType *protobuf.NameOption `protobuf:"bytes,3,opt,name=offset_type,json=offsetType,proto3" json:"offset_type,omitempty"`
	// 繰り返し
	Repeat bool `protobuf:"varint,4,opt,name=repeat,proto3" json:"repeat,omitempty"`
	// 繰り返しサイト
	RepeatInterval int64 `protobuf:"varint,5,opt,name=repeat_interval,json=repeatInterval,proto3" json:"repeat_interval,omitempty"`
	// メッセージ区分
	// 名称マスタにおける種別 A1000003 (催促メッセージ区分) を参照
	ContentType *protobuf.NameOption `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// メッセージ内容
	Content string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	// システム通知先リスト
	// 名称マスタにおける種別 A1000005 (通知先設定) を参照
	SystemNotificationTos []*protobuf.NameOption `protobuf:"bytes,8,rep,name=system_notification_tos,json=systemNotificationTos,proto3" json:"system_notification_tos,omitempty"`
}

func (x *AutoReminderSetting) Reset() {
	*x = AutoReminderSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveybasictypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoReminderSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoReminderSetting) ProtoMessage() {}

func (x *AutoReminderSetting) ProtoReflect() protoreflect.Message {
	mi := &file_surveybasictypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoReminderSetting.ProtoReflect.Descriptor instead.
func (*AutoReminderSetting) Descriptor() ([]byte, []int) {
	return file_surveybasictypes_proto_rawDescGZIP(), []int{0}
}

func (x *AutoReminderSetting) GetAutoReminder() bool {
	if x != nil {
		return x.AutoReminder
	}
	return false
}

func (x *AutoReminderSetting) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AutoReminderSetting) GetOffsetType() *protobuf.NameOption {
	if x != nil {
		return x.OffsetType
	}
	return nil
}

func (x *AutoReminderSetting) GetRepeat() bool {
	if x != nil {
		return x.Repeat
	}
	return false
}

func (x *AutoReminderSetting) GetRepeatInterval() int64 {
	if x != nil {
		return x.RepeatInterval
	}
	return 0
}

func (x *AutoReminderSetting) GetContentType() *protobuf.NameOption {
	if x != nil {
		return x.ContentType
	}
	return nil
}

func (x *AutoReminderSetting) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AutoReminderSetting) GetSystemNotificationTos() []*protobuf.NameOption {
	if x != nil {
		return x.SystemNotificationTos
	}
	return nil
}

// 取引先情報
type BusinessUnitProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 取引先管理ID
	BusinessUnitManagementId string `protobuf:"bytes,1,opt,name=business_unit_management_id,json=businessUnitManagementId,proto3" json:"business_unit_management_id,omitempty"`
	// 取引先コード
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// 取引先名
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// 取引先担当者
	MainContact []*protobuf.UserReference `protobuf:"bytes,4,rep,name=main_contact,json=mainContact,proto3" json:"main_contact,omitempty"`
	// 取引先アシスタント
	MainContactAssistant []*protobuf.UserReference `protobuf:"bytes,5,rep,name=main_contact_assistant,json=mainContactAssistant,proto3" json:"main_contact_assistant,omitempty"`
	// 自社担当者
	Staff []*protobuf.UserReference `protobuf:"bytes,6,rep,name=staff,proto3" json:"staff,omitempty"`
	// 自社アシスタント
	StaffAssistant []*protobuf.UserReference `protobuf:"bytes,7,rep,name=staff_assistant,json=staffAssistant,proto3" json:"staff_assistant,omitempty"`
	// 取引先会社ID
	CompanyId string `protobuf:"bytes,8,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *BusinessUnitProperties) Reset() {
	*x = BusinessUnitProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveybasictypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessUnitProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessUnitProperties) ProtoMessage() {}

func (x *BusinessUnitProperties) ProtoReflect() protoreflect.Message {
	mi := &file_surveybasictypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessUnitProperties.ProtoReflect.Descriptor instead.
func (*BusinessUnitProperties) Descriptor() ([]byte, []int) {
	return file_surveybasictypes_proto_rawDescGZIP(), []int{1}
}

func (x *BusinessUnitProperties) GetBusinessUnitManagementId() string {
	if x != nil {
		return x.BusinessUnitManagementId
	}
	return ""
}

func (x *BusinessUnitProperties) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BusinessUnitProperties) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *BusinessUnitProperties) GetMainContact() []*protobuf.UserReference {
	if x != nil {
		return x.MainContact
	}
	return nil
}

func (x *BusinessUnitProperties) GetMainContactAssistant() []*protobuf.UserReference {
	if x != nil {
		return x.MainContactAssistant
	}
	return nil
}

func (x *BusinessUnitProperties) GetStaff() []*protobuf.UserReference {
	if x != nil {
		return x.Staff
	}
	return nil
}

func (x *BusinessUnitProperties) GetStaffAssistant() []*protobuf.UserReference {
	if x != nil {
		return x.StaffAssistant
	}
	return nil
}

func (x *BusinessUnitProperties) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

// 承認依頼
type ApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 承認依頼日時
	RequestedAt int64 `protobuf:"varint,1,opt,name=requested_at,json=requestedAt,proto3" json:"requested_at,omitempty"`
	// 承認依頼ユーザー
	RequestedBy *protobuf.EmbeddedUser `protobuf:"bytes,2,opt,name=requested_by,json=requestedBy,proto3" json:"requested_by,omitempty"`
	// 承認依頼担当者
	RequestedStaff *protobuf.StaffReference `protobuf:"bytes,3,opt,name=requested_staff,json=requestedStaff,proto3" json:"requested_staff,omitempty"`
	// 承認依頼部門
	RequestedOrganization *protobuf.ComponentUnitReference `protobuf:"bytes,4,opt,name=requested_organization,json=requestedOrganization,proto3" json:"requested_organization,omitempty"`
	// 承認担当者（予定）
	ApprovalPlanStaff *protobuf.StaffReference `protobuf:"bytes,5,opt,name=approval_plan_staff,json=approvalPlanStaff,proto3" json:"approval_plan_staff,omitempty"`
	// 承認部門（予定）
	ApprovalPlanOrganization *protobuf.ComponentUnitReference `protobuf:"bytes,6,opt,name=approval_plan_organization,json=approvalPlanOrganization,proto3" json:"approval_plan_organization,omitempty"`
	// 承認依頼コメント
	RequestedComment string `protobuf:"bytes,7,opt,name=requested_comment,json=requestedComment,proto3" json:"requested_comment,omitempty"`
}

func (x *ApprovalRequest) Reset() {
	*x = ApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveybasictypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalRequest) ProtoMessage() {}

func (x *ApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_surveybasictypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalRequest.ProtoReflect.Descriptor instead.
func (*ApprovalRequest) Descriptor() ([]byte, []int) {
	return file_surveybasictypes_proto_rawDescGZIP(), []int{2}
}

func (x *ApprovalRequest) GetRequestedAt() int64 {
	if x != nil {
		return x.RequestedAt
	}
	return 0
}

func (x *ApprovalRequest) GetRequestedBy() *protobuf.EmbeddedUser {
	if x != nil {
		return x.RequestedBy
	}
	return nil
}

func (x *ApprovalRequest) GetRequestedStaff() *protobuf.StaffReference {
	if x != nil {
		return x.RequestedStaff
	}
	return nil
}

func (x *ApprovalRequest) GetRequestedOrganization() *protobuf.ComponentUnitReference {
	if x != nil {
		return x.RequestedOrganization
	}
	return nil
}

func (x *ApprovalRequest) GetApprovalPlanStaff() *protobuf.StaffReference {
	if x != nil {
		return x.ApprovalPlanStaff
	}
	return nil
}

func (x *ApprovalRequest) GetApprovalPlanOrganization() *protobuf.ComponentUnitReference {
	if x != nil {
		return x.ApprovalPlanOrganization
	}
	return nil
}

func (x *ApprovalRequest) GetRequestedComment() string {
	if x != nil {
		return x.RequestedComment
	}
	return ""
}

// 承認実績
type ApprovalResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 承認実績日時
	DecidedAt int64 `protobuf:"varint,1,opt,name=decided_at,json=decidedAt,proto3" json:"decided_at,omitempty"`
	// 承認実績ユーザー
	DecidedBy *protobuf.EmbeddedUser `protobuf:"bytes,2,opt,name=decided_by,json=decidedBy,proto3" json:"decided_by,omitempty"`
	// 承認担当者（実績）
	DecidedStaff *protobuf.StaffReference `protobuf:"bytes,3,opt,name=decided_staff,json=decidedStaff,proto3" json:"decided_staff,omitempty"`
	// 承認部門（実績）
	DecidedOrganization *protobuf.ComponentUnitReference `protobuf:"bytes,4,opt,name=decided_organization,json=decidedOrganization,proto3" json:"decided_organization,omitempty"`
	// 承認実績コメント
	DecidedComment string `protobuf:"bytes,5,opt,name=decided_comment,json=decidedComment,proto3" json:"decided_comment,omitempty"`
}

func (x *ApprovalResult) Reset() {
	*x = ApprovalResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveybasictypes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalResult) ProtoMessage() {}

func (x *ApprovalResult) ProtoReflect() protoreflect.Message {
	mi := &file_surveybasictypes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalResult.ProtoReflect.Descriptor instead.
func (*ApprovalResult) Descriptor() ([]byte, []int) {
	return file_surveybasictypes_proto_rawDescGZIP(), []int{3}
}

func (x *ApprovalResult) GetDecidedAt() int64 {
	if x != nil {
		return x.DecidedAt
	}
	return 0
}

func (x *ApprovalResult) GetDecidedBy() *protobuf.EmbeddedUser {
	if x != nil {
		return x.DecidedBy
	}
	return nil
}

func (x *ApprovalResult) GetDecidedStaff() *protobuf.StaffReference {
	if x != nil {
		return x.DecidedStaff
	}
	return nil
}

func (x *ApprovalResult) GetDecidedOrganization() *protobuf.ComponentUnitReference {
	if x != nil {
		return x.DecidedOrganization
	}
	return nil
}

func (x *ApprovalResult) GetDecidedComment() string {
	if x != nil {
		return x.DecidedComment
	}
	return ""
}

// 付属情報付き添付ファイル
type AttachmentAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 添付ファイル
	Attachment *protobuf.Attachment `protobuf:"bytes,1,opt,name=attachment,proto3" json:"attachment,omitempty"`
	// 取引先管理ID
	BusinessUnitManagementId string `protobuf:"bytes,2,opt,name=business_unit_management_id,json=businessUnitManagementId,proto3" json:"business_unit_management_id,omitempty"`
	// 取引先名
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *AttachmentAttribute) Reset() {
	*x = AttachmentAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveybasictypes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentAttribute) ProtoMessage() {}

func (x *AttachmentAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_surveybasictypes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentAttribute.ProtoReflect.Descriptor instead.
func (*AttachmentAttribute) Descriptor() ([]byte, []int) {
	return file_surveybasictypes_proto_rawDescGZIP(), []int{4}
}

func (x *AttachmentAttribute) GetAttachment() *protobuf.Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *AttachmentAttribute) GetBusinessUnitManagementId() string {
	if x != nil {
		return x.BusinessUnitManagementId
	}
	return ""
}

func (x *AttachmentAttribute) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// 実行情報
type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 実行日時
	ProcessedAt int64 `protobuf:"varint,1,opt,name=processed_at,json=processedAt,proto3" json:"processed_at,omitempty"`
	// 実行ユーザー
	ProcessedBy *protobuf.EmbeddedUser `protobuf:"bytes,2,opt,name=processed_by,json=processedBy,proto3" json:"processed_by,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surveybasictypes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_surveybasictypes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_surveybasictypes_proto_rawDescGZIP(), []int{5}
}

func (x *Process) GetProcessedAt() int64 {
	if x != nil {
		return x.ProcessedAt
	}
	return 0
}

func (x *Process) GetProcessedBy() *protobuf.EmbeddedUser {
	if x != nil {
		return x.ProcessedBy
	}
	return nil
}

var File_surveybasictypes_proto protoreflect.FileDescriptor

var file_surveybasictypes_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x62, 0x61, 0x73, 0x69, 0x63, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x1a, 0x14,
	0x6d, 0x74, 0x6e, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x02, 0x0a, 0x13, 0x41, 0x75,
	0x74, 0x6f, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x35,
	0x0a, 0x0b, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x17, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x15, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x73, 0x22, 0xd1, 0x03, 0x0a, 0x16, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x75,
	0x6e, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x55, 0x6e, 0x69, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02,
	0x02, 0x08, 0x03, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x42, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42,
	0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x12, 0x55, 0x0a, 0x16, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06, 0xc2,
	0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x14, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x12, 0x48, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x0e, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0xe2, 0x03, 0x0a, 0x0f,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x41, 0x0a,
	0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x12, 0x57, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x13, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x11, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x12, 0x5e, 0x0a, 0x1a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c,
	0x69, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x18, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0xa3, 0x02, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69,
	0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09,
	0x64, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64, 0x42, 0x79, 0x12, 0x3d, 0x0a, 0x0d, 0x64, 0x65, 0x63,
	0x69, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0c, 0x64, 0x65, 0x63, 0x69,
	0x64, 0x65, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x53, 0x0a, 0x14, 0x64, 0x65, 0x63, 0x69,
	0x64, 0x65, 0x64, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x13, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65,
	0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a,
	0x0f, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x69, 0x64, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x34,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x77, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x29, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x41, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02,
	0x08, 0x03, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x42, 0x79, 0x42,
	0x1b, 0x5a, 0x19, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_surveybasictypes_proto_rawDescOnce sync.Once
	file_surveybasictypes_proto_rawDescData = file_surveybasictypes_proto_rawDesc
)

func file_surveybasictypes_proto_rawDescGZIP() []byte {
	file_surveybasictypes_proto_rawDescOnce.Do(func() {
		file_surveybasictypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_surveybasictypes_proto_rawDescData)
	})
	return file_surveybasictypes_proto_rawDescData
}

var file_surveybasictypes_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_surveybasictypes_proto_goTypes = []interface{}{
	(*AutoReminderSetting)(nil),             // 0: mtechnavi.api.survey.AutoReminderSetting
	(*BusinessUnitProperties)(nil),          // 1: mtechnavi.api.survey.BusinessUnitProperties
	(*ApprovalRequest)(nil),                 // 2: mtechnavi.api.survey.ApprovalRequest
	(*ApprovalResult)(nil),                  // 3: mtechnavi.api.survey.ApprovalResult
	(*AttachmentAttribute)(nil),             // 4: mtechnavi.api.survey.AttachmentAttribute
	(*Process)(nil),                         // 5: mtechnavi.api.survey.Process
	(*protobuf.NameOption)(nil),             // 6: sharelib.NameOption
	(*protobuf.UserReference)(nil),          // 7: sharelib.UserReference
	(*protobuf.EmbeddedUser)(nil),           // 8: sharelib.EmbeddedUser
	(*protobuf.StaffReference)(nil),         // 9: sharelib.StaffReference
	(*protobuf.ComponentUnitReference)(nil), // 10: sharelib.ComponentUnitReference
	(*protobuf.Attachment)(nil),             // 11: sharelib.Attachment
}
var file_surveybasictypes_proto_depIdxs = []int32{
	6,  // 0: mtechnavi.api.survey.AutoReminderSetting.offset_type:type_name -> sharelib.NameOption
	6,  // 1: mtechnavi.api.survey.AutoReminderSetting.content_type:type_name -> sharelib.NameOption
	6,  // 2: mtechnavi.api.survey.AutoReminderSetting.system_notification_tos:type_name -> sharelib.NameOption
	7,  // 3: mtechnavi.api.survey.BusinessUnitProperties.main_contact:type_name -> sharelib.UserReference
	7,  // 4: mtechnavi.api.survey.BusinessUnitProperties.main_contact_assistant:type_name -> sharelib.UserReference
	7,  // 5: mtechnavi.api.survey.BusinessUnitProperties.staff:type_name -> sharelib.UserReference
	7,  // 6: mtechnavi.api.survey.BusinessUnitProperties.staff_assistant:type_name -> sharelib.UserReference
	8,  // 7: mtechnavi.api.survey.ApprovalRequest.requested_by:type_name -> sharelib.EmbeddedUser
	9,  // 8: mtechnavi.api.survey.ApprovalRequest.requested_staff:type_name -> sharelib.StaffReference
	10, // 9: mtechnavi.api.survey.ApprovalRequest.requested_organization:type_name -> sharelib.ComponentUnitReference
	9,  // 10: mtechnavi.api.survey.ApprovalRequest.approval_plan_staff:type_name -> sharelib.StaffReference
	10, // 11: mtechnavi.api.survey.ApprovalRequest.approval_plan_organization:type_name -> sharelib.ComponentUnitReference
	8,  // 12: mtechnavi.api.survey.ApprovalResult.decided_by:type_name -> sharelib.EmbeddedUser
	9,  // 13: mtechnavi.api.survey.ApprovalResult.decided_staff:type_name -> sharelib.StaffReference
	10, // 14: mtechnavi.api.survey.ApprovalResult.decided_organization:type_name -> sharelib.ComponentUnitReference
	11, // 15: mtechnavi.api.survey.AttachmentAttribute.attachment:type_name -> sharelib.Attachment
	8,  // 16: mtechnavi.api.survey.Process.processed_by:type_name -> sharelib.EmbeddedUser
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_surveybasictypes_proto_init() }
func file_surveybasictypes_proto_init() {
	if File_surveybasictypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_surveybasictypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoReminderSetting); i {
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
		file_surveybasictypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessUnitProperties); i {
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
		file_surveybasictypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalRequest); i {
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
		file_surveybasictypes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalResult); i {
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
		file_surveybasictypes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentAttribute); i {
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
		file_surveybasictypes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
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
			RawDescriptor: file_surveybasictypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_surveybasictypes_proto_goTypes,
		DependencyIndexes: file_surveybasictypes_proto_depIdxs,
		MessageInfos:      file_surveybasictypes_proto_msgTypes,
	}.Build()
	File_surveybasictypes_proto = out.File
	file_surveybasictypes_proto_rawDesc = nil
	file_surveybasictypes_proto_goTypes = nil
	file_surveybasictypes_proto_depIdxs = nil
}
