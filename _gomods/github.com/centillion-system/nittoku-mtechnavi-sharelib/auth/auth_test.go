package auth

import (
	pb "mtechnavi/idp/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsAnyGroupID(t *testing.T) {

	cases := []struct {
		inUserGroups    []*pb.UserGroup
		inTargetUserIDs []string
		exResult        bool
	}{
		{
			inUserGroups:    nil,
			inTargetUserIDs: nil,
			exResult:        false,
		},
		{
			inUserGroups:    []*pb.UserGroup{},
			inTargetUserIDs: []string{},
			exResult:        false,
		},
		{
			inUserGroups:    nil,
			inTargetUserIDs: []string{},
			exResult:        false,
		},
		{
			inUserGroups:    []*pb.UserGroup{},
			inTargetUserIDs: nil,
			exResult:        false,
		},
		{
			inUserGroups: []*pb.UserGroup{
				{UserGroupId: "UGI0001"},
				{UserGroupId: "UGI0002"},
			},
			inTargetUserIDs: []string{
				"UGI0001",
			},
			exResult: true,
		},
		{
			inUserGroups: []*pb.UserGroup{
				{UserGroupId: "UGI0001"},
				{UserGroupId: "UGI0002"},
			},
			inTargetUserIDs: []string{
				"UGI0002",
			},
			exResult: true,
		},
		{
			inUserGroups: []*pb.UserGroup{
				{UserGroupId: "UGI0001"},
				{UserGroupId: "UGI0002"},
			},
			inTargetUserIDs: []string{
				"UGI0003",
			},
			exResult: false,
		},
	}

	for i, c := range cases {
		actual := ContainsAnyGroupID(c.inUserGroups, c.inTargetUserIDs)
		assert.Equalf(t, actual, c.exResult, "Case Number:%d,inUserGroups:%v,inTargetUserIDs:%v", i, c.inUserGroups, c.inTargetUserIDs)
	}
}

func TestIsTenantAdmin(t *testing.T) {

	cases := []struct {
		inUserGroups []*pb.UserGroup
		exResult     bool
	}{
		{
			inUserGroups: nil,
			exResult:     false,
		},
		{
			inUserGroups: []*pb.UserGroup{},
			exResult:     false,
		},
		{
			inUserGroups: nil,
			exResult:     false,
		},
		{
			inUserGroups: []*pb.UserGroup{},
			exResult:     false,
		},
		{
			inUserGroups: []*pb.UserGroup{
				{UserGroupCode: "UGC0001"},
				{UserGroupCode: "Group-Tenant-Admin"},
			},
			exResult: true,
		},
		{
			inUserGroups: []*pb.UserGroup{
				{UserGroupCode: "Group-Tenant-Admin"},
				{UserGroupCode: "UGC0002"},
			},
			exResult: true,
		},
		{
			inUserGroups: []*pb.UserGroup{
				{UserGroupCode: "UGC0001"},
				{UserGroupCode: "UGC0002"},
			},
			exResult: false,
		},
	}

	for i, c := range cases {
		actual := IsTenantAdmin(c.inUserGroups)
		assert.Equalf(t, actual, c.exResult, "Case Number:%d,inUserGroups:%v", i, c.inUserGroups)
	}
}
