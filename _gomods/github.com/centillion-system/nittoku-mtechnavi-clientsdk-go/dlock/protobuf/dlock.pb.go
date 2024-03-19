// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: dlock.proto

package protobuf

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LockStatus int32

const (
	LockStatus_AVAILABLE       LockStatus = 0
	LockStatus_LOCKED_BY_ME    LockStatus = 1
	LockStatus_LOCKED_BY_OTHER LockStatus = 2
)

// Enum value maps for LockStatus.
var (
	LockStatus_name = map[int32]string{
		0: "AVAILABLE",
		1: "LOCKED_BY_ME",
		2: "LOCKED_BY_OTHER",
	}
	LockStatus_value = map[string]int32{
		"AVAILABLE":       0,
		"LOCKED_BY_ME":    1,
		"LOCKED_BY_OTHER": 2,
	}
)

func (x LockStatus) Enum() *LockStatus {
	p := new(LockStatus)
	*p = x
	return p
}

func (x LockStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LockStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_dlock_proto_enumTypes[0].Descriptor()
}

func (LockStatus) Type() protoreflect.EnumType {
	return &file_dlock_proto_enumTypes[0]
}

func (x LockStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LockStatus.Descriptor instead.
func (LockStatus) EnumDescriptor() ([]byte, []int) {
	return file_dlock_proto_rawDescGZIP(), []int{0}
}

type Lock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// ロックを取得（Acquire）した場合のみ、取得可能
	LockId    string `protobuf:"bytes,2,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty"`
	LockedBy  string `protobuf:"bytes,3,opt,name=locked_by,json=lockedBy,proto3" json:"locked_by,omitempty"`
	ExpiredAt int64  `protobuf:"varint,4,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
}

func (x *Lock) Reset() {
	*x = Lock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dlock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lock) ProtoMessage() {}

func (x *Lock) ProtoReflect() protoreflect.Message {
	mi := &file_dlock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lock.ProtoReflect.Descriptor instead.
func (*Lock) Descriptor() ([]byte, []int) {
	return file_dlock_proto_rawDescGZIP(), []int{0}
}

func (x *Lock) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Lock) GetLockId() string {
	if x != nil {
		return x.LockId
	}
	return ""
}

func (x *Lock) GetLockedBy() string {
	if x != nil {
		return x.LockedBy
	}
	return ""
}

func (x *Lock) GetExpiredAt() int64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

type AcquireLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// ロック期間（秒）
	// 1s ~ 24h（86400s）の範囲で指定する
	Expire int64 `protobuf:"varint,2,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *AcquireLockRequest) Reset() {
	*x = AcquireLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dlock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcquireLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcquireLockRequest) ProtoMessage() {}

func (x *AcquireLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dlock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcquireLockRequest.ProtoReflect.Descriptor instead.
func (*AcquireLockRequest) Descriptor() ([]byte, []int) {
	return file_dlock_proto_rawDescGZIP(), []int{1}
}

func (x *AcquireLockRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AcquireLockRequest) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type GetLockStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetLockStatusRequest) Reset() {
	*x = GetLockStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dlock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLockStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLockStatusRequest) ProtoMessage() {}

