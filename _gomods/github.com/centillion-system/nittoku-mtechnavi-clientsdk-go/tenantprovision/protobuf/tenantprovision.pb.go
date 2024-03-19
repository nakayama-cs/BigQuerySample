// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: tenantprovision.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	protobuf "mtechnavi/idp/protobuf"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// required tenant_request_id
type GetTenantRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// テナント申請ID
	TenantRequestId string `protobuf:"bytes,1,opt,name=tenant_request_id,json=tenantRequestId,proto3" json:"tenant_request_id,omitempty"`
}

func (x *GetTenantRequestRequest) Reset() {
	*x = GetTenantRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenantprovision_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequestRequest) ProtoMessage() {}

func (x *GetTenantRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenantprovision_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequestRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequestRequest) Descriptor() ([]byte, []int) {
	return file_tenantprovision_proto_rawDescGZIP(), []int{0}
}

func (x *GetTenantRequestRequest) GetTenantRequestId() string {
	if x != nil {
		return x.TenantRequestId
	}
	return ""
}

type GetTenantRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// テナント申請
	TenantRequest *TenantRequest `protobuf:"bytes,1,opt,name=tenant_request,json=tenantRequest,proto3" json:"tenant_request,omitempty"`
}

func (x *GetTenantRequestResponse) Reset() {
	*x = GetTenantRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenantprovision_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequestResponse) ProtoMessage() {}

func (x *GetTenantRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenantprovision_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequestResponse.ProtoReflect.Descriptor instead.
func (*GetTenantRequestResponse) Descriptor() ([]byte, []int) {
	return file_tenantprovision_proto_rawDescGZIP(), []int{1}
}

func (x *GetTenantRequestResponse) GetTenantRequest() *TenantRequest {
	if x != nil {
		return x.TenantRequest
	}
	return nil
}

// required tenant_request
type VerifyTenantRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// テナント申請
	// required tenant_request_id
	// required tenant_display_name_lang
	// required admin_display_name
	// required admin_email
	TenantRequest *TenantRequest `protobuf:"bytes,1,opt,name=tenant_request,json=tenantRequest,proto3" json:"tenant_request,omitempty"`
}

func (x *VerifyTenantRequestRequest) Reset() {
	*x = VerifyTenantRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenantprovision_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTenantRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTenantRequestRequest) ProtoMessage() {}

func (x *VerifyTenantRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenantprovision_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTenantRequestRequest.ProtoReflect.Descriptor instead.
func (*VerifyTenantRequestRequest) Descriptor() ([]byte, []int) {
	return file_tenantprovision_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyTenantRequestRequest) GetTenantRequest() *TenantRequest {
	if x != nil {
		return x.TenantRequest
	}
	return nil
}

type VerifyTenantRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VerifyTenantRequestResponse) Reset() {
	*x = VerifyTenantRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenantprovision_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTenantRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTenantRequestResponse) ProtoMessage() {}

func (x *VerifyTenantRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenantprovision_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTenantRequestResponse.ProtoReflect.Descriptor instead.
func (*VerifyTenantRequestResponse) Descriptor() ([]byte, []int) {
	return file_tenantprovision_proto_rawDescGZIP(), []int{3}
}

// required tenant_request_id
// required verify_tenant_request_id
type ApplyTenantRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// テナント申請ID
	TenantRequestId string `protobuf:"bytes,1,opt,name=tenant_request_id,json=tenantRequestId,proto3" json:"tenant_request_id,omitempty"`
	// テナント申請確認ID
	VerifyTenantRequestId string `protobuf:"bytes,2,opt,name=verify_tenant_request_id,json=verifyTenantRequestId,proto3" json:"verify_tenant_request_id,omitempty"`
}

func (x *ApplyTenantRequestRequest) Reset() {
	*x = ApplyTenantRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenantprovision_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyTenantRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyTenantRequestRequest) ProtoMessage() {}

func (x *ApplyTenantRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenantprovision_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyTenantRequestRequest.ProtoReflect.Descriptor instead.
func (*ApplyTenantRequestRequest) Descriptor() ([]byte, []int) {
	return file_tenantprovision_proto_rawDescGZIP(), []int{4}
}

func (x *ApplyTenantRequestRequest) GetTenantRequestId() string {
	if x != nil {
		return x.TenantRequestId
	}
	return ""
}

func (x *ApplyTenantRequestRequest) GetVerifyTenantRequestId() string {
	if x != nil {
		return x.VerifyTenantRequestId
	}
	return ""
}

type ApplyTenantRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApplyTenantRequestResponse) Reset() {
	*x = ApplyTenantRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenantprovision_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyTenantRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyTenantRequestResponse) ProtoMessage() {}

