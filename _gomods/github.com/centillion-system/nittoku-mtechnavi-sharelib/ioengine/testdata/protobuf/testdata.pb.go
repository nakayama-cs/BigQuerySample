// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: testdata.proto

package testdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	protobuf "mtechnavi/sharelib/protobuf"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestEnum int32

const (
	TestEnum_A TestEnum = 0
	TestEnum_B TestEnum = 1
	TestEnum_C TestEnum = 2
)

// Enum value maps for TestEnum.
var (
	TestEnum_name = map[int32]string{
		0: "A",
		1: "B",
		2: "C",
	}
	TestEnum_value = map[string]int32{
		"A": 0,
		"B": 1,
		"C": 2,
	}
)

func (x TestEnum) Enum() *TestEnum {
	p := new(TestEnum)
	*p = x
	return p
}

func (x TestEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_testdata_proto_enumTypes[0].Descriptor()
}

func (TestEnum) Type() protoreflect.EnumType {
	return &file_testdata_proto_enumTypes[0]
}

func (x TestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestEnum.Descriptor instead.
func (TestEnum) EnumDescriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{0}
}

type TestRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Double_   float64 `protobuf:"fixed64,1,opt,name=double_,json=double,proto3" json:"double_,omitempty"`
	Float_    float32 `protobuf:"fixed32,3,opt,name=float_,json=float,proto3" json:"float_,omitempty"`
	Int32_    int32   `protobuf:"varint,5,opt,name=int32_,json=int32,proto3" json:"int32_,omitempty"`
	Int64_    int64   `protobuf:"varint,7,opt,name=int64_,json=int64,proto3" json:"int64_,omitempty"`
	Uint32_   uint32  `protobuf:"varint,9,opt,name=uint32_,json=uint32,proto3" json:"uint32_,omitempty"`
	Uint64_   uint64  `protobuf:"varint,11,opt,name=uint64_,json=uint64,proto3" json:"uint64_,omitempty"`
	Sint32_   int32   `protobuf:"zigzag32,13,opt,name=sint32_,json=sint32,proto3" json:"sint32_,omitempty"`
	Sint64_   int64   `protobuf:"zigzag64,15,opt,name=sint64_,json=sint64,proto3" json:"sint64_,omitempty"`
	Fixed32_  uint32  `protobuf:"fixed32,17,opt,name=fixed32_,json=fixed32,proto3" json:"fixed32_,omitempty"`
	Fixed64_  uint64  `protobuf:"fixed64,19,opt,name=fixed64_,json=fixed64,proto3" json:"fixed64_,omitempty"`
	Sfixed32_ int32   `protobuf:"fixed32,21,opt,name=sfixed32_,json=sfixed32,proto3" json:"sfixed32_,omitempty"`
	Sfixed64_ int64   `protobuf:"fixed64,23,opt,name=sfixed64_,json=sfixed64,proto3" json:"sfixed64_,omitempty"`
	Bool_     bool    `protobuf:"varint,25,opt,name=bool_,json=bool,proto3" json:"bool_,omitempty"`
	String_   string  `protobuf:"bytes,27,opt,name=string_,json=string,proto3" json:"string_,omitempty"`
	// RootEnum        enum_            = 29;
	RepeatedString_ []string `protobuf:"bytes,31,rep,name=repeated_string_,json=repeatedString,proto3" json:"repeated_string_,omitempty"`
	RepeatedInt64_  []int64  `protobuf:"varint,33,rep,packed,name=repeated_int64_,json=repeatedInt64,proto3" json:"repeated_int64_,omitempty"`
	DoubleBase_     float64  `protobuf:"fixed64,35,opt,name=double_base_,json=doubleBase,proto3" json:"double_base_,omitempty"`
	FloatBase_      float32  `protobuf:"fixed32,37,opt,name=float_base_,json=floatBase,proto3" json:"float_base_,omitempty"`
	Int32Base_      int32    `protobuf:"varint,39,opt,name=int32_base_,json=int32Base,proto3" json:"int32_base_,omitempty"`
	Int64Base_      int64    `protobuf:"varint,41,opt,name=int64_base_,json=int64Base,proto3" json:"int64_base_,omitempty"`
	Uint32Base_     uint32   `protobuf:"varint,43,opt,name=uint32_base_,json=uint32Base,proto3" json:"uint32_base_,omitempty"`
	Uint64Base_     uint64   `protobuf:"varint,45,opt,name=uint64_base_,json=uint64Base,proto3" json:"uint64_base_,omitempty"`
	Sint32Base_     int32    `protobuf:"zigzag32,47,opt,name=sint32_base_,json=sint32Base,proto3" json:"sint32_base_,omitempty"`
	Sint64Base_     int64    `protobuf:"zigzag64,49,opt,name=sint64_base_,json=sint64Base,proto3" json:"sint64_base_,omitempty"`
	Fixed32Base_    uint32   `protobuf:"fixed32,51,opt,name=fixed32_base_,json=fixed32Base,proto3" json:"fixed32_base_,omitempty"`
	Fixed64Base_    uint64   `protobuf:"fixed64,53,opt,name=fixed64_base_,json=fixed64Base,proto3" json:"fixed64_base_,omitempty"`
	Sfixed32Base_   int32    `protobuf:"fixed32,55,opt,name=sfixed32_base_,json=sfixed32Base,proto3" json:"sfixed32_base_,omitempty"`
	Sfixed64Base_   int64    `protobuf:"fixed64,56,opt,name=sfixed64_base_,json=sfixed64Base,proto3" json:"sfixed64_base_,omitempty"`
	BoolBase_       bool     `protobuf:"varint,58,opt,name=bool_base_,json=boolBase,proto3" json:"bool_base_,omitempty"`
	StringBase_     string   `protobuf:"bytes,61,opt,name=string_base_,json=stringBase,proto3" json:"string_base_,omitempty"`
	// RootEnum        enum_            = 29;
	RepeatedStringBase_ []string `protobuf:"bytes,63,rep,name=repeated_string_base_,json=repeatedStringBase,proto3" json:"repeated_string_base_,omitempty"`
	RepeatedInt64Base_  []int64  `protobuf:"varint,65,rep,packed,name=repeated_int64_base_,json=repeatedInt64Base,proto3" json:"repeated_int64_base_,omitempty"`
	TestRecordId        string   `protobuf:"bytes,100,opt,name=test_record_id,json=testRecordId,proto3" json:"test_record_id,omitempty"`
	DeletedAt           int64    `protobuf:"varint,101,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	CreatedAt           int64    `protobuf:"varint,102,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt           int64    `protobuf:"varint,103,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TestRecord) Reset() {
	*x = TestRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRecord) ProtoMessage() {}

