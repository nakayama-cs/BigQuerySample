// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: worktaskrelation.proto

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

// 作業案件紐付けデータ
type WorkTaskRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 作業案件紐付けID
	WorkTaskRelationId string `protobuf:"bytes,1,opt,name=work_task_relation_id,json=workTaskRelationId,proto3" json:"work_task_relation_id,omitempty"`
	// 親指図案件ID
	LinkedWorkTaskId string `protobuf:"bytes,2,opt,name=linked_work_task_id,json=linkedWorkTaskId,proto3" json:"linked_work_task_id,omitempty"`
	// readonly 子伝票urn
	Urn string `protobuf:"bytes,3,opt,name=urn,proto3" json:"urn,omitempty"`
	// type_name
	TypeName string `protobuf:"bytes,4,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// record_id
	RecordId string `protobuf:"bytes,5,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	// システム分類
	// 名称マスタにおける種別 A2000002 (システム分類) を参照
	SystemCategory *protobuf.NameOption `protobuf:"bytes,6,opt,name=system_category,json=systemCategory,proto3" json:"system_category,omitempty"`
	// created_at
	CreatedAt int64 `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// updated_at
	UpdatedAt int64 `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// deleted_at
	DeletedAt int64 `protobuf:"varint,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *WorkTaskRelation) Reset() {
	*x = WorkTaskRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worktaskrelation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkTaskRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkTaskRelation) ProtoMessage() {}

func (x *WorkTaskRelation) ProtoReflect() protoreflect.Message {
	mi := &file_worktaskrelation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkTaskRelation.ProtoReflect.Descriptor instead.
func (*WorkTaskRelation) Descriptor() ([]byte, []int) {
	return file_worktaskrelation_proto_rawDescGZIP(), []int{0}
}

func (x *WorkTaskRelation) GetWorkTaskRelationId() string {
	if x != nil {
		return x.WorkTaskRelationId
	}
	return ""
}

func (x *WorkTaskRelation) GetLinkedWorkTaskId() string {
	if x != nil {
		return x.LinkedWorkTaskId
	}
	return ""
}

func (x *WorkTaskRelation) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *WorkTaskRelation) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *WorkTaskRelation) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *WorkTaskRelation) GetSystemCategory() *protobuf.NameOption {
	if x != nil {
		return x.SystemCategory
	}
	return nil
}

func (x *WorkTaskRelation) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *WorkTaskRelation) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *WorkTaskRelation) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

// 指図案件紐付けコンテンツ
type WorkTaskRelationContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 指図案件紐付けデータ
	WorkTaskRelation *WorkTaskRelation `protobuf:"bytes,1,opt,name=work_task_relation,json=workTaskRelation,proto3" json:"work_task_relation,omitempty"`
	// 番号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// 件名
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// ステータス
	// 名称マスタにおける種別 A0000040 (完了ステータス) を参照
	Status *protobuf.NameOption `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *WorkTaskRelationContent) Reset() {
	*x = WorkTaskRelationContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worktaskrelation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkTaskRelationContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkTaskRelationContent) ProtoMessage() {}

func (x *WorkTaskRelationContent) ProtoReflect() protoreflect.Message {
	mi := &file_worktaskrelation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkTaskRelationContent.ProtoReflect.Descriptor instead.
func (*WorkTaskRelationContent) Descriptor() ([]byte, []int) {
	return file_worktaskrelation_proto_rawDescGZIP(), []int{1}
}

func (x *WorkTaskRelationContent) GetWorkTaskRelation() *WorkTaskRelation {
	if x != nil {
		return x.WorkTaskRelation
	}
	return nil
}

func (x *WorkTaskRelationContent) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *WorkTaskRelationContent) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *WorkTaskRelationContent) GetStatus() *protobuf.NameOption {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_worktaskrelation_proto protoreflect.FileDescriptor

var file_worktaskrelation_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61, 0x73, 0x6b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61, 0x73, 0x6b,
	0x1a, 0x0f, 0x6d, 0x74, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x99, 0x03, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x15, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x13, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xd2, 0xb8, 0x02, 0x17, 0x0a, 0x15,
	0x0a, 0x13, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x52, 0x10, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x57, 0x6f, 0x72,
	0x6b, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xd2, 0xb8, 0x02, 0x07, 0x0a, 0x05, 0x0a, 0x03, 0x75, 0x72,
	0x6e, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xd2, 0xb8, 0x02, 0x0d, 0x0a,
	0x0b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xd6, 0x01, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x12, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x1d, 0x5a, 0x1b, 0x6d, 0x74, 0x65, 0x63,
	0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worktaskrelation_proto_rawDescOnce sync.Once
	file_worktaskrelation_proto_rawDescData = file_worktaskrelation_proto_rawDesc
)

func file_worktaskrelation_proto_rawDescGZIP() []byte {
	file_worktaskrelation_proto_rawDescOnce.Do(func() {
		file_worktaskrelation_proto_rawDescData = protoimpl.X.CompressGZIP(file_worktaskrelation_proto_rawDescData)
	})
	return file_worktaskrelation_proto_rawDescData
}

var file_worktaskrelation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_worktaskrelation_proto_goTypes = []interface{}{
	(*WorkTaskRelation)(nil),        // 0: mtechnavi.api.worktask.WorkTaskRelation
	(*WorkTaskRelationContent)(nil), // 1: mtechnavi.api.worktask.WorkTaskRelationContent
	(*protobuf.NameOption)(nil),     // 2: sharelib.NameOption
}
var file_worktaskrelation_proto_depIdxs = []int32{
	2, // 0: mtechnavi.api.worktask.WorkTaskRelation.system_category:type_name -> sharelib.NameOption
	0, // 1: mtechnavi.api.worktask.WorkTaskRelationContent.work_task_relation:type_name -> mtechnavi.api.worktask.WorkTaskRelation
	2, // 2: mtechnavi.api.worktask.WorkTaskRelationContent.status:type_name -> sharelib.NameOption
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_worktaskrelation_proto_init() }
func file_worktaskrelation_proto_init() {
	if File_worktaskrelation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worktaskrelation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkTaskRelation); i {
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
		file_worktaskrelation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkTaskRelationContent); i {
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
			RawDescriptor: file_worktaskrelation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_worktaskrelation_proto_goTypes,
		DependencyIndexes: file_worktaskrelation_proto_depIdxs,
		MessageInfos:      file_worktaskrelation_proto_msgTypes,
	}.Build()
	File_worktaskrelation_proto = out.File
	file_worktaskrelation_proto_rawDesc = nil
	file_worktaskrelation_proto_goTypes = nil
	file_worktaskrelation_proto_depIdxs = nil
}
