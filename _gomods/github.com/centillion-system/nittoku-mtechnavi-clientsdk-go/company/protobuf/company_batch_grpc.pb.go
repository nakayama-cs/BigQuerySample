// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CompanyBatchServiceClient is the client API for CompanyBatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyBatchServiceClient interface {
	TaskLinkBusinessUnitManagement(ctx context.Context, in *TaskLinkBusinessUnitManagementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 履歴作成
	TaskCreateHistory(ctx context.Context, in *TaskCreateHistoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) メール送信
	TaskNotification(ctx context.Context, in *TaskNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 日次（夜間）スケジュール実行
	ScheduleTaskSharetoBusinessUnitManagement(ctx context.Context, in *ScheduleTaskSharetoBusinessUnitManagementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 自社企業情報公開
	TaskSharetoBusinessUnitManagement(ctx context.Context, in *TaskSharetoBusinessUnitManagementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 自社代表連絡先作成
	TaskCreateMainContact(ctx context.Context, in *TaskCreateMainContactRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 取引先管理（連絡先）公開
	TaskSharetoBusinessUnitContact(ctx context.Context, in *TaskSharetoBusinessUnitContactRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 更新/通知有効化
	TaskActivateBusinessUnit(ctx context.Context, in *TaskActivateBusinessUnitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type companyBatchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyBatchServiceClient(cc grpc.ClientConnInterface) CompanyBatchServiceClient {
	return &companyBatchServiceClient{cc}
}

func (c *companyBatchServiceClient) TaskLinkBusinessUnitManagement(ctx context.Context, in *TaskLinkBusinessUnitManagementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.company.CompanyBatchService/TaskLinkBusinessUnitManagement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyBatchServiceClient) TaskCreateHistory(ctx context.Context, in *TaskCreateHistoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.company.CompanyBatchService/TaskCreateHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyBatchServiceClient) TaskNotification(ctx context.Context, in *TaskNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.company.CompanyBatchService/TaskNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyBatchServiceClient) ScheduleTaskSharetoBusinessUnitManagement(ctx context.Context, in *ScheduleTaskSharetoBusinessUnitManagementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.company.CompanyBatchService/ScheduleTaskSharetoBusinessUnitManagement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyBatchServiceClient) TaskSharetoBusinessUnitManagement(ctx context.Context, in *TaskSharetoBusinessUnitManagementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.company.CompanyBatchService/TaskSharetoBusinessUnitManagement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyBatchServiceClient) TaskCreateMainContact(ctx context.Context, in *TaskCreateMainContactRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.company.CompanyBatchService/TaskCreateMainContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyBatchServiceClient) TaskSharetoBusinessUnitContact(ctx context.Context, in *TaskSharetoBusinessUnitContactRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.company.CompanyBatchService/TaskSharetoBusinessUnitContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyBatchServiceClient) TaskActivateBusinessUnit(ctx context.Context, in *TaskActivateBusinessUnitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.company.CompanyBatchService/TaskActivateBusinessUnit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyBatchServiceServer is the server API for CompanyBatchService service.
// All implementations must embed UnimplementedCompanyBatchServiceServer
// for forward compatibility
type CompanyBatchServiceServer interface {
	TaskLinkBusinessUnitManagement(context.Context, *TaskLinkBusinessUnitManagementRequest) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 履歴作成
	TaskCreateHistory(context.Context, *TaskCreateHistoryRequest) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) メール送信
	TaskNotification(context.Context, *TaskNotificationRequest) (*emptypb.Empty, error)
	// 日次（夜間）スケジュール実行
	ScheduleTaskSharetoBusinessUnitManagement(context.Context, *ScheduleTaskSharetoBusinessUnitManagementRequest) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 自社企業情報公開
	TaskSharetoBusinessUnitManagement(context.Context, *TaskSharetoBusinessUnitManagementRequest) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 自社代表連絡先作成
	TaskCreateMainContact(context.Context, *TaskCreateMainContactRequest) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 取引先管理（連絡先）公開
	TaskSharetoBusinessUnitContact(context.Context, *TaskSharetoBusinessUnitContactRequest) (*emptypb.Empty, error)
	// 非同期Queue(バッチ) 更新/通知有効化
	TaskActivateBusinessUnit(context.Context, *TaskActivateBusinessUnitRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCompanyBatchServiceServer()
}

// UnimplementedCompanyBatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyBatchServiceServer struct {
}

func (UnimplementedCompanyBatchServiceServer) TaskLinkBusinessUnitManagement(context.Context, *TaskLinkBusinessUnitManagementRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskLinkBusinessUnitManagement not implemented")
}
func (UnimplementedCompanyBatchServiceServer) TaskCreateHistory(context.Context, *TaskCreateHistoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskCreateHistory not implemented")
}
func (UnimplementedCompanyBatchServiceServer) TaskNotification(context.Context, *TaskNotificationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskNotification not implemented")
}
func (UnimplementedCompanyBatchServiceServer) ScheduleTaskSharetoBusinessUnitManagement(context.Context, *ScheduleTaskSharetoBusinessUnitManagementRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleTaskSharetoBusinessUnitManagement not implemented")
}
func (UnimplementedCompanyBatchServiceServer) TaskSharetoBusinessUnitManagement(context.Context, *TaskSharetoBusinessUnitManagementRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskSharetoBusinessUnitManagement not implemented")
}
func (UnimplementedCompanyBatchServiceServer) TaskCreateMainContact(context.Context, *TaskCreateMainContactRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskCreateMainContact not implemented")
}
func (UnimplementedCompanyBatchServiceServer) TaskSharetoBusinessUnitContact(context.Context, *TaskSharetoBusinessUnitContactRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskSharetoBusinessUnitContact not implemented")
}
func (UnimplementedCompanyBatchServiceServer) TaskActivateBusinessUnit(context.Context, *TaskActivateBusinessUnitRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskActivateBusinessUnit not implemented")
}
func (UnimplementedCompanyBatchServiceServer) mustEmbedUnimplementedCompanyBatchServiceServer() {}

// UnsafeCompanyBatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyBatchServiceServer will
// result in compilation errors.
type UnsafeCompanyBatchServiceServer interface {
	mustEmbedUnimplementedCompanyBatchServiceServer()
}

func RegisterCompanyBatchServiceServer(s grpc.ServiceRegistrar, srv CompanyBatchServiceServer) {
	s.RegisterService(&CompanyBatchService_ServiceDesc, srv)
}

func _CompanyBatchService_TaskLinkBusinessUnitManagement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskLinkBusinessUnitManagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyBatchServiceServer).TaskLinkBusinessUnitManagement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.company.CompanyBatchService/TaskLinkBusinessUnitManagement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyBatchServiceServer).TaskLinkBusinessUnitManagement(ctx, req.(*TaskLinkBusinessUnitManagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyBatchService_TaskCreateHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyBatchServiceServer).TaskCreateHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.company.CompanyBatchService/TaskCreateHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyBatchServiceServer).TaskCreateHistory(ctx, req.(*TaskCreateHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyBatchService_TaskNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyBatchServiceServer).TaskNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.company.CompanyBatchService/TaskNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyBatchServiceServer).TaskNotification(ctx, req.(*TaskNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyBatchService_ScheduleTaskSharetoBusinessUnitManagement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleTaskSharetoBusinessUnitManagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyBatchServiceServer).ScheduleTaskSharetoBusinessUnitManagement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.company.CompanyBatchService/ScheduleTaskSharetoBusinessUnitManagement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyBatchServiceServer).ScheduleTaskSharetoBusinessUnitManagement(ctx, req.(*ScheduleTaskSharetoBusinessUnitManagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyBatchService_TaskSharetoBusinessUnitManagement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskSharetoBusinessUnitManagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyBatchServiceServer).TaskSharetoBusinessUnitManagement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.company.CompanyBatchService/TaskSharetoBusinessUnitManagement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyBatchServiceServer).TaskSharetoBusinessUnitManagement(ctx, req.(*TaskSharetoBusinessUnitManagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyBatchService_TaskCreateMainContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateMainContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyBatchServiceServer).TaskCreateMainContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.company.CompanyBatchService/TaskCreateMainContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyBatchServiceServer).TaskCreateMainContact(ctx, req.(*TaskCreateMainContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyBatchService_TaskSharetoBusinessUnitContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskSharetoBusinessUnitContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyBatchServiceServer).TaskSharetoBusinessUnitContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.company.CompanyBatchService/TaskSharetoBusinessUnitContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyBatchServiceServer).TaskSharetoBusinessUnitContact(ctx, req.(*TaskSharetoBusinessUnitContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyBatchService_TaskActivateBusinessUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskActivateBusinessUnitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyBatchServiceServer).TaskActivateBusinessUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.company.CompanyBatchService/TaskActivateBusinessUnit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyBatchServiceServer).TaskActivateBusinessUnit(ctx, req.(*TaskActivateBusinessUnitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyBatchService_ServiceDesc is the grpc.ServiceDesc for CompanyBatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyBatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mtechnavi.api.company.CompanyBatchService",
	HandlerType: (*CompanyBatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskLinkBusinessUnitManagement",
			Handler:    _CompanyBatchService_TaskLinkBusinessUnitManagement_Handler,
		},
		{
			MethodName: "TaskCreateHistory",
			Handler:    _CompanyBatchService_TaskCreateHistory_Handler,
		},
		{
			MethodName: "TaskNotification",
			Handler:    _CompanyBatchService_TaskNotification_Handler,
		},
		{
			MethodName: "ScheduleTaskSharetoBusinessUnitManagement",
			Handler:    _CompanyBatchService_ScheduleTaskSharetoBusinessUnitManagement_Handler,
		},
		{
			MethodName: "TaskSharetoBusinessUnitManagement",
			Handler:    _CompanyBatchService_TaskSharetoBusinessUnitManagement_Handler,
		},
		{
			MethodName: "TaskCreateMainContact",
			Handler:    _CompanyBatchService_TaskCreateMainContact_Handler,
		},
		{
			MethodName: "TaskSharetoBusinessUnitContact",
			Handler:    _CompanyBatchService_TaskSharetoBusinessUnitContact_Handler,
		},
		{
			MethodName: "TaskActivateBusinessUnit",
			Handler:    _CompanyBatchService_TaskActivateBusinessUnit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "company_batch.proto",
}
