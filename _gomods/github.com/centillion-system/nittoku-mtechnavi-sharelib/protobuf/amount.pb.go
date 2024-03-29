// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: amount.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type DistantAmount_Unit int32

const (
	DistantAmount_NONE        DistantAmount_Unit = 0
	DistantAmount_METER       DistantAmount_Unit = 1
	DistantAmount_CENTI_METER DistantAmount_Unit = 2
	DistantAmount_MILLI_METER DistantAmount_Unit = 3
	DistantAmount_MICRO_METER DistantAmount_Unit = 4
	DistantAmount_NANO_METER  DistantAmount_Unit = 5
)

// Enum value maps for DistantAmount_Unit.
var (
	DistantAmount_Unit_name = map[int32]string{
		0: "NONE",
		1: "METER",
		2: "CENTI_METER",
		3: "MILLI_METER",
		4: "MICRO_METER",
		5: "NANO_METER",
	}
	DistantAmount_Unit_value = map[string]int32{
		"NONE":        0,
		"METER":       1,
		"CENTI_METER": 2,
		"MILLI_METER": 3,
		"MICRO_METER": 4,
		"NANO_METER":  5,
	}
)

func (x DistantAmount_Unit) Enum() *DistantAmount_Unit {
	p := new(DistantAmount_Unit)
	*p = x
	return p
}

func (x DistantAmount_Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DistantAmount_Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_amount_proto_enumTypes[0].Descriptor()
}

func (DistantAmount_Unit) Type() protoreflect.EnumType {
	return &file_amount_proto_enumTypes[0]
}

func (x DistantAmount_Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DistantAmount_Unit.Descriptor instead.
func (DistantAmount_Unit) EnumDescriptor() ([]byte, []int) {
	return file_amount_proto_rawDescGZIP(), []int{1, 0}
}

type MonetaryAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 通貨コード
	// JPY：日本
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// 整数
	IntegralAmount int64 `protobuf:"varint,2,opt,name=integral_amount,json=integralAmount,proto3" json:"integral_amount,omitempty"`
	// 少数（４桁）
	FractionalAmount int32 `protobuf:"varint,3,opt,name=fractional_amount,json=fractionalAmount,proto3" json:"fractional_amount,omitempty"`
}

func (x *MonetaryAmount) Reset() {
	*x = MonetaryAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonetaryAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonetaryAmount) ProtoMessage() {}

func (x *MonetaryAmount) ProtoReflect() protoreflect.Message {
	mi := &file_amount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonetaryAmount.ProtoReflect.Descriptor instead.
func (*MonetaryAmount) Descriptor() ([]byte, []int) {
	return file_amount_proto_rawDescGZIP(), []int{0}
}

func (x *MonetaryAmount) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *MonetaryAmount) GetIntegralAmount() int64 {
	if x != nil {
		return x.IntegralAmount
	}
	return 0
}

func (x *MonetaryAmount) GetFractionalAmount() int32 {
	if x != nil {
		return x.FractionalAmount
	}
	return 0
}

type DistantAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 単位コード
	Unit DistantAmount_Unit `protobuf:"varint,1,opt,name=unit,proto3,enum=sharelib.DistantAmount_Unit" json:"unit,omitempty"`
	// 整数
	IntegralAmount int64 `protobuf:"varint,2,opt,name=integral_amount,json=integralAmount,proto3" json:"integral_amount,omitempty"`
	// 少数（４桁）
	FractionalAmount int32 `protobuf:"varint,3,opt,name=fractional_amount,json=fractionalAmount,proto3" json:"fractional_amount,omitempty"`
}

func (x *DistantAmount) Reset() {
	*x = DistantAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistantAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistantAmount) ProtoMessage() {}