func (x *GetLockStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dlock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLockStatusRequest.ProtoReflect.Descriptor instead.
func (*GetLockStatusRequest) Descriptor() ([]byte, []int) {
	return file_dlock_proto_rawDescGZIP(), []int{2}
}

func (x *GetLockStatusRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetLockStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status LockStatus `protobuf:"varint,1,opt,name=status,proto3,enum=mtechnavi.api.dlock.LockStatus" json:"status,omitempty"` // enumで定義
}

func (x *GetLockStatusResponse) Reset() {
	*x = GetLockStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dlock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLockStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLockStatusResponse) ProtoMessage() {}

func (x *GetLockStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dlock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLockStatusResponse.ProtoReflect.Descriptor instead.
func (*GetLockStatusResponse) Descriptor() ([]byte, []int) {
	return file_dlock_proto_rawDescGZIP(), []int{3}
}

func (x *GetLockStatusResponse) GetStatus() LockStatus {
	if x != nil {
		return x.Status
	}
	return LockStatus_AVAILABLE
}

type ExtendLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	LockId string `protobuf:"bytes,2,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty"`
	// ロック期間（秒）
	// 1s ~ 24h（86400s）の範囲で指定する
	Expire int64 `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"` // １秒以上でバリデート（1s ~ 24h）
}

func (x *ExtendLockRequest) Reset() {
	*x = ExtendLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dlock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendLockRequest) ProtoMessage() {}

func (x *ExtendLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dlock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendLockRequest.ProtoReflect.Descriptor instead.
func (*ExtendLockRequest) Descriptor() ([]byte, []int) {
	return file_dlock_proto_rawDescGZIP(), []int{4}
}

func (x *ExtendLockRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ExtendLockRequest) GetLockId() string {
	if x != nil {
		return x.LockId
	}
	return ""
}

func (x *ExtendLockRequest) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type ExtendLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lock *Lock `protobuf:"bytes,1,opt,name=lock,proto3" json:"lock,omitempty"`
}

func (x *ExtendLockResponse) Reset() {
	*x = ExtendLockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dlock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendLockResponse) ProtoMessage() {}

func (x *ExtendLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dlock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendLockResponse.ProtoReflect.Descriptor instead.
func (*ExtendLockResponse) Descriptor() ([]byte, []int) {
	return file_dlock_proto_rawDescGZIP(), []int{5}
}

func (x *ExtendLockResponse) GetLock() *Lock {
	if x != nil {
		return x.Lock
	}
	return nil
}

type ReleaseLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	LockId string `protobuf:"bytes,2,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty"`
}

func (x *ReleaseLockRequest) Reset() {
	*x = ReleaseLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dlock_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseLockRequest) ProtoMessage() {}

func (x *ReleaseLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dlock_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseLockRequest.ProtoReflect.Descriptor instead.
func (*ReleaseLockRequest) Descriptor() ([]byte, []int) {
	return file_dlock_proto_rawDescGZIP(), []int{6}
}

func (x *ReleaseLockRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReleaseLockRequest) GetLockId() string {
	if x != nil {
		return x.LockId
	}
	return ""
}

var File_dlock_proto protoreflect.FileDescriptor

var file_dlock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6c, 0x6f,
	0x63, 0x6b, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0x54, 0x0a, 0x12, 0x41, 0x63, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x22, 0x06, 0x18,
	0x80, 0xa3, 0x05, 0x28, 0x01, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x31, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x50, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x74, 0x65, 0x63,
	0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x68, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x20, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x4d, 0x0a, 0x12,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x51, 0x0a, 0x12, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x07,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x2a, 0x42,
	0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4c,
	0x4f, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52,
	0x10, 0x02, 0x32, 0xe4, 0x02, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x51, 0x0a,
	0x0b, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x27, 0x2e, 0x6d,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x63, 0x6b,
	0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x29, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6d,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61,
	0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x4e, 0x0a, 0x0b, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x27, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x74, 0x65,
	0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x64, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dlock_proto_rawDescOnce sync.Once
	file_dlock_proto_rawDescData = file_dlock_proto_rawDesc
)

func file_dlock_proto_rawDescGZIP() []byte {
	file_dlock_proto_rawDescOnce.Do(func() {
		file_dlock_proto_rawDescData = protoimpl.X.CompressGZIP(file_dlock_proto_rawDescData)
	})
	return file_dlock_proto_rawDescData
}

var file_dlock_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dlock_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_dlock_proto_goTypes = []interface{}{
	(LockStatus)(0),               // 0: mtechnavi.api.dlock.LockStatus
	(*Lock)(nil),                  // 1: mtechnavi.api.dlock.Lock
	(*AcquireLockRequest)(nil),    // 2: mtechnavi.api.dlock.AcquireLockRequest
	(*GetLockStatusRequest)(nil),  // 3: mtechnavi.api.dlock.GetLockStatusRequest
	(*GetLockStatusResponse)(nil), // 4: mtechnavi.api.dlock.GetLockStatusResponse
	(*ExtendLockRequest)(nil),     // 5: mtechnavi.api.dlock.ExtendLockRequest
	(*ExtendLockResponse)(nil),    // 6: mtechnavi.api.dlock.ExtendLockResponse
	(*ReleaseLockRequest)(nil),    // 7: mtechnavi.api.dlock.ReleaseLockRequest
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_dlock_proto_depIdxs = []int32{
	0, // 0: mtechnavi.api.dlock.GetLockStatusResponse.status:type_name -> mtechnavi.api.dlock.LockStatus
	1, // 1: mtechnavi.api.dlock.ExtendLockResponse.lock:type_name -> mtechnavi.api.dlock.Lock
	2, // 2: mtechnavi.api.dlock.Locker.AcquireLock:input_type -> mtechnavi.api.dlock.AcquireLockRequest
	3, // 3: mtechnavi.api.dlock.Locker.GetLockStatus:input_type -> mtechnavi.api.dlock.GetLockStatusRequest
	5, // 4: mtechnavi.api.dlock.Locker.ExtendLock:input_type -> mtechnavi.api.dlock.ExtendLockRequest
	7, // 5: mtechnavi.api.dlock.Locker.ReleaseLock:input_type -> mtechnavi.api.dlock.ReleaseLockRequest
	1, // 6: mtechnavi.api.dlock.Locker.AcquireLock:output_type -> mtechnavi.api.dlock.Lock
	4, // 7: mtechnavi.api.dlock.Locker.GetLockStatus:output_type -> mtechnavi.api.dlock.GetLockStatusResponse
	1, // 8: mtechnavi.api.dlock.Locker.ExtendLock:output_type -> mtechnavi.api.dlock.Lock
	8, // 9: mtechnavi.api.dlock.Locker.ReleaseLock:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dlock_proto_init() }
func file_dlock_proto_init() {
	if File_dlock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dlock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lock); i {
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
		file_dlock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcquireLockRequest); i {
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
		file_dlock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLockStatusRequest); i {
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
		file_dlock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLockStatusResponse); i {
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
		file_dlock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendLockRequest); i {
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
		file_dlock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendLockResponse); i {
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
		file_dlock_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseLockRequest); i {
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
			RawDescriptor: file_dlock_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dlock_proto_goTypes,
		DependencyIndexes: file_dlock_proto_depIdxs,
		EnumInfos:         file_dlock_proto_enumTypes,
		MessageInfos:      file_dlock_proto_msgTypes,
	}.Build()
	File_dlock_proto = out.File
	file_dlock_proto_rawDesc = nil
	file_dlock_proto_goTypes = nil
	file_dlock_proto_depIdxs = nil
}
