// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	record "mtechnavi/sharelib/protobuf/record"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataproxyClient is the client API for Dataproxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataproxyClient interface {
	// レコード読み書き
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*record.Record, error)
	GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*record.Record, error)
	ListRecords(ctx context.Context, in *ListRecordsRequest, opts ...grpc.CallOption) (*ListRecordsResponse, error)
	UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*record.Record, error)
	DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// レコード複数書込み（同一トランザクションで実行する）
	CreateRecords(ctx context.Context, in *CreateRecordsRequest, opts ...grpc.CallOption) (*CreateRecordsResponse, error)
	UpdateRecords(ctx context.Context, in *UpdateRecordsRequest, opts ...grpc.CallOption) (*UpdateRecordsResponse, error)
	DeleteRecords(ctx context.Context, in *DeleteRecordsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateOrUpdateRecords(ctx context.Context, in *CreateOrUpdateRecordsRequest, opts ...grpc.CallOption) (*CreateOrUpdateRecordsResponse, error)
	// レコードのStream
	StreamRecords(ctx context.Context, in *StreamRecordsRequest, opts ...grpc.CallOption) (Dataproxy_StreamRecordsClient, error)
	// レコード共有
	GetShareRecord(ctx context.Context, in *GetShareRecordRequest, opts ...grpc.CallOption) (*record.Record, error)
	ShareRecord(ctx context.Context, in *ShareRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListSharedTenants(ctx context.Context, in *ListSharedTenantsRequest, opts ...grpc.CallOption) (*ListSharedTenantsResponse, error)
	UnshareRecord(ctx context.Context, in *UnshareRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// アセット管理
	CreateAssetData(ctx context.Context, in *CreateAssetDataRequest, opts ...grpc.CallOption) (*AssetData, error)
	GetAssetData(ctx context.Context, in *GetAssetDataRequest, opts ...grpc.CallOption) (*AssetData, error)
	DeleteAssetData(ctx context.Context, in *DeleteAssetDataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GenerateSignedUrl(ctx context.Context, in *SignedUrlRequest, opts ...grpc.CallOption) (*SignedUrlResponse, error)
	// ログ
	ListAuditLogs(ctx context.Context, in *ListAuditLogsRequest, opts ...grpc.CallOption) (*ListAuditLogsResponse, error)
}

type dataproxyClient struct {
	cc grpc.ClientConnInterface
}

func NewDataproxyClient(cc grpc.ClientConnInterface) DataproxyClient {
	return &dataproxyClient{cc}
}

func (c *dataproxyClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*record.Record, error) {
	out := new(record.Record)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/CreateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*record.Record, error) {
	out := new(record.Record)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/GetRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) ListRecords(ctx context.Context, in *ListRecordsRequest, opts ...grpc.CallOption) (*ListRecordsResponse, error) {
	out := new(ListRecordsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/ListRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*record.Record, error) {
	out := new(record.Record)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/UpdateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/DeleteRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) CreateRecords(ctx context.Context, in *CreateRecordsRequest, opts ...grpc.CallOption) (*CreateRecordsResponse, error) {
	out := new(CreateRecordsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/CreateRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) UpdateRecords(ctx context.Context, in *UpdateRecordsRequest, opts ...grpc.CallOption) (*UpdateRecordsResponse, error) {
	out := new(UpdateRecordsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/UpdateRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) DeleteRecords(ctx context.Context, in *DeleteRecordsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/DeleteRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) CreateOrUpdateRecords(ctx context.Context, in *CreateOrUpdateRecordsRequest, opts ...grpc.CallOption) (*CreateOrUpdateRecordsResponse, error) {
	out := new(CreateOrUpdateRecordsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/CreateOrUpdateRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) StreamRecords(ctx context.Context, in *StreamRecordsRequest, opts ...grpc.CallOption) (Dataproxy_StreamRecordsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dataproxy_ServiceDesc.Streams[0], "/mtechnavi.api.dataproxy.Dataproxy/StreamRecords", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataproxyStreamRecordsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dataproxy_StreamRecordsClient interface {
	Recv() (*record.Record, error)
	grpc.ClientStream
}

type dataproxyStreamRecordsClient struct {
	grpc.ClientStream
}

func (x *dataproxyStreamRecordsClient) Recv() (*record.Record, error) {
	m := new(record.Record)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataproxyClient) GetShareRecord(ctx context.Context, in *GetShareRecordRequest, opts ...grpc.CallOption) (*record.Record, error) {
	out := new(record.Record)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/GetShareRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) ShareRecord(ctx context.Context, in *ShareRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/ShareRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) ListSharedTenants(ctx context.Context, in *ListSharedTenantsRequest, opts ...grpc.CallOption) (*ListSharedTenantsResponse, error) {
	out := new(ListSharedTenantsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/ListSharedTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) UnshareRecord(ctx context.Context, in *UnshareRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/UnshareRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) CreateAssetData(ctx context.Context, in *CreateAssetDataRequest, opts ...grpc.CallOption) (*AssetData, error) {
	out := new(AssetData)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/CreateAssetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) GetAssetData(ctx context.Context, in *GetAssetDataRequest, opts ...grpc.CallOption) (*AssetData, error) {
	out := new(AssetData)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/GetAssetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) DeleteAssetData(ctx context.Context, in *DeleteAssetDataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/DeleteAssetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) GenerateSignedUrl(ctx context.Context, in *SignedUrlRequest, opts ...grpc.CallOption) (*SignedUrlResponse, error) {
	out := new(SignedUrlResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/GenerateSignedUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyClient) ListAuditLogs(ctx context.Context, in *ListAuditLogsRequest, opts ...grpc.CallOption) (*ListAuditLogsResponse, error) {
	out := new(ListAuditLogsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.dataproxy.Dataproxy/ListAuditLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataproxyServer is the server API for Dataproxy service.
// All implementations must embed UnimplementedDataproxyServer
// for forward compatibility
type DataproxyServer interface {
	// レコード読み書き
	CreateRecord(context.Context, *CreateRecordRequest) (*record.Record, error)
	GetRecord(context.Context, *GetRecordRequest) (*record.Record, error)
	ListRecords(context.Context, *ListRecordsRequest) (*ListRecordsResponse, error)
	UpdateRecord(context.Context, *UpdateRecordRequest) (*record.Record, error)
	DeleteRecord(context.Context, *DeleteRecordRequest) (*emptypb.Empty, error)
	// レコード複数書込み（同一トランザクションで実行する）
	CreateRecords(context.Context, *CreateRecordsRequest) (*CreateRecordsResponse, error)
	UpdateRecords(context.Context, *UpdateRecordsRequest) (*UpdateRecordsResponse, error)
	DeleteRecords(context.Context, *DeleteRecordsRequest) (*emptypb.Empty, error)
	CreateOrUpdateRecords(context.Context, *CreateOrUpdateRecordsRequest) (*CreateOrUpdateRecordsResponse, error)
	// レコードのStream
	StreamRecords(*StreamRecordsRequest, Dataproxy_StreamRecordsServer) error
	// レコード共有
	GetShareRecord(context.Context, *GetShareRecordRequest) (*record.Record, error)
	ShareRecord(context.Context, *ShareRecordRequest) (*emptypb.Empty, error)
	ListSharedTenants(context.Context, *ListSharedTenantsRequest) (*ListSharedTenantsResponse, error)
	UnshareRecord(context.Context, *UnshareRecordRequest) (*emptypb.Empty, error)
	// アセット管理
	CreateAssetData(context.Context, *CreateAssetDataRequest) (*AssetData, error)
	GetAssetData(context.Context, *GetAssetDataRequest) (*AssetData, error)
	DeleteAssetData(context.Context, *DeleteAssetDataRequest) (*emptypb.Empty, error)
	GenerateSignedUrl(context.Context, *SignedUrlRequest) (*SignedUrlResponse, error)
	// ログ
	ListAuditLogs(context.Context, *ListAuditLogsRequest) (*ListAuditLogsResponse, error)
	mustEmbedUnimplementedDataproxyServer()
}

// UnimplementedDataproxyServer must be embedded to have forward compatible implementations.
type UnimplementedDataproxyServer struct {
}

func (UnimplementedDataproxyServer) CreateRecord(context.Context, *CreateRecordRequest) (*record.Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedDataproxyServer) GetRecord(context.Context, *GetRecordRequest) (*record.Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}
func (UnimplementedDataproxyServer) ListRecords(context.Context, *ListRecordsRequest) (*ListRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecords not implemented")
}
func (UnimplementedDataproxyServer) UpdateRecord(context.Context, *UpdateRecordRequest) (*record.Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecord not implemented")
}
func (UnimplementedDataproxyServer) DeleteRecord(context.Context, *DeleteRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecord not implemented")
}
func (UnimplementedDataproxyServer) CreateRecords(context.Context, *CreateRecordsRequest) (*CreateRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecords not implemented")
}
func (UnimplementedDataproxyServer) UpdateRecords(context.Context, *UpdateRecordsRequest) (*UpdateRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecords not implemented")
}
func (UnimplementedDataproxyServer) DeleteRecords(context.Context, *DeleteRecordsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecords not implemented")
}
func (UnimplementedDataproxyServer) CreateOrUpdateRecords(context.Context, *CreateOrUpdateRecordsRequest) (*CreateOrUpdateRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateRecords not implemented")
}
func (UnimplementedDataproxyServer) StreamRecords(*StreamRecordsRequest, Dataproxy_StreamRecordsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRecords not implemented")
}
func (UnimplementedDataproxyServer) GetShareRecord(context.Context, *GetShareRecordRequest) (*record.Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShareRecord not implemented")
}
func (UnimplementedDataproxyServer) ShareRecord(context.Context, *ShareRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareRecord not implemented")
}
func (UnimplementedDataproxyServer) ListSharedTenants(context.Context, *ListSharedTenantsRequest) (*ListSharedTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSharedTenants not implemented")
}
func (UnimplementedDataproxyServer) UnshareRecord(context.Context, *UnshareRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnshareRecord not implemented")
}
func (UnimplementedDataproxyServer) CreateAssetData(context.Context, *CreateAssetDataRequest) (*AssetData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssetData not implemented")
}
func (UnimplementedDataproxyServer) GetAssetData(context.Context, *GetAssetDataRequest) (*AssetData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetData not implemented")
}
func (UnimplementedDataproxyServer) DeleteAssetData(context.Context, *DeleteAssetDataRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAssetData not implemented")
}
func (UnimplementedDataproxyServer) GenerateSignedUrl(context.Context, *SignedUrlRequest) (*SignedUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSignedUrl not implemented")
}
func (UnimplementedDataproxyServer) ListAuditLogs(context.Context, *ListAuditLogsRequest) (*ListAuditLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuditLogs not implemented")
}
func (UnimplementedDataproxyServer) mustEmbedUnimplementedDataproxyServer() {}

// UnsafeDataproxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataproxyServer will
// result in compilation errors.
type UnsafeDataproxyServer interface {
	mustEmbedUnimplementedDataproxyServer()
}

func RegisterDataproxyServer(s grpc.ServiceRegistrar, srv DataproxyServer) {
	s.RegisterService(&Dataproxy_ServiceDesc, srv)
}

func _Dataproxy_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).GetRecord(ctx, req.(*GetRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_ListRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).ListRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/ListRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).ListRecords(ctx, req.(*ListRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/UpdateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).UpdateRecord(ctx, req.(*UpdateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).DeleteRecord(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_CreateRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).CreateRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/CreateRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).CreateRecords(ctx, req.(*CreateRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_UpdateRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).UpdateRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/UpdateRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).UpdateRecords(ctx, req.(*UpdateRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_DeleteRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).DeleteRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/DeleteRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).DeleteRecords(ctx, req.(*DeleteRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_CreateOrUpdateRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).CreateOrUpdateRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/CreateOrUpdateRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).CreateOrUpdateRecords(ctx, req.(*CreateOrUpdateRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_StreamRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRecordsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataproxyServer).StreamRecords(m, &dataproxyStreamRecordsServer{stream})
}

type Dataproxy_StreamRecordsServer interface {
	Send(*record.Record) error
	grpc.ServerStream
}

type dataproxyStreamRecordsServer struct {
	grpc.ServerStream
}

func (x *dataproxyStreamRecordsServer) Send(m *record.Record) error {
	return x.ServerStream.SendMsg(m)
}

func _Dataproxy_GetShareRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShareRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).GetShareRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/GetShareRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).GetShareRecord(ctx, req.(*GetShareRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_ShareRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).ShareRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/ShareRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).ShareRecord(ctx, req.(*ShareRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_ListSharedTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSharedTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).ListSharedTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/ListSharedTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).ListSharedTenants(ctx, req.(*ListSharedTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_UnshareRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnshareRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).UnshareRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/UnshareRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).UnshareRecord(ctx, req.(*UnshareRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_CreateAssetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).CreateAssetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/CreateAssetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).CreateAssetData(ctx, req.(*CreateAssetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_GetAssetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).GetAssetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/GetAssetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).GetAssetData(ctx, req.(*GetAssetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_DeleteAssetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAssetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).DeleteAssetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/DeleteAssetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).DeleteAssetData(ctx, req.(*DeleteAssetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_GenerateSignedUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).GenerateSignedUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/GenerateSignedUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).GenerateSignedUrl(ctx, req.(*SignedUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataproxy_ListAuditLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuditLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServer).ListAuditLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.dataproxy.Dataproxy/ListAuditLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServer).ListAuditLogs(ctx, req.(*ListAuditLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dataproxy_ServiceDesc is the grpc.ServiceDesc for Dataproxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dataproxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mtechnavi.api.dataproxy.Dataproxy",
	HandlerType: (*DataproxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecord",
			Handler:    _Dataproxy_CreateRecord_Handler,
		},
		{
			MethodName: "GetRecord",
			Handler:    _Dataproxy_GetRecord_Handler,
		},
		{
			MethodName: "ListRecords",
			Handler:    _Dataproxy_ListRecords_Handler,
		},
		{
			MethodName: "UpdateRecord",
			Handler:    _Dataproxy_UpdateRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _Dataproxy_DeleteRecord_Handler,
		},
		{
			MethodName: "CreateRecords",
			Handler:    _Dataproxy_CreateRecords_Handler,
		},
		{
			MethodName: "UpdateRecords",
			Handler:    _Dataproxy_UpdateRecords_Handler,
		},
		{
			MethodName: "DeleteRecords",
			Handler:    _Dataproxy_DeleteRecords_Handler,
		},
		{
			MethodName: "CreateOrUpdateRecords",
			Handler:    _Dataproxy_CreateOrUpdateRecords_Handler,
		},
		{
			MethodName: "GetShareRecord",
			Handler:    _Dataproxy_GetShareRecord_Handler,
		},
		{
			MethodName: "ShareRecord",
			Handler:    _Dataproxy_ShareRecord_Handler,
		},
		{
			MethodName: "ListSharedTenants",
			Handler:    _Dataproxy_ListSharedTenants_Handler,
		},
		{
			MethodName: "UnshareRecord",
			Handler:    _Dataproxy_UnshareRecord_Handler,
		},
		{
			MethodName: "CreateAssetData",
			Handler:    _Dataproxy_CreateAssetData_Handler,
		},
		{
			MethodName: "GetAssetData",
			Handler:    _Dataproxy_GetAssetData_Handler,
		},
		{
			MethodName: "DeleteAssetData",
			Handler:    _Dataproxy_DeleteAssetData_Handler,
		},
		{
			MethodName: "GenerateSignedUrl",
			Handler:    _Dataproxy_GenerateSignedUrl_Handler,
		},
		{
			MethodName: "ListAuditLogs",
			Handler:    _Dataproxy_ListAuditLogs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRecords",
			Handler:       _Dataproxy_StreamRecords_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dataproxy.proto",
}
