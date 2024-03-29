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

// BatchServiceClient is the client API for BatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BatchServiceClient interface {
	ScheduleJobDailyMail(ctx context.Context, in *ScheduleJobDailyMailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ScheduleJobSharetoSurveyReminder(ctx context.Context, in *ScheduleJobSharetoSurveyReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ScheduleJobSharetoBusinessUnitManagement(ctx context.Context, in *ScheduleJobSharetoBusinessUnitManagementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type batchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBatchServiceClient(cc grpc.ClientConnInterface) BatchServiceClient {
	return &batchServiceClient{cc}
}

func (c *batchServiceClient) ScheduleJobDailyMail(ctx context.Context, in *ScheduleJobDailyMailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.batch.BatchService/ScheduleJobDailyMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchServiceClient) ScheduleJobSharetoSurveyReminder(ctx context.Context, in *ScheduleJobSharetoSurveyReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.batch.BatchService/ScheduleJobSharetoSurveyReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchServiceClient) ScheduleJobSharetoBusinessUnitManagement(ctx context.Context, in *ScheduleJobSharetoBusinessUnitManagementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.batch.BatchService/ScheduleJobSharetoBusinessUnitManagement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatchServiceServer is the server API for BatchService service.
// All implementations must embed UnimplementedBatchServiceServer
// for forward compatibility
type BatchServiceServer interface {
	ScheduleJobDailyMail(context.Context, *ScheduleJobDailyMailRequest) (*emptypb.Empty, error)
	ScheduleJobSharetoSurveyReminder(context.Context, *ScheduleJobSharetoSurveyReminderRequest) (*emptypb.Empty, error)
	ScheduleJobSharetoBusinessUnitManagement(context.Context, *ScheduleJobSharetoBusinessUnitManagementRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBatchServiceServer()
}

// UnimplementedBatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBatchServiceServer struct {
}

func (UnimplementedBatchServiceServer) ScheduleJobDailyMail(context.Context, *ScheduleJobDailyMailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleJobDailyMail not implemented")
}
func (UnimplementedBatchServiceServer) ScheduleJobSharetoSurveyReminder(context.Context, *ScheduleJobSharetoSurveyReminderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleJobSharetoSurveyReminder not implemented")
}
func (UnimplementedBatchServiceServer) ScheduleJobSharetoBusinessUnitManagement(context.Context, *ScheduleJobSharetoBusinessUnitManagementRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleJobSharetoBusinessUnitManagement not implemented")
}
func (UnimplementedBatchServiceServer) mustEmbedUnimplementedBatchServiceServer() {}

// UnsafeBatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BatchServiceServer will
// result in compilation errors.
type UnsafeBatchServiceServer interface {
	mustEmbedUnimplementedBatchServiceServer()
}

func RegisterBatchServiceServer(s grpc.ServiceRegistrar, srv BatchServiceServer) {
	s.RegisterService(&BatchService_ServiceDesc, srv)
}

func _BatchService_ScheduleJobDailyMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleJobDailyMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchServiceServer).ScheduleJobDailyMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.batch.BatchService/ScheduleJobDailyMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchServiceServer).ScheduleJobDailyMail(ctx, req.(*ScheduleJobDailyMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchService_ScheduleJobSharetoSurveyReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleJobSharetoSurveyReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchServiceServer).ScheduleJobSharetoSurveyReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.batch.BatchService/ScheduleJobSharetoSurveyReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchServiceServer).ScheduleJobSharetoSurveyReminder(ctx, req.(*ScheduleJobSharetoSurveyReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchService_ScheduleJobSharetoBusinessUnitManagement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleJobSharetoBusinessUnitManagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchServiceServer).ScheduleJobSharetoBusinessUnitManagement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.batch.BatchService/ScheduleJobSharetoBusinessUnitManagement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchServiceServer).ScheduleJobSharetoBusinessUnitManagement(ctx, req.(*ScheduleJobSharetoBusinessUnitManagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BatchService_ServiceDesc is the grpc.ServiceDesc for BatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mtechnavi.api.batch.BatchService",
	HandlerType: (*BatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleJobDailyMail",
			Handler:    _BatchService_ScheduleJobDailyMail_Handler,
		},
		{
			MethodName: "ScheduleJobSharetoSurveyReminder",
			Handler:    _BatchService_ScheduleJobSharetoSurveyReminder_Handler,
		},
		{
			MethodName: "ScheduleJobSharetoBusinessUnitManagement",
			Handler:    _BatchService_ScheduleJobSharetoBusinessUnitManagement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "batch.proto",
}
