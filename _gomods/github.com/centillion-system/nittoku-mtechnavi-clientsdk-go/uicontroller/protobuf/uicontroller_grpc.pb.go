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

// UiControllerClient is the client API for UiController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UiControllerClient interface {
	CreatePresetMenu(ctx context.Context, in *CreatePresetMenuRequest, opts ...grpc.CallOption) (*PresetMenu, error)
	ListPresetMenus(ctx context.Context, in *ListPresetMenusRequest, opts ...grpc.CallOption) (*ListPresetMenusResponse, error)
	GetPresetMenu(ctx context.Context, in *GetPresetMenuRequest, opts ...grpc.CallOption) (*PresetMenu, error)
	UpdatePresetMenu(ctx context.Context, in *UpdatePresetMenuRequest, opts ...grpc.CallOption) (*PresetMenu, error)
	DeletePresetMenu(ctx context.Context, in *DeletePresetMenuRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreatePreset(ctx context.Context, in *CreatePresetRequest, opts ...grpc.CallOption) (*Preset, error)
	ListPresets(ctx context.Context, in *ListPresetsRequest, opts ...grpc.CallOption) (*ListPresetsResponse, error)
	GetPreset(ctx context.Context, in *GetPresetRequest, opts ...grpc.CallOption) (*Preset, error)
	UpdatePreset(ctx context.Context, in *UpdatePresetRequest, opts ...grpc.CallOption) (*Preset, error)
	DeletePreset(ctx context.Context, in *DeletePresetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type uiControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewUiControllerClient(cc grpc.ClientConnInterface) UiControllerClient {
	return &uiControllerClient{cc}
}

func (c *uiControllerClient) CreatePresetMenu(ctx context.Context, in *CreatePresetMenuRequest, opts ...grpc.CallOption) (*PresetMenu, error) {
	out := new(PresetMenu)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/CreatePresetMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uiControllerClient) ListPresetMenus(ctx context.Context, in *ListPresetMenusRequest, opts ...grpc.CallOption) (*ListPresetMenusResponse, error) {
	out := new(ListPresetMenusResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/ListPresetMenus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uiControllerClient) GetPresetMenu(ctx context.Context, in *GetPresetMenuRequest, opts ...grpc.CallOption) (*PresetMenu, error) {
	out := new(PresetMenu)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/GetPresetMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uiControllerClient) UpdatePresetMenu(ctx context.Context, in *UpdatePresetMenuRequest, opts ...grpc.CallOption) (*PresetMenu, error) {
	out := new(PresetMenu)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/UpdatePresetMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uiControllerClient) DeletePresetMenu(ctx context.Context, in *DeletePresetMenuRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/DeletePresetMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uiControllerClient) CreatePreset(ctx context.Context, in *CreatePresetRequest, opts ...grpc.CallOption) (*Preset, error) {
	out := new(Preset)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/CreatePreset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uiControllerClient) ListPresets(ctx context.Context, in *ListPresetsRequest, opts ...grpc.CallOption) (*ListPresetsResponse, error) {
	out := new(ListPresetsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/ListPresets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uiControllerClient) GetPreset(ctx context.Context, in *GetPresetRequest, opts ...grpc.CallOption) (*Preset, error) {
	out := new(Preset)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/GetPreset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uiControllerClient) UpdatePreset(ctx context.Context, in *UpdatePresetRequest, opts ...grpc.CallOption) (*Preset, error) {
	out := new(Preset)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/UpdatePreset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uiControllerClient) DeletePreset(ctx context.Context, in *DeletePresetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.uicontroller.UiController/DeletePreset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UiControllerServer is the server API for UiController service.
// All implementations must embed UnimplementedUiControllerServer
// for forward compatibility
type UiControllerServer interface {
	CreatePresetMenu(context.Context, *CreatePresetMenuRequest) (*PresetMenu, error)
	ListPresetMenus(context.Context, *ListPresetMenusRequest) (*ListPresetMenusResponse, error)
	GetPresetMenu(context.Context, *GetPresetMenuRequest) (*PresetMenu, error)
	UpdatePresetMenu(context.Context, *UpdatePresetMenuRequest) (*PresetMenu, error)
	DeletePresetMenu(context.Context, *DeletePresetMenuRequest) (*emptypb.Empty, error)
	CreatePreset(context.Context, *CreatePresetRequest) (*Preset, error)
	ListPresets(context.Context, *ListPresetsRequest) (*ListPresetsResponse, error)
	GetPreset(context.Context, *GetPresetRequest) (*Preset, error)
	UpdatePreset(context.Context, *UpdatePresetRequest) (*Preset, error)
	DeletePreset(context.Context, *DeletePresetRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUiControllerServer()
}

// UnimplementedUiControllerServer must be embedded to have forward compatible implementations.
type UnimplementedUiControllerServer struct {
}

func (UnimplementedUiControllerServer) CreatePresetMenu(context.Context, *CreatePresetMenuRequest) (*PresetMenu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePresetMenu not implemented")
}
func (UnimplementedUiControllerServer) ListPresetMenus(context.Context, *ListPresetMenusRequest) (*ListPresetMenusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPresetMenus not implemented")
}
func (UnimplementedUiControllerServer) GetPresetMenu(context.Context, *GetPresetMenuRequest) (*PresetMenu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPresetMenu not implemented")
}
func (UnimplementedUiControllerServer) UpdatePresetMenu(context.Context, *UpdatePresetMenuRequest) (*PresetMenu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePresetMenu not implemented")
}
func (UnimplementedUiControllerServer) DeletePresetMenu(context.Context, *DeletePresetMenuRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePresetMenu not implemented")
}
func (UnimplementedUiControllerServer) CreatePreset(context.Context, *CreatePresetRequest) (*Preset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePreset not implemented")
}
func (UnimplementedUiControllerServer) ListPresets(context.Context, *ListPresetsRequest) (*ListPresetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPresets not implemented")
}
func (UnimplementedUiControllerServer) GetPreset(context.Context, *GetPresetRequest) (*Preset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreset not implemented")
}
func (UnimplementedUiControllerServer) UpdatePreset(context.Context, *UpdatePresetRequest) (*Preset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePreset not implemented")
}
func (UnimplementedUiControllerServer) DeletePreset(context.Context, *DeletePresetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePreset not implemented")
}
func (UnimplementedUiControllerServer) mustEmbedUnimplementedUiControllerServer() {}

// UnsafeUiControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UiControllerServer will
// result in compilation errors.
type UnsafeUiControllerServer interface {
	mustEmbedUnimplementedUiControllerServer()
}

func RegisterUiControllerServer(s grpc.ServiceRegistrar, srv UiControllerServer) {
	s.RegisterService(&UiController_ServiceDesc, srv)
}

func _UiController_CreatePresetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePresetMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).CreatePresetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/CreatePresetMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).CreatePresetMenu(ctx, req.(*CreatePresetMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UiController_ListPresetMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPresetMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).ListPresetMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/ListPresetMenus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).ListPresetMenus(ctx, req.(*ListPresetMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UiController_GetPresetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPresetMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).GetPresetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/GetPresetMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).GetPresetMenu(ctx, req.(*GetPresetMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UiController_UpdatePresetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePresetMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).UpdatePresetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/UpdatePresetMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).UpdatePresetMenu(ctx, req.(*UpdatePresetMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UiController_DeletePresetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePresetMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).DeletePresetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/DeletePresetMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).DeletePresetMenu(ctx, req.(*DeletePresetMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UiController_CreatePreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).CreatePreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/CreatePreset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).CreatePreset(ctx, req.(*CreatePresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UiController_ListPresets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPresetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).ListPresets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/ListPresets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).ListPresets(ctx, req.(*ListPresetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UiController_GetPreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).GetPreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/GetPreset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).GetPreset(ctx, req.(*GetPresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UiController_UpdatePreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).UpdatePreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/UpdatePreset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).UpdatePreset(ctx, req.(*UpdatePresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UiController_DeletePreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UiControllerServer).DeletePreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.uicontroller.UiController/DeletePreset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UiControllerServer).DeletePreset(ctx, req.(*DeletePresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UiController_ServiceDesc is the grpc.ServiceDesc for UiController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UiController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mtechnavi.api.uicontroller.UiController",
	HandlerType: (*UiControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePresetMenu",
			Handler:    _UiController_CreatePresetMenu_Handler,
		},
		{
			MethodName: "ListPresetMenus",
			Handler:    _UiController_ListPresetMenus_Handler,
		},
		{
			MethodName: "GetPresetMenu",
			Handler:    _UiController_GetPresetMenu_Handler,
		},
		{
			MethodName: "UpdatePresetMenu",
			Handler:    _UiController_UpdatePresetMenu_Handler,
		},
		{
			MethodName: "DeletePresetMenu",
			Handler:    _UiController_DeletePresetMenu_Handler,
		},
		{
			MethodName: "CreatePreset",
			Handler:    _UiController_CreatePreset_Handler,
		},
		{
			MethodName: "ListPresets",
			Handler:    _UiController_ListPresets_Handler,
		},
		{
			MethodName: "GetPreset",
			Handler:    _UiController_GetPreset_Handler,
		},
		{
			MethodName: "UpdatePreset",
			Handler:    _UiController_UpdatePreset_Handler,
		},
		{
			MethodName: "DeletePreset",
			Handler:    _UiController_DeletePreset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uicontroller.proto",
}
