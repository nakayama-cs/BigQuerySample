package auth

import (
	pb "mtechnavi/idp/protobuf"
)

type Invoker struct {
	Tenant *pb.Tenant

	User *pb.User

	Groups []*pb.UserGroup

	Roles []*pb.Role
}

func (v *Invoker) GetTenant() *pb.Tenant {
	return v.Tenant
}

func (v *Invoker) GetUser() *pb.User {
	return v.User
}

func (v *Invoker) GetGroups() []*pb.UserGroup {
	return v.Groups
}

func (v *Invoker) GetRoles() []*pb.Role {
	return v.Roles
}
