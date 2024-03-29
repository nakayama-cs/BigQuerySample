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

// ForumServiceClient is the client API for ForumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ForumServiceClient interface {
	// 親スレッド
	//
	// Implementation Note::
	//
	//	Create/Updateされたタイミングでレコードをsharetoする
	//	shareto先はBaseThread.company_idsを参照
	CreateBaseThread(ctx context.Context, in *CreateBaseThreadRequest, opts ...grpc.CallOption) (*BaseThread, error)
	ListBaseThreads(ctx context.Context, in *ListBaseThreadsRequest, opts ...grpc.CallOption) (*ListBaseThreadsResponse, error)
	GetBaseThread(ctx context.Context, in *GetBaseThreadRequest, opts ...grpc.CallOption) (*BaseThread, error)
	UpdateBaseThread(ctx context.Context, in *UpdateBaseThreadRequest, opts ...grpc.CallOption) (*BaseThread, error)
	DeleteBaseThread(ctx context.Context, in *DeleteBaseThreadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 問合せ
	//
	// Implementation Note::
	//
	//	Create/Updateされたタイミングで公開設定が公開/企業指定の場合はレコードをsharetoする
	//	shareto先は親スレッド（BaseThread.company_ids）を参照
	CreateThread(ctx context.Context, in *CreateThreadRequest, opts ...grpc.CallOption) (*Thread, error)
	ListThreads(ctx context.Context, in *ListThreadsRequest, opts ...grpc.CallOption) (*ListThreadsResponse, error)
	GetThread(ctx context.Context, in *GetThreadRequest, opts ...grpc.CallOption) (*Thread, error)
	UpdateThread(ctx context.Context, in *UpdateThreadRequest, opts ...grpc.CallOption) (*Thread, error)
	DeleteThread(ctx context.Context, in *DeleteThreadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 投稿
	//
	// Implementation Note::
	//
	//	Create/Update のタイミングでThreadSummaryをUpdateする
	//
	//	Create/Updateされたタイミングで公開設定が公開/企業指定の場合はレコードをsharetoする
	//	shareto先は親スレッド（BaseThread.company_ids）を参照
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*Comment, error)
	ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error)
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*Comment, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*Comment, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// マーカーをつける
	//
	// Implementation Note::
	//
	//  1. 指定されたmodeでMarkerを更新する
	//     mode = Complete の場合、指定されたマーカーに置き換える（元のマーカーを削除）
	//     mode = Incremental の場合、指定されたマーカーを追加
	//
	//  2. ThreadUserSummaryをUpdateする
	SetMarker(ctx context.Context, in *SetMarkerRequest, opts ...grpc.CallOption) (*SetMarkerResponse, error)
	// マーカーをはずす
	//
	// Implementation Note::
	//
	//  1. 指定されたmodeでMarkerを更新する
	//     mode = Complete の場合、全てのマーカーを削除
	//     mode = Incremental の場合、指定されたマーカーのみを削除
	//
	//  2. ThreadUserSummaryをUpdateする
	UnsetMarker(ctx context.Context, in *UnsetMarkerRequest, opts ...grpc.CallOption) (*UnsetMarkerResponse, error)
}

type forumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewForumServiceClient(cc grpc.ClientConnInterface) ForumServiceClient {
	return &forumServiceClient{cc}
}