func (x *DistantAmount) ProtoReflect() protoreflect.Message {
	mi := &file_amount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistantAmount.ProtoReflect.Descriptor instead.
func (*DistantAmount) Descriptor() ([]byte, []int) {
	return file_amount_proto_rawDescGZIP(), []int{1}
}

func (x *DistantAmount) GetUnit() DistantAmount_Unit {
	if x != nil {
		return x.Unit
	}
	return DistantAmount_NONE
}

func (x *DistantAmount) GetIntegralAmount() int64 {
	if x != nil {
		return x.IntegralAmount
	}
	return 0
}

func (x *DistantAmount) GetFractionalAmount() int32 {
	if x != nil {
		return x.FractionalAmount
	}
	return 0
}

// 数量
type QuantityAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 名称マスタにおける種別 A0000005 (数量単位) を参照
	Unit *NameOption `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	// 整数
	IntegralAmount int64 `protobuf:"varint,2,opt,name=integral_amount,json=integralAmount,proto3" json:"integral_amount,omitempty"`
	// 小数（４桁）
	FractionalAmount int32 `protobuf:"varint,3,opt,name=fractional_amount,json=fractionalAmount,proto3" json:"fractional_amount,omitempty"`
}

func (x *QuantityAmount) Reset() {
	*x = QuantityAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amount_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantityAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantityAmount) ProtoMessage() {}

func (x *QuantityAmount) ProtoReflect() protoreflect.Message {
	mi := &file_amount_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantityAmount.ProtoReflect.Descriptor instead.
func (*QuantityAmount) Descriptor() ([]byte, []int) {
	return file_amount_proto_rawDescGZIP(), []int{2}
}

func (x *QuantityAmount) GetUnit() *NameOption {
	if x != nil {
		return x.Unit
	}
	return nil
}

func (x *QuantityAmount) GetIntegralAmount() int64 {
	if x != nil {
		return x.IntegralAmount
	}
	return 0
}

func (x *QuantityAmount) GetFractionalAmount() int32 {
	if x != nil {
		return x.FractionalAmount
	}
	return 0
}

// 数値
type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 整数
	IntegralAmount int64 `protobuf:"varint,1,opt,name=integral_amount,json=integralAmount,proto3" json:"integral_amount,omitempty"`
	// 小数
	FractionalAmount int32 `protobuf:"varint,2,opt,name=fractional_amount,json=fractionalAmount,proto3" json:"fractional_amount,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amount_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_amount_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_amount_proto_rawDescGZIP(), []int{3}
}

func (x *Amount) GetIntegralAmount() int64 {
	if x != nil {
		return x.IntegralAmount
	}
	return 0
}

func (x *Amount) GetFractionalAmount() int32 {
	if x != nil {
		return x.FractionalAmount
	}
	return 0
}

var File_amount_proto protoreflect.FileDescriptor

var file_amount_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x1a, 0x14, 0x6d, 0x74, 0x6e, 0x2f, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x6e, 0x61, 0x6d, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x65, 0x74, 0x61, 0x72, 0x79, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02,
	0x08, 0x03, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x2f, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08,
	0x03, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x33, 0x0a, 0x11, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xc2, 0xb8,
	0x02, 0x02, 0x08, 0x03, 0x52, 0x10, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69,
	0x62, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x55, 0x6e, 0x69, 0x74, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x12, 0x2f, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2, 0xb8, 0x02,
	0x02, 0x08, 0x03, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x11, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06,
	0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x10, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5e, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x45,
	0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x45, 0x4e, 0x54, 0x49, 0x5f, 0x4d,
	0x45, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x5f,
	0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x49, 0x43, 0x52, 0x4f,
	0x5f, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x41, 0x4e, 0x4f,
	0x5f, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x05, 0x22, 0xa8, 0x01, 0x0a, 0x0e, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x6c, 0x69, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x2f, 0x0a,
	0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33,
	0x0a, 0x11, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08,
	0x03, 0x52, 0x10, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x6e, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a,
	0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33,
	0x0a, 0x11, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08,
	0x03, 0x52, 0x10, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_amount_proto_rawDescOnce sync.Once
	file_amount_proto_rawDescData = file_amount_proto_rawDesc
)

func file_amount_proto_rawDescGZIP() []byte {
	file_amount_proto_rawDescOnce.Do(func() {
		file_amount_proto_rawDescData = protoimpl.X.CompressGZIP(file_amount_proto_rawDescData)
	})
	return file_amount_proto_rawDescData
}

var file_amount_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_amount_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_amount_proto_goTypes = []interface{}{
	(DistantAmount_Unit)(0), // 0: sharelib.DistantAmount.Unit
	(*MonetaryAmount)(nil),  // 1: sharelib.MonetaryAmount
	(*DistantAmount)(nil),   // 2: sharelib.DistantAmount
	(*QuantityAmount)(nil),  // 3: sharelib.QuantityAmount
	(*Amount)(nil),          // 4: sharelib.Amount
	(*NameOption)(nil),      // 5: sharelib.NameOption
}
var file_amount_proto_depIdxs = []int32{
	0, // 0: sharelib.DistantAmount.unit:type_name -> sharelib.DistantAmount.Unit
	5, // 1: sharelib.QuantityAmount.unit:type_name -> sharelib.NameOption
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_amount_proto_init() }
func file_amount_proto_init() {
	if File_amount_proto != nil {
		return
	}
	file_nameoption_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_amount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonetaryAmount); i {
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
		file_amount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistantAmount); i {
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
		file_amount_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantityAmount); i {
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
		file_amount_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
			RawDescriptor: file_amount_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_amount_proto_goTypes,
		DependencyIndexes: file_amount_proto_depIdxs,
		EnumInfos:         file_amount_proto_enumTypes,
		MessageInfos:      file_amount_proto_msgTypes,
	}.Build()
	File_amount_proto = out.File
	file_amount_proto_rawDesc = nil
	file_amount_proto_goTypes = nil
	file_amount_proto_depIdxs = nil
}
