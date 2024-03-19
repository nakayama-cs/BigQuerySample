package acl

import (
	"encoding/json"
	"fmt"
	pb "mtechnavi/idp/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Example() {
	masker := &Masker{
		AccessControls: &pb.SchemaAccessControl{
			AccessControls: map[string]*pb.PropertyAccessControl{
				"mtechnavi.api.idp.VerifyUserResponse": &pb.PropertyAccessControl{
					Allow: []string{"*"},
				},
				"mtechnavi.api.idp.User": &pb.PropertyAccessControl{
					Allow: []string{"display_name", "email", "email_verified", "user_group_ids"},
					Deny:  []string{"email", "email_verified"},
				},
				"mtechnavi.api.idp.Tenant": &pb.PropertyAccessControl{
					Allow: []string{"*"},
					Deny:  []string{"domain", "internal_tenant_id", "tenant_code"},
				},
			},
		},
	}
	v := &pb.VerifyUserResponse{
		Tenant: &pb.Tenant{
			TenantId:         "xxx",
			TenantCode:       "xxx",
			Domain:           "xxx",
			DisplayName:      "xxx",
			InternalTenantId: "xxx",
		},
		User: &pb.User{
			UserId:        "xxx",
			Email:         "xxx",
			EmailVerified: true,
			Password:      "xxx",
			DisplayName:   "xxx",
			UserGroupIds:  []string{"xxx"},
			Disabled:      false,
		},
		Groups: []*pb.UserGroup{
			&pb.UserGroup{
				UserGroupId:             "xxx",
				UserGroupCode:           "xxx",
				DisplayName:             "xxx",
				RoleNames:               []string{"xxx"},
				SchemaAccessControlName: "xxx",
			},
		},
		Roles: []*pb.Role{
			&pb.Role{
				RoleName:    "xxx",
				DisplayName: "xxx",
				Permissions: []string{"xxx"},
			},
		},
	}
	masker.Apply(v)
	fmt.Println(formatJSON(v))

	// Output:
	// {
	//   "tenant": {
	//     "tenantId": "xxx",
	//     "displayName": "xxx"
	//   },
	//   "user": {
	//     "userId": "xxx",
	//     "displayName": "xxx",
	//     "userGroupIds": [
	//       "xxx"
	//     ]
	//   },
	//   "groups": [
	//     {
	//       "userGroupId": "xxx"
	//     }
	//   ],
	//   "roles": [
	//     {}
	//   ]
	// }
}

func TestMergeAccessControls(t *testing.T) {
	cases := []struct {
		Merged *pb.SchemaAccessControl
		Src    []*pb.SchemaAccessControl
	}{
		{ // 0
			Merged: &pb.SchemaAccessControl{},
		},
		{ // 1
			Merged: &pb.SchemaAccessControl{
				AccessControls: map[string]*pb.PropertyAccessControl{
					"mtechnavi.api.idp.VerifyUserResponse": &pb.PropertyAccessControl{
						Allow: []string{"*"},
					},
					"mtechnavi.api.idp.User": &pb.PropertyAccessControl{
						Allow: []string{"display_name", "email", "email_verified", "user_group_ids"},
						Deny:  []string{"email", "email_verified"},
					},
					"mtechnavi.api.idp.Tenant": &pb.PropertyAccessControl{
						Allow: []string{"*"},
						Deny:  []string{"domain", "internal_tenant_id", "tenant_code"},
					},
				},
			},
			Src: []*pb.SchemaAccessControl{
				&pb.SchemaAccessControl{
					AccessControls: map[string]*pb.PropertyAccessControl{
						"mtechnavi.api.idp.VerifyUserResponse": &pb.PropertyAccessControl{
							Allow: []string{"*"},
						},
					},
				},
				&pb.SchemaAccessControl{
					AccessControls: map[string]*pb.PropertyAccessControl{
						"mtechnavi.api.idp.User": &pb.PropertyAccessControl{
							Allow: []string{"email", "email_verified", "display_name"},
						},
						"mtechnavi.api.idp.Tenant": &pb.PropertyAccessControl{
							Allow: []string{"display_name"},
						},
					},
				},
				&pb.SchemaAccessControl{
					AccessControls: map[string]*pb.PropertyAccessControl{
						"mtechnavi.api.idp.User": &pb.PropertyAccessControl{
							Allow: []string{"user_group_ids"},
							Deny:  []string{"email", "email_verified"},
						},
						"mtechnavi.api.idp.Tenant": &pb.PropertyAccessControl{
							Allow: []string{"*"},
							Deny:  []string{"tenant_code", "domain", "internal_tenant_id"},
						},
					},
				},
			},
		},
	}
	for i, c := range cases {
		v := mergeAccessControls(c.Src...)
		exp := formatJSON(c.Merged)
		act := formatJSON(v)
		assert.JSONEqf(t, exp, act, "case=%d", i)
	}
}

func formatJSON(v interface{}) string {
	if msg, ok := v.(proto.Message); ok {
		v = json.RawMessage(protojson.Format(msg))
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