func (c *forumServiceClient) CreateBaseThread(ctx context.Context, in *CreateBaseThreadRequest, opts ...grpc.CallOption) (*BaseThread, error) {
	out := new(BaseThread)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/CreateBaseThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) ListBaseThreads(ctx context.Context, in *ListBaseThreadsRequest, opts ...grpc.CallOption) (*ListBaseThreadsResponse, error) {
	out := new(ListBaseThreadsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/ListBaseThreads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetBaseThread(ctx context.Context, in *GetBaseThreadRequest, opts ...grpc.CallOption) (*BaseThread, error) {
	out := new(BaseThread)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/GetBaseThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) UpdateBaseThread(ctx context.Context, in *UpdateBaseThreadRequest, opts ...grpc.CallOption) (*BaseThread, error) {
	out := new(BaseThread)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/UpdateBaseThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) DeleteBaseThread(ctx context.Context, in *DeleteBaseThreadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/DeleteBaseThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) CreateThread(ctx context.Context, in *CreateThreadRequest, opts ...grpc.CallOption) (*Thread, error) {
	out := new(Thread)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/CreateThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) ListThreads(ctx context.Context, in *ListThreadsRequest, opts ...grpc.CallOption) (*ListThreadsResponse, error) {
	out := new(ListThreadsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/ListThreads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetThread(ctx context.Context, in *GetThreadRequest, opts ...grpc.CallOption) (*Thread, error) {
	out := new(Thread)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/GetThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) UpdateThread(ctx context.Context, in *UpdateThreadRequest, opts ...grpc.CallOption) (*Thread, error) {
	out := new(Thread)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/UpdateThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) DeleteThread(ctx context.Context, in *DeleteThreadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/DeleteThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error) {
	out := new(ListCommentsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/ListComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) SetMarker(ctx context.Context, in *SetMarkerRequest, opts ...grpc.CallOption) (*SetMarkerResponse, error) {
	out := new(SetMarkerResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/SetMarker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) UnsetMarker(ctx context.Context, in *UnsetMarkerRequest, opts ...grpc.CallOption) (*UnsetMarkerResponse, error) {
	out := new(UnsetMarkerResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumService/UnsetMarker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForumServiceServer is the server API for ForumService service.
// All implementations must embed UnimplementedForumServiceServer
// for forward compatibility
type ForumServiceServer interface {
	// 親スレッド
	//
	// Implementation Note::
	//
	//	Create/Updateされたタイミングでレコードをsharetoする
	//	shareto先はBaseThread.company_idsを参照
	CreateBaseThread(context.Context, *CreateBaseThreadRequest) (*BaseThread, error)
	ListBaseThreads(context.Context, *ListBaseThreadsRequest) (*ListBaseThreadsResponse, error)
	GetBaseThread(context.Context, *GetBaseThreadRequest) (*BaseThread, error)
	UpdateBaseThread(context.Context, *UpdateBaseThreadRequest) (*BaseThread, error)
	DeleteBaseThread(context.Context, *DeleteBaseThreadRequest) (*emptypb.Empty, error)
	// 問合せ
	//
	// Implementation Note::
	//
	//	Create/Updateされたタイミングで公開設定が公開/企業指定の場合はレコードをsharetoする
	//	shareto先は親スレッド（BaseThread.company_ids）を参照
	CreateThread(context.Context, *CreateThreadRequest) (*Thread, error)
	ListThreads(context.Context, *ListThreadsRequest) (*ListThreadsResponse, error)
	GetThread(context.Context, *GetThreadRequest) (*Thread, error)
	UpdateThread(context.Context, *UpdateThreadRequest) (*Thread, error)
	DeleteThread(context.Context, *DeleteThreadRequest) (*emptypb.Empty, error)
	// 投稿
	//
	// Implementation Note::
	//
	//	Create/Update のタイミングでThreadSummaryをUpdateする
	//
	//	Create/Updateされたタイミングで公開設定が公開/企業指定の場合はレコードをsharetoする
	//	shareto先は親スレッド（BaseThread.company_ids）を参照
	CreateComment(context.Context, *CreateCommentRequest) (*Comment, error)
	ListComments(context.Context, *ListCommentsRequest) (*ListCommentsResponse, error)
	GetComment(context.Context, *GetCommentRequest) (*Comment, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*Comment, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error)
	// マーカーをつける
	//
	// Implementation Note::
	//
	//  1. 指定されたmodeでMarkerを更新する
	//     mode = Complete の場合、指定されたマーカーに置き換える（元のマーカーを削除）
	//     mode = Incremental の場合、指定されたマーカーを追加
	//
	//  2. ThreadUserSummaryをUpdateする
	SetMarker(context.Context, *SetMarkerRequest) (*SetMarkerResponse, error)
	// マーカーをはずす
	//
	// Implementation Note::
	//
	//  1. 指定されたmodeでMarkerを更新する
	//     mode = Complete の場合、全てのマーカーを削除
	//     mode = Incremental の場合、指定されたマーカーのみを削除
	//
	//  2. ThreadUserSummaryをUpdateする
	UnsetMarker(context.Context, *UnsetMarkerRequest) (*UnsetMarkerResponse, error)
	mustEmbedUnimplementedForumServiceServer()
}

// UnimplementedForumServiceServer must be embedded to have forward compatible implementations.
type UnimplementedForumServiceServer struct {
}

func (UnimplementedForumServiceServer) CreateBaseThread(context.Context, *CreateBaseThreadRequest) (*BaseThread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBaseThread not implemented")
}
func (UnimplementedForumServiceServer) ListBaseThreads(context.Context, *ListBaseThreadsRequest) (*ListBaseThreadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBaseThreads not implemented")
}
func (UnimplementedForumServiceServer) GetBaseThread(context.Context, *GetBaseThreadRequest) (*BaseThread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBaseThread not implemented")
}
func (UnimplementedForumServiceServer) UpdateBaseThread(context.Context, *UpdateBaseThreadRequest) (*BaseThread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBaseThread not implemented")
}
func (UnimplementedForumServiceServer) DeleteBaseThread(context.Context, *DeleteBaseThreadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBaseThread not implemented")
}
func (UnimplementedForumServiceServer) CreateThread(context.Context, *CreateThreadRequest) (*Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateThread not implemented")
}
func (UnimplementedForumServiceServer) ListThreads(context.Context, *ListThreadsRequest) (*ListThreadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListThreads not implemented")
}
func (UnimplementedForumServiceServer) GetThread(context.Context, *GetThreadRequest) (*Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThread not implemented")
}
func (UnimplementedForumServiceServer) UpdateThread(context.Context, *UpdateThreadRequest) (*Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateThread not implemented")
}
func (UnimplementedForumServiceServer) DeleteThread(context.Context, *DeleteThreadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteThread not implemented")
}
func (UnimplementedForumServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedForumServiceServer) ListComments(context.Context, *ListCommentsRequest) (*ListCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComments not implemented")
}
func (UnimplementedForumServiceServer) GetComment(context.Context, *GetCommentRequest) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedForumServiceServer) UpdateComment(context.Context, *UpdateCommentRequest) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedForumServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedForumServiceServer) SetMarker(context.Context, *SetMarkerRequest) (*SetMarkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMarker not implemented")
}
func (UnimplementedForumServiceServer) UnsetMarker(context.Context, *UnsetMarkerRequest) (*UnsetMarkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsetMarker not implemented")
}
func (UnimplementedForumServiceServer) mustEmbedUnimplementedForumServiceServer() {}

// UnsafeForumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForumServiceServer will
// result in compilation errors.
type UnsafeForumServiceServer interface {
	mustEmbedUnimplementedForumServiceServer()
}

func RegisterForumServiceServer(s grpc.ServiceRegistrar, srv ForumServiceServer) {
	s.RegisterService(&ForumService_ServiceDesc, srv)
}

func _ForumService_CreateBaseThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBaseThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreateBaseThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/CreateBaseThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreateBaseThread(ctx, req.(*CreateBaseThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_ListBaseThreads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBaseThreadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).ListBaseThreads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/ListBaseThreads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).ListBaseThreads(ctx, req.(*ListBaseThreadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetBaseThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBaseThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetBaseThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/GetBaseThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetBaseThread(ctx, req.(*GetBaseThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_UpdateBaseThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBaseThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).UpdateBaseThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/UpdateBaseThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).UpdateBaseThread(ctx, req.(*UpdateBaseThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_DeleteBaseThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBaseThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).DeleteBaseThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/DeleteBaseThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).DeleteBaseThread(ctx, req.(*DeleteBaseThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_CreateThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreateThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/CreateThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreateThread(ctx, req.(*CreateThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_ListThreads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListThreadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).ListThreads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/ListThreads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).ListThreads(ctx, req.(*ListThreadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/GetThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetThread(ctx, req.(*GetThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_UpdateThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).UpdateThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/UpdateThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).UpdateThread(ctx, req.(*UpdateThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_DeleteThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).DeleteThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/DeleteThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).DeleteThread(ctx, req.(*DeleteThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_ListComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).ListComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/ListComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).ListComments(ctx, req.(*ListCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_SetMarker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMarkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).SetMarker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/SetMarker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).SetMarker(ctx, req.(*SetMarkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_UnsetMarker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsetMarkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).UnsetMarker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumService/UnsetMarker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).UnsetMarker(ctx, req.(*UnsetMarkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ForumService_ServiceDesc is the grpc.ServiceDesc for ForumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ForumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mtechnavi.api.forum.ForumService",
	HandlerType: (*ForumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBaseThread",
			Handler:    _ForumService_CreateBaseThread_Handler,
		},
		{
			MethodName: "ListBaseThreads",
			Handler:    _ForumService_ListBaseThreads_Handler,
		},
		{
			MethodName: "GetBaseThread",
			Handler:    _ForumService_GetBaseThread_Handler,
		},
		{
			MethodName: "UpdateBaseThread",
			Handler:    _ForumService_UpdateBaseThread_Handler,
		},
		{
			MethodName: "DeleteBaseThread",
			Handler:    _ForumService_DeleteBaseThread_Handler,
		},
		{
			MethodName: "CreateThread",
			Handler:    _ForumService_CreateThread_Handler,
		},
		{
			MethodName: "ListThreads",
			Handler:    _ForumService_ListThreads_Handler,
		},
		{
			MethodName: "GetThread",
			Handler:    _ForumService_GetThread_Handler,
		},
		{
			MethodName: "UpdateThread",
			Handler:    _ForumService_UpdateThread_Handler,
		},
		{
			MethodName: "DeleteThread",
			Handler:    _ForumService_DeleteThread_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _ForumService_CreateComment_Handler,
		},
		{
			MethodName: "ListComments",
			Handler:    _ForumService_ListComments_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _ForumService_GetComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _ForumService_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _ForumService_DeleteComment_Handler,
		},
		{
			MethodName: "SetMarker",
			Handler:    _ForumService_SetMarker_Handler,
		},
		{
			MethodName: "UnsetMarker",
			Handler:    _ForumService_UnsetMarker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "forum.proto",
}

// ForumTaskListServiceClient is the client API for ForumTaskListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ForumTaskListServiceClient interface {
	// 伝票未読
	UnreadListResources(ctx context.Context, in *UnreadListResourcesRequest, opts ...grpc.CallOption) (*UnreadListResourcesResponse, error)
}

type forumTaskListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewForumTaskListServiceClient(cc grpc.ClientConnInterface) ForumTaskListServiceClient {
	return &forumTaskListServiceClient{cc}
}

func (c *forumTaskListServiceClient) UnreadListResources(ctx context.Context, in *UnreadListResourcesRequest, opts ...grpc.CallOption) (*UnreadListResourcesResponse, error) {
	out := new(UnreadListResourcesResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.forum.ForumTaskListService/UnreadListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForumTaskListServiceServer is the server API for ForumTaskListService service.
// All implementations must embed UnimplementedForumTaskListServiceServer
// for forward compatibility
type ForumTaskListServiceServer interface {
	// 伝票未読
	UnreadListResources(context.Context, *UnreadListResourcesRequest) (*UnreadListResourcesResponse, error)
	mustEmbedUnimplementedForumTaskListServiceServer()
}

// UnimplementedForumTaskListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedForumTaskListServiceServer struct {
}

func (UnimplementedForumTaskListServiceServer) UnreadListResources(context.Context, *UnreadListResourcesRequest) (*UnreadListResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnreadListResources not implemented")
}
func (UnimplementedForumTaskListServiceServer) mustEmbedUnimplementedForumTaskListServiceServer() {}

// UnsafeForumTaskListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForumTaskListServiceServer will
// result in compilation errors.
type UnsafeForumTaskListServiceServer interface {
	mustEmbedUnimplementedForumTaskListServiceServer()
}

func RegisterForumTaskListServiceServer(s grpc.ServiceRegistrar, srv ForumTaskListServiceServer) {
	s.RegisterService(&ForumTaskListService_ServiceDesc, srv)
}

func _ForumTaskListService_UnreadListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnreadListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumTaskListServiceServer).UnreadListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.forum.ForumTaskListService/UnreadListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumTaskListServiceServer).UnreadListResources(ctx, req.(*UnreadListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ForumTaskListService_ServiceDesc is the grpc.ServiceDesc for ForumTaskListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ForumTaskListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mtechnavi.api.forum.ForumTaskListService",
	HandlerType: (*ForumTaskListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnreadListResources",
			Handler:    _ForumTaskListService_UnreadListResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "forum.proto",
}