func (x *TestRecord) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRecord.ProtoReflect.Descriptor instead.
func (*TestRecord) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{0}
}

func (x *TestRecord) GetDouble_() float64 {
	if x != nil {
		return x.Double_
	}
	return 0
}

func (x *TestRecord) GetFloat_() float32 {
	if x != nil {
		return x.Float_
	}
	return 0
}

func (x *TestRecord) GetInt32_() int32 {
	if x != nil {
		return x.Int32_
	}
	return 0
}

func (x *TestRecord) GetInt64_() int64 {
	if x != nil {
		return x.Int64_
	}
	return 0
}

func (x *TestRecord) GetUint32_() uint32 {
	if x != nil {
		return x.Uint32_
	}
	return 0
}

func (x *TestRecord) GetUint64_() uint64 {
	if x != nil {
		return x.Uint64_
	}
	return 0
}

func (x *TestRecord) GetSint32_() int32 {
	if x != nil {
		return x.Sint32_
	}
	return 0
}

func (x *TestRecord) GetSint64_() int64 {
	if x != nil {
		return x.Sint64_
	}
	return 0
}

func (x *TestRecord) GetFixed32_() uint32 {
	if x != nil {
		return x.Fixed32_
	}
	return 0
}

func (x *TestRecord) GetFixed64_() uint64 {
	if x != nil {
		return x.Fixed64_
	}
	return 0
}

func (x *TestRecord) GetSfixed32_() int32 {
	if x != nil {
		return x.Sfixed32_
	}
	return 0
}

func (x *TestRecord) GetSfixed64_() int64 {
	if x != nil {
		return x.Sfixed64_
	}
	return 0
}

