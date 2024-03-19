// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: batch_task.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BatchTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// サービス名
	// protoファイルで「service」で定義されているものを指定
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// メソッド名
	// protoファイルで「service」で定義されているものを指定
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// メソッドのリクエストのjson
	// protojsonパッケージで対象Requestのデータを
	// Marshalしたjsonデータ
	ReqeustJson string `protobuf:"bytes,3,opt,name=reqeust_json,json=reqeustJson,proto3" json:"reqeust_json,omitempty"`
}

func (x *BatchTask) Reset() {
	*x = BatchTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchTask) ProtoMessage() {}

func (x *BatchTask) ProtoReflect() protoreflect.Message {
	mi := &file_batch_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchTask.ProtoReflect.Descriptor instead.
func (*BatchTask) Descriptor() ([]byte, []int) {
	return file_batch_task_proto_rawDescGZIP(), []int{0}
}

func (x *BatchTask) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *BatchTask) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *BatchTask) GetReqeustJson() string {
	if x != nil {
		return x.ReqeustJson
	}
	return ""
}

var File_batch_task_proto protoreflect.FileDescriptor

var file_batch_task_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x60, 0x0a, 0x09, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x65, 0x75, 0x73,
	0x74, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x71, 0x65, 0x75, 0x73, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x74, 0x65,
	0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_batch_task_proto_rawDescOnce sync.Once
	file_batch_task_proto_rawDescData = file_batch_task_proto_rawDesc
)

func file_batch_task_proto_rawDescGZIP() []byte {
	file_batch_task_proto_rawDescOnce.Do(func() {
		file_batch_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_batch_task_proto_rawDescData)
	})
	return file_batch_task_proto_rawDescData
}

var file_batch_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_batch_task_proto_goTypes = []interface{}{
	(*BatchTask)(nil), // 0: mtechnavi.api.batch.BatchTask
}
var file_batch_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_batch_task_proto_init() }
func file_batch_task_proto_init() {
	if File_batch_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_batch_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchTask); i {
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
			RawDescriptor: file_batch_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_batch_task_proto_goTypes,
		DependencyIndexes: file_batch_task_proto_depIdxs,
		MessageInfos:      file_batch_task_proto_msgTypes,
	}.Build()
	File_batch_task_proto = out.File
	file_batch_task_proto_rawDesc = nil
	file_batch_task_proto_goTypes = nil
	file_batch_task_proto_depIdxs = nil
}