func (x *ApplyTenantRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenantprovision_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyTenantRequestResponse.ProtoReflect.Descriptor instead.
func (*ApplyTenantRequestResponse) Descriptor() ([]byte, []int) {
	return file_tenantprovision_proto_rawDescGZIP(), []int{5}
}

type TenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// テナント申請ID
	TenantRequestId string           `protobuf:"bytes,1,opt,name=tenant_request_id,json=tenantRequestId,proto3" json:"tenant_request_id,omitempty"`
	Tenant          *protobuf.Tenant `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	TenantAdmin     *protobuf.User   `protobuf:"bytes,3,opt,name=tenant_admin,json=tenantAdmin,proto3" json:"tenant_admin,omitempty"`
	// 招待日時 readonly
	// Implementation Note::
	// 格納Format：yyyy/MM/dd HH:mm:ss(JST)
	InvitedAt string `protobuf:"bytes,4,opt,name=invited_at,json=invitedAt,proto3" json:"invited_at,omitempty"`
	// 申請受付日時 readonly
	// Implementation Note::
	// 格納Format：yyyy/MM/dd HH:mm:ss(JST)
	RequestVerifiedAt string `protobuf:"bytes,5,opt,name=request_verified_at,json=requestVerifiedAt,proto3" json:"request_verified_at,omitempty"`
	// ライセンスラベル readonly
	// テナントに組み込むライセンスの種類
	LicenseLabels []string `protobuf:"bytes,6,rep,name=license_labels,json=licenseLabels,proto3" json:"license_labels,omitempty"`
	// 招待元テナントID readonly
	// 招待元テナントのテナントID
	InviteBy string `protobuf:"bytes,7,opt,name=invite_by,json=inviteBy,proto3" json:"invite_by,omitempty"`
	// 招待元取引先管理ID readonly
	// 招待元の取引先管理マスタのrecord_id
	ReleateBusinessUnitManagementId string `protobuf:"bytes,8,opt,name=releate_business_unit_management_id,json=releateBusinessUnitManagementId,proto3" json:"releate_business_unit_management_id,omitempty"`
}

func (x *TenantRequest) Reset() {
	*x = TenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenantprovision_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantRequest) ProtoMessage() {}

func (x *TenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenantprovision_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantRequest.ProtoReflect.Descriptor instead.
func (*TenantRequest) Descriptor() ([]byte, []int) {
	return file_tenantprovision_proto_rawDescGZIP(), []int{6}
}

func (x *TenantRequest) GetTenantRequestId() string {
	if x != nil {
		return x.TenantRequestId
	}
	return ""
}

func (x *TenantRequest) GetTenant() *protobuf.Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

func (x *TenantRequest) GetTenantAdmin() *protobuf.User {
	if x != nil {
		return x.TenantAdmin
	}
	return nil
}

func (x *TenantRequest) GetInvitedAt() string {
	if x != nil {
		return x.InvitedAt
	}
	return ""
}

func (x *TenantRequest) GetRequestVerifiedAt() string {
	if x != nil {
		return x.RequestVerifiedAt
	}
	return ""
}

func (x *TenantRequest) GetLicenseLabels() []string {
	if x != nil {
		return x.LicenseLabels
	}
	return nil
}

func (x *TenantRequest) GetInviteBy() string {
	if x != nil {
		return x.InviteBy
	}
	return ""
}

func (x *TenantRequest) GetReleateBusinessUnitManagementId() string {
	if x != nil {
		return x.ReleateBusinessUnitManagementId
	}
	return ""
}

var File_tenantprovision_proto protoreflect.FileDescriptor

var file_tenantprovision_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61,
	0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x6f, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x71,
	0x0a, 0x1a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x0e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x80, 0x01, 0x0a, 0x19, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x8b, 0x03, 0x0a, 0x0d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x70, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a,
	0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x62,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x4c, 0x0a, 0x23, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1f,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x55,
	0x6e, 0x69, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x32,
	0xb9, 0x03, 0x0a, 0x16, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x36, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x8c, 0x01, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x89, 0x01, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61,
	0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x39, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x6d,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tenantprovision_proto_rawDescOnce sync.Once
	file_tenantprovision_proto_rawDescData = file_tenantprovision_proto_rawDesc
)

func file_tenantprovision_proto_rawDescGZIP() []byte {
	file_tenantprovision_proto_rawDescOnce.Do(func() {
		file_tenantprovision_proto_rawDescData = protoimpl.X.CompressGZIP(file_tenantprovision_proto_rawDescData)
	})
	return file_tenantprovision_proto_rawDescData
}

var file_tenantprovision_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tenantprovision_proto_goTypes = []interface{}{
	(*GetTenantRequestRequest)(nil),     // 0: mtechnavi.api.tenantprovision.GetTenantRequestRequest
	(*GetTenantRequestResponse)(nil),    // 1: mtechnavi.api.tenantprovision.GetTenantRequestResponse
	(*VerifyTenantRequestRequest)(nil),  // 2: mtechnavi.api.tenantprovision.VerifyTenantRequestRequest
	(*VerifyTenantRequestResponse)(nil), // 3: mtechnavi.api.tenantprovision.VerifyTenantRequestResponse
	(*ApplyTenantRequestRequest)(nil),   // 4: mtechnavi.api.tenantprovision.ApplyTenantRequestRequest
	(*ApplyTenantRequestResponse)(nil),  // 5: mtechnavi.api.tenantprovision.ApplyTenantRequestResponse
	(*TenantRequest)(nil),               // 6: mtechnavi.api.tenantprovision.TenantRequest
	(*protobuf.Tenant)(nil),             // 7: mtechnavi.api.idp.Tenant
	(*protobuf.User)(nil),               // 8: mtechnavi.api.idp.User
}
var file_tenantprovision_proto_depIdxs = []int32{
	6, // 0: mtechnavi.api.tenantprovision.GetTenantRequestResponse.tenant_request:type_name -> mtechnavi.api.tenantprovision.TenantRequest
	6, // 1: mtechnavi.api.tenantprovision.VerifyTenantRequestRequest.tenant_request:type_name -> mtechnavi.api.tenantprovision.TenantRequest
	7, // 2: mtechnavi.api.tenantprovision.TenantRequest.tenant:type_name -> mtechnavi.api.idp.Tenant
	8, // 3: mtechnavi.api.tenantprovision.TenantRequest.tenant_admin:type_name -> mtechnavi.api.idp.User
	0, // 4: mtechnavi.api.tenantprovision.TenantProvisionService.GetTenantRequest:input_type -> mtechnavi.api.tenantprovision.GetTenantRequestRequest
	2, // 5: mtechnavi.api.tenantprovision.TenantProvisionService.VerifyTenantRequest:input_type -> mtechnavi.api.tenantprovision.VerifyTenantRequestRequest
	4, // 6: mtechnavi.api.tenantprovision.TenantProvisionService.ApplyTenantRequest:input_type -> mtechnavi.api.tenantprovision.ApplyTenantRequestRequest
	1, // 7: mtechnavi.api.tenantprovision.TenantProvisionService.GetTenantRequest:output_type -> mtechnavi.api.tenantprovision.GetTenantRequestResponse
	3, // 8: mtechnavi.api.tenantprovision.TenantProvisionService.VerifyTenantRequest:output_type -> mtechnavi.api.tenantprovision.VerifyTenantRequestResponse
	5, // 9: mtechnavi.api.tenantprovision.TenantProvisionService.ApplyTenantRequest:output_type -> mtechnavi.api.tenantprovision.ApplyTenantRequestResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tenantprovision_proto_init() }
func file_tenantprovision_proto_init() {
	if File_tenantprovision_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tenantprovision_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantRequestRequest); i {
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
		file_tenantprovision_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantRequestResponse); i {
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
		file_tenantprovision_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTenantRequestRequest); i {
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
		file_tenantprovision_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTenantRequestResponse); i {
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
		file_tenantprovision_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyTenantRequestRequest); i {
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
		file_tenantprovision_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyTenantRequestResponse); i {
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
		file_tenantprovision_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantRequest); i {
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
			RawDescriptor: file_tenantprovision_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tenantprovision_proto_goTypes,
		DependencyIndexes: file_tenantprovision_proto_depIdxs,
		MessageInfos:      file_tenantprovision_proto_msgTypes,
	}.Build()
	File_tenantprovision_proto = out.File
	file_tenantprovision_proto_rawDesc = nil
	file_tenantprovision_proto_goTypes = nil
	file_tenantprovision_proto_depIdxs = nil
}