func (x *TestRecord) GetBool_() bool {
	if x != nil {
		return x.Bool_
	}
	return false
}

func (x *TestRecord) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *TestRecord) GetRepeatedString_() []string {
	if x != nil {
		return x.RepeatedString_
	}
	return nil
}

func (x *TestRecord) GetRepeatedInt64_() []int64 {
	if x != nil {
		return x.RepeatedInt64_
	}
	return nil
}

func (x *TestRecord) GetDoubleBase_() float64 {
	if x != nil {
		return x.DoubleBase_
	}
	return 0
}

func (x *TestRecord) GetFloatBase_() float32 {
	if x != nil {
		return x.FloatBase_
	}
	return 0
}

func (x *TestRecord) GetInt32Base_() int32 {
	if x != nil {
		return x.Int32Base_
	}
	return 0
}

func (x *TestRecord) GetInt64Base_() int64 {
	if x != nil {
		return x.Int64Base_
	}
	return 0
}

func (x *TestRecord) GetUint32Base_() uint32 {
	if x != nil {
		return x.Uint32Base_
	}
	return 0
}

func (x *TestRecord) GetUint64Base_() uint64 {
	if x != nil {
		return x.Uint64Base_
	}
	return 0
}

func (x *TestRecord) GetSint32Base_() int32 {
	if x != nil {
		return x.Sint32Base_
	}
	return 0
}

func (x *TestRecord) GetSint64Base_() int64 {
	if x != nil {
		return x.Sint64Base_
	}
	return 0
}

func (x *TestRecord) GetFixed32Base_() uint32 {
	if x != nil {
		return x.Fixed32Base_
	}
	return 0
}

func (x *TestRecord) GetFixed64Base_() uint64 {
	if x != nil {
		return x.Fixed64Base_
	}
	return 0
}

func (x *TestRecord) GetSfixed32Base_() int32 {
	if x != nil {
		return x.Sfixed32Base_
	}
	return 0
}

func (x *TestRecord) GetSfixed64Base_() int64 {
	if x != nil {
		return x.Sfixed64Base_
	}
	return 0
}

func (x *TestRecord) GetBoolBase_() bool {
	if x != nil {
		return x.BoolBase_
	}
	return false
}

func (x *TestRecord) GetStringBase_() string {
	if x != nil {
		return x.StringBase_
	}
	return ""
}

func (x *TestRecord) GetRepeatedStringBase_() []string {
	if x != nil {
		return x.RepeatedStringBase_
	}
	return nil
}

func (x *TestRecord) GetRepeatedInt64Base_() []int64 {
	if x != nil {
		return x.RepeatedInt64Base_
	}
	return nil
}

func (x *TestRecord) GetTestRecordId() string {
	if x != nil {
		return x.TestRecordId
	}
	return ""
}

func (x *TestRecord) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *TestRecord) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TestRecord) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type DatetimeRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datetime1_ *protobuf.Datetime `protobuf:"bytes,1,opt,name=datetime1_,json=datetime1,proto3" json:"datetime1_,omitempty"`
	Datetime3_ *protobuf.Datetime `protobuf:"bytes,3,opt,name=datetime3_,json=datetime3,proto3" json:"datetime3_,omitempty"`
	Datetime5_ *protobuf.Datetime `protobuf:"bytes,5,opt,name=datetime5_,json=datetime5,proto3" json:"datetime5_,omitempty"`
	Datetime7_ *protobuf.Datetime `protobuf:"bytes,7,opt,name=datetime7_,json=datetime7,proto3" json:"datetime7_,omitempty"`
	Datetime9_ *protobuf.Datetime `protobuf:"bytes,9,opt,name=datetime9_,json=datetime9,proto3" json:"datetime9_,omitempty"`
}

func (x *DatetimeRecord) Reset() {
	*x = DatetimeRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatetimeRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatetimeRecord) ProtoMessage() {}

func (x *DatetimeRecord) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatetimeRecord.ProtoReflect.Descriptor instead.
func (*DatetimeRecord) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{1}
}

func (x *DatetimeRecord) GetDatetime1_() *protobuf.Datetime {
	if x != nil {
		return x.Datetime1_
	}
	return nil
}

