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

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityClient interface {
	// 認証前に実行
	CreateBrowserSession(ctx context.Context, in *CreateBrowserSessionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 認証前に実行
	DeleteBrowserSession(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 認証前に実行
	GetTenantConfig(ctx context.Context, in *GetTenantConfigRequest, opts ...grpc.CallOption) (*TenantConfig, error)
	// 認証前に実行
	VerifyUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VerifyUserResponse, error)
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateUserGroup(ctx context.Context, in *CreateUserGroupRequest, opts ...grpc.CallOption) (*UserGroup, error)
	ListUserGroups(ctx context.Context, in *ListUserGroupsRequest, opts ...grpc.CallOption) (*ListUserGroupsResponse, error)
	GetUserGroup(ctx context.Context, in *GetUserGroupRequest, opts ...grpc.CallOption) (*UserGroup, error)
	UpdateUserGroup(ctx context.Context, in *UpdateUserGroupRequest, opts ...grpc.CallOption) (*UserGroup, error)
	DeleteUserGroup(ctx context.Context, in *DeleteUserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error)
	ListSchemaAccessControls(ctx context.Context, in *ListSchemaAccessControlsRequest, opts ...grpc.CallOption) (*ListSchemaAccessControlsResponse, error)
	GetSchemaAccessControl(ctx context.Context, in *GetSchemaAccessControlRequest, opts ...grpc.CallOption) (*SchemaAccessControl, error)
}

type identityClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityClient(cc grpc.ClientConnInterface) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) CreateBrowserSession(ctx context.Context, in *CreateBrowserSessionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/CreateBrowserSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) DeleteBrowserSession(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/DeleteBrowserSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetTenantConfig(ctx context.Context, in *GetTenantConfigRequest, opts ...grpc.CallOption) (*TenantConfig, error) {
	out := new(TenantConfig)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/GetTenantConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) VerifyUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VerifyUserResponse, error) {
	out := new(VerifyUserResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/VerifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/ListTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/UpdateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) CreateUserGroup(ctx context.Context, in *CreateUserGroupRequest, opts ...grpc.CallOption) (*UserGroup, error) {
	out := new(UserGroup)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/CreateUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ListUserGroups(ctx context.Context, in *ListUserGroupsRequest, opts ...grpc.CallOption) (*ListUserGroupsResponse, error) {
	out := new(ListUserGroupsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/ListUserGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetUserGroup(ctx context.Context, in *GetUserGroupRequest, opts ...grpc.CallOption) (*UserGroup, error) {
	out := new(UserGroup)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/GetUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) UpdateUserGroup(ctx context.Context, in *UpdateUserGroupRequest, opts ...grpc.CallOption) (*UserGroup, error) {
	out := new(UserGroup)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/UpdateUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) DeleteUserGroup(ctx context.Context, in *DeleteUserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/DeleteUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/ListRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ListSchemaAccessControls(ctx context.Context, in *ListSchemaAccessControlsRequest, opts ...grpc.CallOption) (*ListSchemaAccessControlsResponse, error) {
	out := new(ListSchemaAccessControlsResponse)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/ListSchemaAccessControls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetSchemaAccessControl(ctx context.Context, in *GetSchemaAccessControlRequest, opts ...grpc.CallOption) (*SchemaAccessControl, error) {
	out := new(SchemaAccessControl)
	err := c.cc.Invoke(ctx, "/mtechnavi.api.idp.Identity/GetSchemaAccessControl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
// All implementations must embed UnimplementedIdentityServer
// for forward compatibility
type IdentityServer interface {
	// 認証前に実行
	CreateBrowserSession(context.Context, *CreateBrowserSessionRequest) (*emptypb.Empty, error)
	// 認証前に実行
	DeleteBrowserSession(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// 認証前に実行
	GetTenantConfig(context.Context, *GetTenantConfigRequest) (*TenantConfig, error)
	// 認証前に実行
	VerifyUser(context.Context, *emptypb.Empty) (*VerifyUserResponse, error)
	CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error)
	ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error)
	DeleteTenant(context.Context, *DeleteTenantRequest) (*emptypb.Empty, error)
	CreateUserGroup(context.Context, *CreateUserGroupRequest) (*UserGroup, error)
	ListUserGroups(context.Context, *ListUserGroupsRequest) (*ListUserGroupsResponse, error)
	GetUserGroup(context.Context, *GetUserGroupRequest) (*UserGroup, error)
	UpdateUserGroup(context.Context, *UpdateUserGroupRequest) (*UserGroup, error)
	DeleteUserGroup(context.Context, *DeleteUserGroupRequest) (*emptypb.Empty, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error)
	GetRole(context.Context, *GetRoleRequest) (*Role, error)
	ListSchemaAccessControls(context.Context, *ListSchemaAccessControlsRequest) (*ListSchemaAccessControlsResponse, error)
	GetSchemaAccessControl(context.Context, *GetSchemaAccessControlRequest) (*SchemaAccessControl, error)
	mustEmbedUnimplementedIdentityServer()
}

// UnimplementedIdentityServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (UnimplementedIdentityServer) CreateBrowserSession(context.Context, *CreateBrowserSessionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrowserSession not implemented")
}
func (UnimplementedIdentityServer) DeleteBrowserSession(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrowserSession not implemented")
}
func (UnimplementedIdentityServer) GetTenantConfig(context.Context, *GetTenantConfigRequest) (*TenantConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantConfig not implemented")
}
func (UnimplementedIdentityServer) VerifyUser(context.Context, *emptypb.Empty) (*VerifyUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}
func (UnimplementedIdentityServer) CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedIdentityServer) ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenants not implemented")
}
func (UnimplementedIdentityServer) GetTenant(context.Context, *GetTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedIdentityServer) UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenant not implemented")
}
func (UnimplementedIdentityServer) DeleteTenant(context.Context, *DeleteTenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (UnimplementedIdentityServer) CreateUserGroup(context.Context, *CreateUserGroupRequest) (*UserGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserGroup not implemented")
}
func (UnimplementedIdentityServer) ListUserGroups(context.Context, *ListUserGroupsRequest) (*ListUserGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserGroups not implemented")
}
func (UnimplementedIdentityServer) GetUserGroup(context.Context, *GetUserGroupRequest) (*UserGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroup not implemented")
}
func (UnimplementedIdentityServer) UpdateUserGroup(context.Context, *UpdateUserGroupRequest) (*UserGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserGroup not implemented")
}
func (UnimplementedIdentityServer) DeleteUserGroup(context.Context, *DeleteUserGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserGroup not implemented")
}
func (UnimplementedIdentityServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedIdentityServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedIdentityServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedIdentityServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedIdentityServer) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedIdentityServer) ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (UnimplementedIdentityServer) GetRole(context.Context, *GetRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedIdentityServer) ListSchemaAccessControls(context.Context, *ListSchemaAccessControlsRequest) (*ListSchemaAccessControlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchemaAccessControls not implemented")
}
func (UnimplementedIdentityServer) GetSchemaAccessControl(context.Context, *GetSchemaAccessControlRequest) (*SchemaAccessControl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchemaAccessControl not implemented")
}
func (UnimplementedIdentityServer) mustEmbedUnimplementedIdentityServer() {}

// UnsafeIdentityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServer will
// result in compilation errors.
type UnsafeIdentityServer interface {
	mustEmbedUnimplementedIdentityServer()
}

func RegisterIdentityServer(s grpc.ServiceRegistrar, srv IdentityServer) {
	s.RegisterService(&Identity_ServiceDesc, srv)
}

func _Identity_CreateBrowserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBrowserSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).CreateBrowserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/CreateBrowserSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).CreateBrowserSession(ctx, req.(*CreateBrowserSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_DeleteBrowserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).DeleteBrowserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/DeleteBrowserSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).DeleteBrowserSession(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetTenantConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetTenantConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/GetTenantConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetTenantConfig(ctx, req.(*GetTenantConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/VerifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).VerifyUser(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/ListTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ListTenants(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/UpdateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).DeleteTenant(ctx, req.(*DeleteTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_CreateUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).CreateUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/CreateUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).CreateUserGroup(ctx, req.(*CreateUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ListUserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ListUserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/ListUserGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ListUserGroups(ctx, req.(*ListUserGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/GetUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetUserGroup(ctx, req.(*GetUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_UpdateUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).UpdateUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/UpdateUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).UpdateUserGroup(ctx, req.(*UpdateUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_DeleteUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).DeleteUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/DeleteUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).DeleteUserGroup(ctx, req.(*DeleteUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/ListRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ListRoles(ctx, req.(*ListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ListSchemaAccessControls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSchemaAccessControlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ListSchemaAccessControls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/ListSchemaAccessControls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ListSchemaAccessControls(ctx, req.(*ListSchemaAccessControlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetSchemaAccessControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaAccessControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetSchemaAccessControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtechnavi.api.idp.Identity/GetSchemaAccessControl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetSchemaAccessControl(ctx, req.(*GetSchemaAccessControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Identity_ServiceDesc is the grpc.ServiceDesc for Identity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Identity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mtechnavi.api.idp.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBrowserSession",
			Handler:    _Identity_CreateBrowserSession_Handler,
		},
		{
			MethodName: "DeleteBrowserSession",
			Handler:    _Identity_DeleteBrowserSession_Handler,
		},
		{
			MethodName: "GetTenantConfig",
			Handler:    _Identity_GetTenantConfig_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _Identity_VerifyUser_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _Identity_CreateTenant_Handler,
		},
		{
			MethodName: "ListTenants",
			Handler:    _Identity_ListTenants_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _Identity_GetTenant_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _Identity_UpdateTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _Identity_DeleteTenant_Handler,
		},
		{
			MethodName: "CreateUserGroup",
			Handler:    _Identity_CreateUserGroup_Handler,
		},
		{
			MethodName: "ListUserGroups",
			Handler:    _Identity_ListUserGroups_Handler,
		},
		{
			MethodName: "GetUserGroup",
			Handler:    _Identity_GetUserGroup_Handler,
		},
		{
			MethodName: "UpdateUserGroup",
			Handler:    _Identity_UpdateUserGroup_Handler,
		},
		{
			MethodName: "DeleteUserGroup",
			Handler:    _Identity_DeleteUserGroup_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Identity_CreateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Identity_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Identity_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Identity_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Identity_DeleteUser_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _Identity_ListRoles_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _Identity_GetRole_Handler,
		},
		{
			MethodName: "ListSchemaAccessControls",
			Handler:    _Identity_ListSchemaAccessControls_Handler,
		},
		{
			MethodName: "GetSchemaAccessControl",
			Handler:    _Identity_GetSchemaAccessControl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}
