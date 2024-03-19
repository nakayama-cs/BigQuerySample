// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: batch.proto

package protobuf

import (
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

type ScheduleJobDailyMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 処理対象開始時間
	// e.g. HHMMSS
	Since string `protobuf:"bytes,1,opt,name=since,proto3" json:"since,omitempty"`
	// 処理対象終了時間
	// e.g. HHMMSS
	Until string `protobuf:"bytes,2,opt,name=until,proto3" json:"until,omitempty"`
}

func (x *ScheduleJobDailyMailRequest) Reset() {
	*x = ScheduleJobDailyMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleJobDailyMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleJobDailyMailRequest) ProtoMessage() {}

func (x *ScheduleJobDailyMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleJobDailyMailRequest.ProtoReflect.Descriptor instead.
func (*ScheduleJobDailyMailRequest) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleJobDailyMailRequest) GetSince() string {
	if x != nil {
		return x.Since
	}
	return ""
}

func (x *ScheduleJobDailyMailRequest) GetUntil() string {
	if x != nil {
		return x.Until
	}
	return ""
}

type ScheduleJobSharetoSurveyReminderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScheduleJobSharetoSurveyReminderRequest) Reset() {
	*x = ScheduleJobSharetoSurveyReminderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleJobSharetoSurveyReminderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleJobSharetoSurveyReminderRequest) ProtoMessage() {}

func (x *ScheduleJobSharetoSurveyReminderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleJobSharetoSurveyReminderRequest.ProtoReflect.Descriptor instead.
func (*ScheduleJobSharetoSurveyReminderRequest) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{1}
}

type ScheduleJobSharetoBusinessUnitManagementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScheduleJobSharetoBusinessUnitManagementRequest) Reset() {
	*x = ScheduleJobSharetoBusinessUnitManagementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleJobSharetoBusinessUnitManagementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleJobSharetoBusinessUnitManagementRequest) ProtoMessage() {}

func (x *ScheduleJobSharetoBusinessUnitManagementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleJobSharetoBusinessUnitManagementRequest.ProtoReflect.Descriptor instead.
func (*ScheduleJobSharetoBusinessUnitManagementRequest) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{2}
}

var File_batch_proto protoreflect.FileDescriptor

var file_batch_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x49, 0x0a, 0x1b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x44, 0x61,
	0x69, 0x6c, 0x79, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x69, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x22, 0x29, 0x0a, 0x27, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x53, 0x68, 0x61, 0x72, 0x65, 0x74, 0x6f,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x2f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x4a, 0x6f, 0x62, 0x53, 0x68, 0x61, 0x72, 0x65, 0x74, 0x6f, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xf5, 0x02, 0x0a, 0x0c, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x14, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x4d, 0x61, 0x69,
	0x6c, 0x12, 0x30, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x4a, 0x6f, 0x62, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x78, 0x0a, 0x20, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x53, 0x68, 0x61, 0x72, 0x65, 0x74,
	0x6f, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x3c, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4a, 0x6f,
	0x62, 0x53, 0x68, 0x61, 0x72, 0x65, 0x74, 0x6f, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x88, 0x01, 0x0a, 0x28, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x53, 0x68, 0x61, 0x72, 0x65, 0x74, 0x6f, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x44, 0x2e, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x4a, 0x6f, 0x62, 0x53, 0x68, 0x61, 0x72, 0x65, 0x74, 0x6f, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x61, 0x76, 0x69, 0x2f, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_batch_proto_rawDescOnce sync.Once
	file_batch_proto_rawDescData = file_batch_proto_rawDesc
)

func file_batch_proto_rawDescGZIP() []byte {
	file_batch_proto_rawDescOnce.Do(func() {
		file_batch_proto_rawDescData = protoimpl.X.CompressGZIP(file_batch_proto_rawDescData)
	})
	return file_batch_proto_rawDescData
}

var file_batch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_batch_proto_goTypes = []interface{}{
	(*ScheduleJobDailyMailRequest)(nil),                     // 0: mtechnavi.api.batch.ScheduleJobDailyMailRequest
	(*ScheduleJobSharetoSurveyReminderRequest)(nil),         // 1: mtechnavi.api.batch.ScheduleJobSharetoSurveyReminderRequest
	(*ScheduleJobSharetoBusinessUnitManagementRequest)(nil), // 2: mtechnavi.api.batch.ScheduleJobSharetoBusinessUnitManagementRequest
	(*emptypb.Empty)(nil),                                   // 3: google.protobuf.Empty
}
var file_batch_proto_depIdxs = []int32{
	0, // 0: mtechnavi.api.batch.BatchService.ScheduleJobDailyMail:input_type -> mtechnavi.api.batch.ScheduleJobDailyMailRequest
	1, // 1: mtechnavi.api.batch.BatchService.ScheduleJobSharetoSurveyReminder:input_type -> mtechnavi.api.batch.ScheduleJobSharetoSurveyReminderRequest
	2, // 2: mtechnavi.api.batch.BatchService.ScheduleJobSharetoBusinessUnitManagement:input_type -> mtechnavi.api.batch.ScheduleJobSharetoBusinessUnitManagementRequest
	3, // 3: mtechnavi.api.batch.BatchService.ScheduleJobDailyMail:output_type -> google.protobuf.Empty
	3, // 4: mtechnavi.api.batch.BatchService.ScheduleJobSharetoSurveyReminder:output_type -> google.protobuf.Empty
	3, // 5: mtechnavi.api.batch.BatchService.ScheduleJobSharetoBusinessUnitManagement:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_batch_proto_init() }
func file_batch_proto_init() {
	if File_batch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_batch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleJobDailyMailRequest); i {
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
		file_batch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleJobSharetoSurveyReminderRequest); i {
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
		file_batch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleJobSharetoBusinessUnitManagementRequest); i {
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
			RawDescriptor: file_batch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_batch_proto_goTypes,
		DependencyIndexes: file_batch_proto_depIdxs,
		MessageInfos:      file_batch_proto_msgTypes,
	}.Build()
	File_batch_proto = out.File
	file_batch_proto_rawDesc = nil
	file_batch_proto_goTypes = nil
	file_batch_proto_depIdxs = nil
}