func (x *DatetimeRecord) GetDatetime3_() *protobuf.Datetime {
	if x != nil {
		return x.Datetime3_
	}
	return nil
}

func (x *DatetimeRecord) GetDatetime5_() *protobuf.Datetime {
	if x != nil {
		return x.Datetime5_
	}
	return nil
}

func (x *DatetimeRecord) GetDatetime7_() *protobuf.Datetime {
	if x != nil {
		return x.Datetime7_
	}
	return nil
}

func (x *DatetimeRecord) GetDatetime9_() *protobuf.Datetime {
	if x != nil {
		return x.Datetime9_
	}
	return nil
}

var File_testdata_proto protoreflect.FileDescriptor

var file_testdata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x09, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x5f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x73,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x5f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x08,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x5f, 0x18, 0x17, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x13, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a,
	0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f,
	0x18, 0x21, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x23, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x25, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x42, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x27, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x42, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x29, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x42, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x75,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x42, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x73,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x2f, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x0a, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x0c, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x31, 0x20,
	0x01, 0x28, 0x12, 0x52, 0x0a, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0b, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x42,
	0x61, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x18, 0x35, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0b, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x42, 0x61, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x37, 0x20, 0x01, 0x28, 0x0f, 0x52,
	0x0c, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x42, 0x61, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18,
	0x38, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0c, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x42,
	0x61, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x42, 0x61, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42,
	0x61, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x3f, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x42, 0x61, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x41,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x42, 0x61, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8f, 0x02, 0x0a, 0x0e, 0x44,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x31, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x31, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x44, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x31,
	0x12, 0x31, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x33, 0x5f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x33, 0x12, 0x31, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x35,
	0x5f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c,
	0x69, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x35, 0x12, 0x31, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x37, 0x5f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x09,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x37, 0x12, 0x31, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x39, 0x5f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x39, 0x2a, 0x1f, 0x0a, 0x08,
	0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x12,
	0x05, 0x0a, 0x01, 0x42, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x43, 0x10, 0x02, 0x42, 0x14, 0x5a,
	0x12, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_proto_rawDescOnce sync.Once
	file_testdata_proto_rawDescData = file_testdata_proto_rawDesc
)

func file_testdata_proto_rawDescGZIP() []byte {
	file_testdata_proto_rawDescOnce.Do(func() {
		file_testdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_proto_rawDescData)
	})
	return file_testdata_proto_rawDescData
}

var file_testdata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_testdata_proto_goTypes = []interface{}{
	(TestEnum)(0),             // 0: mtechnavi.testdata.TestEnum
	(*TestRecord)(nil),        // 1: mtechnavi.testdata.TestRecord
	(*DatetimeRecord)(nil),    // 2: mtechnavi.testdata.DatetimeRecord
	(*protobuf.Datetime)(nil), // 3: sharelib.Datetime
}
var file_testdata_proto_depIdxs = []int32{
	3, // 0: mtechnavi.testdata.DatetimeRecord.datetime1_:type_name -> sharelib.Datetime
	3, // 1: mtechnavi.testdata.DatetimeRecord.datetime3_:type_name -> sharelib.Datetime
	3, // 2: mtechnavi.testdata.DatetimeRecord.datetime5_:type_name -> sharelib.Datetime
	3, // 3: mtechnavi.testdata.DatetimeRecord.datetime7_:type_name -> sharelib.Datetime
	3, // 4: mtechnavi.testdata.DatetimeRecord.datetime9_:type_name -> sharelib.Datetime
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_testdata_proto_init() }
func file_testdata_proto_init() {
	if File_testdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRecord); i {
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
		file_testdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatetimeRecord); i {
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
			RawDescriptor: file_testdata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_proto_goTypes,
		DependencyIndexes: file_testdata_proto_depIdxs,
		EnumInfos:         file_testdata_proto_enumTypes,
		MessageInfos:      file_testdata_proto_msgTypes,
	}.Build()
	File_testdata_proto = out.File
	file_testdata_proto_rawDesc = nil
	file_testdata_proto_goTypes = nil
	file_testdata_proto_depIdxs = nil
}
