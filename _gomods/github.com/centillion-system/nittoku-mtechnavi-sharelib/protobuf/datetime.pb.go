// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: datetime.proto

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

// 精度
// どの精度で取り扱っているかを指定する
// e.g.
// DAYの場合 timestampにはマイクロ秒精度のUnix時刻が格納されるが、時間以下の部分については0クリアした上で扱う
type Datetime_Accuracy int32

const (
	Datetime_MICROSECOND Datetime_Accuracy = 0
	Datetime_MILLISECOND Datetime_Accuracy = 1
	Datetime_SECOND      Datetime_Accuracy = 2
	Datetime_MINUTE      Datetime_Accuracy = 3
	Datetime_HOUR        Datetime_Accuracy = 4
	Datetime_DAY         Datetime_Accuracy = 5
	Datetime_MONTH       Datetime_Accuracy = 6
	Datetime_YEAR        Datetime_Accuracy = 7
)

// Enum value maps for Datetime_Accuracy.
var (
	Datetime_Accuracy_name = map[int32]string{
		0: "MICROSECOND",
		1: "MILLISECOND",
		2: "SECOND",
		3: "MINUTE",
		4: "HOUR",
		5: "DAY",
		6: "MONTH",
		7: "YEAR",
	}
	Datetime_Accuracy_value = map[string]int32{
		"MICROSECOND": 0,
		"MILLISECOND": 1,
		"SECOND":      2,
		"MINUTE":      3,
		"HOUR":        4,
		"DAY":         5,
		"MONTH":       6,
		"YEAR":        7,
	}
)

func (x Datetime_Accuracy) Enum() *Datetime_Accuracy {
	p := new(Datetime_Accuracy)
	*p = x
	return p
}

func (x Datetime_Accuracy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Datetime_Accuracy) Descriptor() protoreflect.EnumDescriptor {
	return file_datetime_proto_enumTypes[0].Descriptor()
}

func (Datetime_Accuracy) Type() protoreflect.EnumType {
	return &file_datetime_proto_enumTypes[0]
}

func (x Datetime_Accuracy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Datetime_Accuracy.Descriptor instead.
func (Datetime_Accuracy) EnumDescriptor() ([]byte, []int) {
	return file_datetime_proto_rawDescGZIP(), []int{0, 0}
}

// 日付型
type Datetime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// タイムゾーン
	// TZ database name の形式で指定する
	// 空の場合 "Asia/Tokyo" とする
	TimezoneName string `protobuf:"bytes,1,opt,name=timezone_name,json=timezoneName,proto3" json:"timezone_name,omitempty"`
	// UNIXエポック マイクロ秒
	Timestamp int64             `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Accuracy  Datetime_Accuracy `protobuf:"varint,3,opt,name=accuracy,proto3,enum=sharelib.Datetime_Accuracy" json:"accuracy,omitempty"`
}

func (x *Datetime) Reset() {
	*x = Datetime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datetime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datetime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datetime) ProtoMessage() {}

func (x *Datetime) ProtoReflect() protoreflect.Message {
	mi := &file_datetime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datetime.ProtoReflect.Descriptor instead.
func (*Datetime) Descriptor() ([]byte, []int) {
	return file_datetime_proto_rawDescGZIP(), []int{0}
}

func (x *Datetime) GetTimezoneName() string {
	if x != nil {
		return x.TimezoneName
	}
	return ""
}

func (x *Datetime) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Datetime) GetAccuracy() Datetime_Accuracy {
	if x != nil {
		return x.Accuracy
	}
	return Datetime_MICROSECOND
}

type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 年
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// 月
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// 日
	Day int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datetime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_datetime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_datetime_proto_rawDescGZIP(), []int{1}
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

var File_datetime_proto protoreflect.FileDescriptor

var file_datetime_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x1a, 0x14, 0x6d, 0x74, 0x6e, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8c, 0x02, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x0c, 0x74, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2,
	0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x3f, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x42,
	0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63,
	0x79, 0x22, 0x6c, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x0f, 0x0a,
	0x0b, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x55, 0x52, 0x10,
	0x04, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41, 0x59, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f,
	0x4e, 0x54, 0x48, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x45, 0x41, 0x52, 0x10, 0x07, 0x22,
	0x5a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x12, 0x18, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06,
	0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x03, 0x64, 0x61, 0x79, 0x42, 0x1d, 0x5a, 0x1b, 0x6d,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_datetime_proto_rawDescOnce sync.Once
	file_datetime_proto_rawDescData = file_datetime_proto_rawDesc
)

func file_datetime_proto_rawDescGZIP() []byte {
	file_datetime_proto_rawDescOnce.Do(func() {
		file_datetime_proto_rawDescData = protoimpl.X.CompressGZIP(file_datetime_proto_rawDescData)
	})
	return file_datetime_proto_rawDescData
}

var file_datetime_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_datetime_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datetime_proto_goTypes = []interface{}{
	(Datetime_Accuracy)(0), // 0: sharelib.Datetime.Accuracy
	(*Datetime)(nil),       // 1: sharelib.Datetime
	(*Date)(nil),           // 2: sharelib.Date
}
var file_datetime_proto_depIdxs = []int32{
	0, // 0: sharelib.Datetime.accuracy:type_name -> sharelib.Datetime.Accuracy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_datetime_proto_init() }
func file_datetime_proto_init() {
	if File_datetime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datetime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datetime); i {
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
		file_datetime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
			RawDescriptor: file_datetime_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datetime_proto_goTypes,
		DependencyIndexes: file_datetime_proto_depIdxs,
		EnumInfos:         file_datetime_proto_enumTypes,
		MessageInfos:      file_datetime_proto_msgTypes,
	}.Build()
	File_datetime_proto = out.File
	file_datetime_proto_rawDesc = nil
	file_datetime_proto_goTypes = nil
	file_datetime_proto_depIdxs = nil
}
