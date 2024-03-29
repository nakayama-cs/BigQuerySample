// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: estimatefeedback.proto

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

// オーナー：発注企業
// 見積結果フィードバック
// e.g. 差戻の処理で利用
type EstimateResultFeedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EstimateResultFeedbackId string `protobuf:"bytes,1,opt,name=estimate_result_feedback_id,json=estimateResultFeedbackId,proto3" json:"estimate_result_feedback_id,omitempty"`
	EstimateResultId         string `protobuf:"bytes,2,opt,name=estimate_result_id,json=estimateResultId,proto3" json:"estimate_result_id,omitempty"`
	// コメント
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	// readonly 送信日
	// フィードバック情報がshareされた日時
	SendedAt int64 `protobuf:"varint,4,opt,name=sended_at,json=sendedAt,proto3" json:"sended_at,omitempty"`
	// readonly 共有情報
	SharedProperties *protobuf.EmbeddedSharedProperties `protobuf:"bytes,5,opt,name=shared_properties,json=sharedProperties,proto3" json:"shared_properties,omitempty"`
	// 共有項目xxx_atのscopeはidと同様
	DeletedAt int64 `protobuf:"varint,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	CreatedAt int64 `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64 `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *EstimateResultFeedback) Reset() {
	*x = EstimateResultFeedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_estimatefeedback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimateResultFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateResultFeedback) ProtoMessage() {}

func (x *EstimateResultFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_estimatefeedback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateResultFeedback.ProtoReflect.Descriptor instead.
func (*EstimateResultFeedback) Descriptor() ([]byte, []int) {
	return file_estimatefeedback_proto_rawDescGZIP(), []int{0}
}

func (x *EstimateResultFeedback) GetEstimateResultFeedbackId() string {
	if x != nil {
		return x.EstimateResultFeedbackId
	}
	return ""
}

func (x *EstimateResultFeedback) GetEstimateResultId() string {
	if x != nil {
		return x.EstimateResultId
	}
	return ""
}

func (x *EstimateResultFeedback) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *EstimateResultFeedback) GetSendedAt() int64 {
	if x != nil {
		return x.SendedAt
	}
	return 0
}

func (x *EstimateResultFeedback) GetSharedProperties() *protobuf.EmbeddedSharedProperties {
	if x != nil {
		return x.SharedProperties
	}
	return nil
}

func (x *EstimateResultFeedback) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *EstimateResultFeedback) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *EstimateResultFeedback) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_estimatefeedback_proto protoreflect.FileDescriptor

var file_estimatefeedback_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x14, 0x6d, 0x74, 0x6e, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6d, 0x74, 0x6e, 0x2f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x03, 0x0a, 0x16, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x3d, 0x0a, 0x1b, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x65, 0x73, 0x74, 0x69, 0x6d,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x12, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x20, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0xd2, 0xb8, 0x02, 0x16, 0x0a, 0x14, 0x0a, 0x12, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69,
	0x64, 0x52, 0x10, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x01,
	0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4f, 0x0a, 0x11, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62,
	0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x10, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x6d, 0x74, 0x65, 0x63,
	0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_estimatefeedback_proto_rawDescOnce sync.Once
	file_estimatefeedback_proto_rawDescData = file_estimatefeedback_proto_rawDesc
)

func file_estimatefeedback_proto_rawDescGZIP() []byte {
	file_estimatefeedback_proto_rawDescOnce.Do(func() {
		file_estimatefeedback_proto_rawDescData = protoimpl.X.CompressGZIP(file_estimatefeedback_proto_rawDescData)
	})
	return file_estimatefeedback_proto_rawDescData
}

var file_estimatefeedback_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_estimatefeedback_proto_goTypes = []interface{}{
	(*EstimateResultFeedback)(nil),            // 0: mtechnavi.api.estimation.EstimateResultFeedback
	(*protobuf.EmbeddedSharedProperties)(nil), // 1: sharelib.EmbeddedSharedProperties
}
var file_estimatefeedback_proto_depIdxs = []int32{
	1, // 0: mtechnavi.api.estimation.EstimateResultFeedback.shared_properties:type_name -> sharelib.EmbeddedSharedProperties
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_estimatefeedback_proto_init() }
func file_estimatefeedback_proto_init() {
	if File_estimatefeedback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_estimatefeedback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimateResultFeedback); i {
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
			RawDescriptor: file_estimatefeedback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_estimatefeedback_proto_goTypes,
		DependencyIndexes: file_estimatefeedback_proto_depIdxs,
		MessageInfos:      file_estimatefeedback_proto_msgTypes,
	}.Build()
	File_estimatefeedback_proto = out.File
	file_estimatefeedback_proto_rawDesc = nil
	file_estimatefeedback_proto_goTypes = nil
	file_estimatefeedback_proto_depIdxs = nil
}
