// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: embedded.proto

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

// Implementation Note:
// "mtechnavi/sharelib/to".ToEmbeddedUser を用いて生成すること。
type EmbeddedUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *EmbeddedUser) Reset() {
	*x = EmbeddedUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_embedded_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbeddedUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbeddedUser) ProtoMessage() {}

func (x *EmbeddedUser) ProtoReflect() protoreflect.Message {
	mi := &file_embedded_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbeddedUser.ProtoReflect.Descriptor instead.
func (*EmbeddedUser) Descriptor() ([]byte, []int) {
	return file_embedded_proto_rawDescGZIP(), []int{0}
}

func (x *EmbeddedUser) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *EmbeddedUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// 他テナントから共有されたmessageに付与される、共有についての情報を格納するmessage
// read-only
//
// このmessageはDataproxyが設定する。
// 利用側は、直接設定してはならず、参照のみとすること。
type EmbeddedSharedProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 共有元のテナントID
	SharedBy string `protobuf:"bytes,1,opt,name=shared_by,json=sharedBy,proto3" json:"shared_by,omitempty"`
	// 自テナントに共有が実施された時刻
	SharedAt int64 `protobuf:"varint,2,opt,name=shared_at,json=sharedAt,proto3" json:"shared_at,omitempty"`
}

func (x *EmbeddedSharedProperties) Reset() {
	*x = EmbeddedSharedProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_embedded_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbeddedSharedProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbeddedSharedProperties) ProtoMessage() {}

func (x *EmbeddedSharedProperties) ProtoReflect() protoreflect.Message {
	mi := &file_embedded_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbeddedSharedProperties.ProtoReflect.Descriptor instead.
func (*EmbeddedSharedProperties) Descriptor() ([]byte, []int) {
	return file_embedded_proto_rawDescGZIP(), []int{1}
}

func (x *EmbeddedSharedProperties) GetSharedBy() string {
	if x != nil {
		return x.SharedBy
	}
	return ""
}

func (x *EmbeddedSharedProperties) GetSharedAt() int64 {
	if x != nil {
		return x.SharedAt
	}
	return 0
}

// 最終更新情報
type EmbeddedUpdatedProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 最終更新日時
	UpdatedAt int64 `protobuf:"varint,1,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// 最終更新ユーザー
	UpdatedBy *EmbeddedUser `protobuf:"bytes,2,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *EmbeddedUpdatedProperties) Reset() {
	*x = EmbeddedUpdatedProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_embedded_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbeddedUpdatedProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbeddedUpdatedProperties) ProtoMessage() {}

func (x *EmbeddedUpdatedProperties) ProtoReflect() protoreflect.Message {
	mi := &file_embedded_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbeddedUpdatedProperties.ProtoReflect.Descriptor instead.
func (*EmbeddedUpdatedProperties) Descriptor() ([]byte, []int) {
	return file_embedded_proto_rawDescGZIP(), []int{2}
}

func (x *EmbeddedUpdatedProperties) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *EmbeddedUpdatedProperties) GetUpdatedBy() *EmbeddedUser {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

// 企業情報
type EmbeddedCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 企業ID
	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// 企業名
	// ja:企業名
	// ja-kana:ふりがな
	// en:英語名
	DisplayNameLang map[string]string `protobuf:"bytes,2,rep,name=display_name_lang,json=displayNameLang,proto3" json:"display_name_lang,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EmbeddedCompany) Reset() {
	*x = EmbeddedCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_embedded_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbeddedCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbeddedCompany) ProtoMessage() {}

func (x *EmbeddedCompany) ProtoReflect() protoreflect.Message {
	mi := &file_embedded_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbeddedCompany.ProtoReflect.Descriptor instead.
func (*EmbeddedCompany) Descriptor() ([]byte, []int) {
	return file_embedded_proto_rawDescGZIP(), []int{3}
}

func (x *EmbeddedCompany) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *EmbeddedCompany) GetDisplayNameLang() map[string]string {
	if x != nil {
		return x.DisplayNameLang
	}
	return nil
}

var File_embedded_proto protoreflect.FileDescriptor

var file_embedded_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x1a, 0x14, 0x6d, 0x74, 0x6e, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x57, 0x0a, 0x0c, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x29, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02,
	0x08, 0x03, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x54, 0x0a, 0x18, 0x45, 0x6d, 0x62,
	0x65, 0x64, 0x64, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x71, 0x0a, 0x19, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x22, 0xe0, 0x01, 0x0a, 0x0f, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x25, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02,
	0x08, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x62, 0x0a,
	0x11, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x61,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x6c, 0x69, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4c,
	0x61, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x06, 0xc2, 0xb8, 0x02, 0x02, 0x08, 0x03,
	0x52, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x61, 0x6e,
	0x67, 0x1a, 0x42, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x4c, 0x61, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61,
	0x76, 0x69, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_embedded_proto_rawDescOnce sync.Once
	file_embedded_proto_rawDescData = file_embedded_proto_rawDesc
)

func file_embedded_proto_rawDescGZIP() []byte {
	file_embedded_proto_rawDescOnce.Do(func() {
		file_embedded_proto_rawDescData = protoimpl.X.CompressGZIP(file_embedded_proto_rawDescData)
	})
	return file_embedded_proto_rawDescData
}

var file_embedded_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_embedded_proto_goTypes = []interface{}{
	(*EmbeddedUser)(nil),              // 0: sharelib.EmbeddedUser
	(*EmbeddedSharedProperties)(nil),  // 1: sharelib.EmbeddedSharedProperties
	(*EmbeddedUpdatedProperties)(nil), // 2: sharelib.EmbeddedUpdatedProperties
	(*EmbeddedCompany)(nil),           // 3: sharelib.EmbeddedCompany
	nil,                               // 4: sharelib.EmbeddedCompany.DisplayNameLangEntry
}
var file_embedded_proto_depIdxs = []int32{
	0, // 0: sharelib.EmbeddedUpdatedProperties.updated_by:type_name -> sharelib.EmbeddedUser
	4, // 1: sharelib.EmbeddedCompany.display_name_lang:type_name -> sharelib.EmbeddedCompany.DisplayNameLangEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_embedded_proto_init() }
func file_embedded_proto_init() {
	if File_embedded_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_embedded_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbeddedUser); i {
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
		file_embedded_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbeddedSharedProperties); i {
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
		file_embedded_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbeddedUpdatedProperties); i {
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
		file_embedded_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbeddedCompany); i {
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
			RawDescriptor: file_embedded_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_embedded_proto_goTypes,
		DependencyIndexes: file_embedded_proto_depIdxs,
		MessageInfos:      file_embedded_proto_msgTypes,
	}.Build()
	File_embedded_proto = out.File
	file_embedded_proto_rawDesc = nil
	file_embedded_proto_goTypes = nil
	file_embedded_proto_depIdxs = nil
}
