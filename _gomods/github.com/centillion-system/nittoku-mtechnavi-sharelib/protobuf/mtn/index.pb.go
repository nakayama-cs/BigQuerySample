// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: mtn/index.proto

package mtn

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataproxyIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// インデックス名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DataproxyIndex) Reset() {
	*x = DataproxyIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtn_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataproxyIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataproxyIndex) ProtoMessage() {}

func (x *DataproxyIndex) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataproxyIndex.ProtoReflect.Descriptor instead.
func (*DataproxyIndex) Descriptor() ([]byte, []int) {
	return file_mtn_index_proto_rawDescGZIP(), []int{0}
}

func (x *DataproxyIndex) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Dataproxy利用に対するオプション設定
type DataproxyOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index *DataproxyIndex `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *DataproxyOptions) Reset() {
	*x = DataproxyOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtn_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataproxyOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataproxyOptions) ProtoMessage() {}

func (x *DataproxyOptions) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataproxyOptions.ProtoReflect.Descriptor instead.
func (*DataproxyOptions) Descriptor() ([]byte, []int) {
	return file_mtn_index_proto_rawDescGZIP(), []int{1}
}

func (x *DataproxyOptions) GetIndex() *DataproxyIndex {
	if x != nil {
		return x.Index
	}
	return nil
}

var file_mtn_index_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*DataproxyOptions)(nil),
		Field:         5002,
		Name:          "mtn.dataproxy",
		Tag:           "bytes,5002,opt,name=dataproxy",
		Filename:      "mtn/index.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional mtn.DataproxyOptions dataproxy = 5002;
	E_Dataproxy = &file_mtn_index_proto_extTypes[0]
)

var File_mtn_index_proto protoreflect.FileDescriptor

var file_mtn_index_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x74, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x6d, 0x74, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d,
	0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x74, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x53, 0x0a,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a, 0x27, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6d, 0x74, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x42, 0x21, 0x5a, 0x1f, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x6d, 0x74, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mtn_index_proto_rawDescOnce sync.Once
	file_mtn_index_proto_rawDescData = file_mtn_index_proto_rawDesc
)

func file_mtn_index_proto_rawDescGZIP() []byte {
	file_mtn_index_proto_rawDescOnce.Do(func() {
		file_mtn_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_mtn_index_proto_rawDescData)
	})
	return file_mtn_index_proto_rawDescData
}

var file_mtn_index_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mtn_index_proto_goTypes = []interface{}{
	(*DataproxyIndex)(nil),            // 0: mtn.DataproxyIndex
	(*DataproxyOptions)(nil),          // 1: mtn.DataproxyOptions
	(*descriptorpb.FieldOptions)(nil), // 2: google.protobuf.FieldOptions
}
var file_mtn_index_proto_depIdxs = []int32{
	0, // 0: mtn.DataproxyOptions.index:type_name -> mtn.DataproxyIndex
	2, // 1: mtn.dataproxy:extendee -> google.protobuf.FieldOptions
	1, // 2: mtn.dataproxy:type_name -> mtn.DataproxyOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mtn_index_proto_init() }
func file_mtn_index_proto_init() {
	if File_mtn_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mtn_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataproxyIndex); i {
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
		file_mtn_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataproxyOptions); i {
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
			RawDescriptor: file_mtn_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_mtn_index_proto_goTypes,
		DependencyIndexes: file_mtn_index_proto_depIdxs,
		MessageInfos:      file_mtn_index_proto_msgTypes,
		ExtensionInfos:    file_mtn_index_proto_extTypes,
	}.Build()
	File_mtn_index_proto = out.File
	file_mtn_index_proto_rawDesc = nil
	file_mtn_index_proto_goTypes = nil
	file_mtn_index_proto_depIdxs = nil
}
