package visibility_test

import (
	"encoding/json"
	"mtechnavi/sharelib/dataproxy/codec"
	sharelibpb "mtechnavi/sharelib/protobuf"
	mtnpb "mtechnavi/sharelib/protobuf/mtn"
	dataproxypb "mtechnavi/sharelib/protobuf/record"
	"mtechnavi/sharelib/visibility"
	testdatapb "mtechnavi/sharelib/visibility/testdata/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestNewScopedMessage_Proto(t *testing.T) {
	type args struct {
		scope mtnpb.Scope
		msg   proto.Message
	}

	tests := []struct {
		name string
		args args
		want proto.Message
	}{
		{
			name: "zero value",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.SimpleVisibilityTypes{},
			},
			want: &testdatapb.SimpleVisibilityTypes{},
		},
		{
			name: "primary field",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.SimpleVisibilityTypes{
					SimpleVisibilityTypesId: "SimpleVisibilityTypesId001",
					CreatedAt:               1,
					UpdatedAt:               2,
					DeletedAt:               3,
					Public:                  "public",
					Granted:                 "granted",
				},
			},
			want: &testdatapb.SimpleVisibilityTypes{
				SimpleVisibilityTypesId: "SimpleVisibilityTypesId001",
				CreatedAt:               1,
				UpdatedAt:               2,
				DeletedAt:               3,
				Public:                  "public",
			},
		},
		{
			name: "public",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
			},
			want: &testdatapb.SimpleVisibilityTypes{
				Public: "public",
			},
		},
		{
			name: "granted",
			args: args{
				mtnpb.Scope_GRANTED,
				&testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
			},
			want: &testdatapb.SimpleVisibilityTypes{
				Public:  "public",
				Granted: "granted",
			},
		},
		{
			name: "private",
			args: args{
				mtnpb.Scope_PRIVATE,
				&testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
			},
			want: &testdatapb.SimpleVisibilityTypes{
				Public:   "public",
				Granted:  "granted",
				Private1: "private1",
				Private2: "private2",
			},
		},
		{
			name: "nested zero value",
			args: args{
				mtnpb.Scope_GRANTED,
				&testdatapb.NestedVisibilityTypes{},
			},
			want: &testdatapb.NestedVisibilityTypes{},
		},
		{
			name: "nested primary field",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.NestedVisibilityTypes{
					NestedVisibilityTypesId: "NestedVisibilityTypesId001",
					CreatedAt:               1,
					UpdatedAt:               2,
					DeletedAt:               3,
					SharedProperties: &sharelibpb.EmbeddedSharedProperties{
						SharedBy: "test001",
						SharedAt: 1,
					},

					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						SimpleVisibilityTypesId: "PublicMessage",
						CreatedAt:               11,
						UpdatedAt:               12,
						DeletedAt:               13,
					},
					GrantedNestedMessage: &testdatapb.NestedMessage{
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							SimpleVisibilityTypesId: "GrantedNestedMessage.PublicMessage",
							CreatedAt:               21,
							UpdatedAt:               22,
							DeletedAt:               23,
						},
						SharedProperties: &sharelibpb.EmbeddedSharedProperties{
							SharedBy: "test002",
							SharedAt: 2,
						},
					},
					PublicNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							SimpleVisibilityTypesId: "PublicNestedMessage.PrivateMessage2",
							CreatedAt:               31,
							UpdatedAt:               32,
							DeletedAt:               33,
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							SimpleVisibilityTypesId: "PublicNestedMessage.PublicMessage",
							CreatedAt:               41,
							UpdatedAt:               42,
							DeletedAt:               43,
						},
						SharedProperties: &sharelibpb.EmbeddedSharedProperties{
							SharedBy: "test003",
							SharedAt: 3,
						},
					},
				},
			},
			want: &testdatapb.NestedVisibilityTypes{
				NestedVisibilityTypesId: "NestedVisibilityTypesId001",
				CreatedAt:               1,
				UpdatedAt:               2,
				DeletedAt:               3,
				SharedProperties: &sharelibpb.EmbeddedSharedProperties{
					SharedBy: "test001",
					SharedAt: 1,
				},
				PublicMessage: &testdatapb.SimpleVisibilityTypes{
					SimpleVisibilityTypesId: "PublicMessage",
					CreatedAt:               11,
					UpdatedAt:               12,
					DeletedAt:               13,
				},
				GrantedNestedMessage: &testdatapb.NestedMessage{},
				PublicNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						SimpleVisibilityTypesId: "PublicNestedMessage.PublicMessage",
						CreatedAt:               41,
						UpdatedAt:               42,
						DeletedAt:               43,
					},
					SharedProperties: &sharelibpb.EmbeddedSharedProperties{
						SharedBy: "test003",
						SharedAt: 3,
					},
				},
			},
		},
		{
			name: "nested public",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.NestedVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
						Message3: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage1: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage2: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					GrantedNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PublicNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateEmptyMessage1: &testdatapb.EmptyMessage{},
					PrivateEmptyMessage2: &testdatapb.EmptyMessage{},
					GrantedEmptyMessage:  &testdatapb.EmptyMessage{},
					PublicEmptyMessage:   &testdatapb.EmptyMessage{},
				},
			},
			want: &testdatapb.NestedVisibilityTypes{
				Public:          "public",
				PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
				PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
				GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
				PublicMessage: &testdatapb.SimpleVisibilityTypes{
					Public: "public",
				},
				TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
					Message3: &testdatapb.SimpleVisibilityTypes{
						Public: "public",
					},
				},
				PrivateNestedMessage1: &testdatapb.NestedMessage{},
				PrivateNestedMessage2: &testdatapb.NestedMessage{},
				GrantedNestedMessage:  &testdatapb.NestedMessage{},
				PublicNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public: "public",
					},
				},
				PrivateEmptyMessage1: &testdatapb.EmptyMessage{},
				PrivateEmptyMessage2: &testdatapb.EmptyMessage{},
				GrantedEmptyMessage:  &testdatapb.EmptyMessage{},
				PublicEmptyMessage:   &testdatapb.EmptyMessage{},
			},
		},
		{
			name: "nested granted",
			args: args{
				mtnpb.Scope_GRANTED,
				&testdatapb.NestedVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
						Message3: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage1: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage2: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					GrantedNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PublicNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
				},
			},
			want: &testdatapb.NestedVisibilityTypes{
				Granted:         "granted",
				Public:          "public",
				PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
				PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
				GrantedMessage: &testdatapb.SimpleVisibilityTypes{
					Granted: "granted",
					Public:  "public",
				},
				PublicMessage: &testdatapb.SimpleVisibilityTypes{
					Granted: "granted",
					Public:  "public",
				},
				TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
					Message3: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
				},
				PrivateNestedMessage1: &testdatapb.NestedMessage{},
				PrivateNestedMessage2: &testdatapb.NestedMessage{},
				GrantedNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
				},
				PublicNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
				},
			},
		},
		{
			name: "nested private",
			args: args{
				mtnpb.Scope_PRIVATE,
				&testdatapb.NestedVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
						Message3: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage1: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage2: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					GrantedNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PublicNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
				},
			},
			want: &testdatapb.NestedVisibilityTypes{
				Public:   "public",
				Granted:  "granted",
				Private1: "private1",
				Private2: "private2",
				PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
				PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
				GrantedMessage: &testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
				PublicMessage: &testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
				TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
					Message3: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				PrivateNestedMessage1: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				PrivateNestedMessage2: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				GrantedNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				PublicNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				// PrivateEmptyMessage1: &testdatapb.EmptyMessage{},
				// PrivateEmptyMessage2: &testdatapb.EmptyMessage{},
				// GrantedEmptyMessage:  &testdatapb.EmptyMessage{},
				// PublicEmptyMessage:   &testdatapb.EmptyMessage{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := visibility.NewScopedMessage(tt.args.scope, tt.args.msg)
			assert.JSONEq(t, jsonStr(tt.want), jsonStr(got))

			got = visibility.NewScopedMessage(tt.args.scope, got)
			assert.JSONEq(t, jsonStr(tt.want), jsonStr(got))
		})
	}
}

func TestNewScopedMessage_Record(t *testing.T) {
	type args struct {
		scope mtnpb.Scope
		msg   proto.Message
	}
	type simpleTest struct {
		name string
		args args
		want *testdatapb.SimpleVisibilityTypes
	}
	simpleTests := []simpleTest{
		{
			name: "zero value",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.SimpleVisibilityTypes{},
			},
			want: &testdatapb.SimpleVisibilityTypes{},
		},
		{
			name: "primary field",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.SimpleVisibilityTypes{
					SimpleVisibilityTypesId: "SimpleVisibilityTypesId001",
					CreatedAt:               1,
					UpdatedAt:               2,
					DeletedAt:               3,
					Public:                  "public",
					Granted:                 "granted",
				},
			},
			want: &testdatapb.SimpleVisibilityTypes{
				SimpleVisibilityTypesId: "SimpleVisibilityTypesId001",
				CreatedAt:               1,
				UpdatedAt:               2,
				DeletedAt:               3,
				Public:                  "public",
			},
		},
		{
			name: "public",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
			},
			want: &testdatapb.SimpleVisibilityTypes{
				Public: "public",
			},
		},
		{
			name: "granted",
			args: args{
				mtnpb.Scope_GRANTED,
				&testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
			},
			want: &testdatapb.SimpleVisibilityTypes{
				Public:  "public",
				Granted: "granted",
			},
		},
		{
			name: "private",
			args: args{
				mtnpb.Scope_PRIVATE,
				&testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
			},
			want: &testdatapb.SimpleVisibilityTypes{
				Public:   "public",
				Granted:  "granted",
				Private1: "private1",
				Private2: "private2",
			},
		},
	}
	for _, tt := range simpleTests {
		t.Run(tt.name, func(t *testing.T) {
			// encode
			srcRec, err := codec.Encode(tt.args.msg)
			if !assert.NoError(t, err) {
				return
			}
			expRec, err := codec.Encode(tt.want)
			if !assert.NoError(t, err) {
				return
			}
			// scoped
			assert.Equal(t, "mtechnavi.visibility.testdata.SimpleVisibilityTypes", srcRec.TypeName)
			assert.Len(t, srcRec.Fields, 8)
			got := visibility.NewScopedMessage(tt.args.scope, srcRec)
			assert.JSONEq(t, jsonStr(&expRec), jsonStr(got.(*dataproxypb.Record)))
			// decode
			var dst testdatapb.SimpleVisibilityTypes
			if assert.NoError(t, codec.Decode(&dst, got.(*dataproxypb.Record))) {
				// 内部フィールドなどは比較したくないため、一度JSON化して比較する
				assert.JSONEq(t, jsonStr(&tt.want), jsonStr(&dst))
			}
		})
	}

	type nestedTest struct {
		name string
		args args
		want *testdatapb.NestedVisibilityTypes
	}
	nestedTests := []nestedTest{
		{
			name: "nested zero value",
			args: args{
				mtnpb.Scope_GRANTED,
				&testdatapb.NestedVisibilityTypes{},
			},
			want: &testdatapb.NestedVisibilityTypes{
				PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
				PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
				GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
				PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				PrivateNestedMessage1: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				PrivateNestedMessage2: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				GrantedNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				PublicNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
			},
		},
		{
			name: "nested primary field",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.NestedVisibilityTypes{
					NestedVisibilityTypesId: "NestedVisibilityTypesId001",
					CreatedAt:               1,
					UpdatedAt:               2,
					DeletedAt:               3,
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						SimpleVisibilityTypesId: "PublicMessage",
						CreatedAt:               11,
						UpdatedAt:               12,
						DeletedAt:               13,
					},
					GrantedNestedMessage: &testdatapb.NestedMessage{
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							SimpleVisibilityTypesId: "GrantedNestedMessage.PublicMessage",
							CreatedAt:               21,
							UpdatedAt:               22,
							DeletedAt:               23,
						},
					},
					PublicNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							SimpleVisibilityTypesId: "PublicNestedMessage.PrivateMessage2",
							CreatedAt:               31,
							UpdatedAt:               32,
							DeletedAt:               33,
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							SimpleVisibilityTypesId: "SimpleVisibilityTypesId",
							CreatedAt:               41,
							UpdatedAt:               42,
							DeletedAt:               43,
						},
					},
				},
			},
			want: &testdatapb.NestedVisibilityTypes{
				NestedVisibilityTypesId: "NestedVisibilityTypesId001",
				CreatedAt:               1,
				UpdatedAt:               2,
				DeletedAt:               3,
				PrivateMessage1:         &testdatapb.SimpleVisibilityTypes{},
				PrivateMessage2:         &testdatapb.SimpleVisibilityTypes{},
				GrantedMessage:          &testdatapb.SimpleVisibilityTypes{},
				PublicMessage: &testdatapb.SimpleVisibilityTypes{
					SimpleVisibilityTypesId: "PublicMessage",
					CreatedAt:               11,
					UpdatedAt:               12,
					DeletedAt:               13,
				},
				PrivateNestedMessage1: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				PrivateNestedMessage2: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				GrantedNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				PublicNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						SimpleVisibilityTypesId: "SimpleVisibilityTypesId",
						CreatedAt:               41,
						UpdatedAt:               42,
						DeletedAt:               43,
					},
				},
			},
		},
		{
			name: "nested public",
			args: args{
				mtnpb.Scope_PUBLIC,
				&testdatapb.NestedVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
						Message3: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage1: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage2: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					GrantedNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PublicNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
				},
			},
			want: &testdatapb.NestedVisibilityTypes{
				Public:          "public",
				PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
				PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
				GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
				PublicMessage: &testdatapb.SimpleVisibilityTypes{
					Public: "public",
				},
				TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
					Message3: &testdatapb.SimpleVisibilityTypes{
						Public: "public",
					},
				},
				PrivateNestedMessage1: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				PrivateNestedMessage2: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				GrantedNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				PublicNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public: "public",
					},
				},
			},
		},
		{
			name: "nested granted",
			args: args{
				mtnpb.Scope_GRANTED,
				&testdatapb.NestedVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
						Message3: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage1: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage2: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					GrantedNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PublicNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
				},
			},
			want: &testdatapb.NestedVisibilityTypes{
				Granted:         "granted",
				Public:          "public",
				PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
				PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
				GrantedMessage: &testdatapb.SimpleVisibilityTypes{
					Granted: "granted",
					Public:  "public",
				},
				PublicMessage: &testdatapb.SimpleVisibilityTypes{
					Granted: "granted",
					Public:  "public",
				},
				TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
					Message3: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
				},

				PrivateNestedMessage1: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				PrivateNestedMessage2: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage:  &testdatapb.SimpleVisibilityTypes{},
					PublicMessage:   &testdatapb.SimpleVisibilityTypes{},
				},
				GrantedNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
				},
				PublicNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Granted: "granted",
						Public:  "public",
					},
				},
			},
		},
		{
			name: "nested private",
			args: args{
				mtnpb.Scope_PRIVATE,
				&testdatapb.NestedVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
						Message3: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage1: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PrivateNestedMessage2: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					GrantedNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
					PublicNestedMessage: &testdatapb.NestedMessage{
						PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						GrantedMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
						PublicMessage: &testdatapb.SimpleVisibilityTypes{
							Public:   "public",
							Granted:  "granted",
							Private1: "private1",
							Private2: "private2",
						},
					},
				},
			},
			want: &testdatapb.NestedVisibilityTypes{
				Public:   "public",
				Granted:  "granted",
				Private1: "private1",
				Private2: "private2",
				PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
				PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
				GrantedMessage: &testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
				PublicMessage: &testdatapb.SimpleVisibilityTypes{
					Public:   "public",
					Granted:  "granted",
					Private1: "private1",
					Private2: "private2",
				},
				TestOneof: &testdatapb.NestedVisibilityTypes_Message3{
					Message3: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				PrivateNestedMessage1: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				PrivateNestedMessage2: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				GrantedNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				PublicNestedMessage: &testdatapb.NestedMessage{
					PrivateMessage1: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PrivateMessage2: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					GrantedMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
					PublicMessage: &testdatapb.SimpleVisibilityTypes{
						Public:   "public",
						Granted:  "granted",
						Private1: "private1",
						Private2: "private2",
					},
				},
				// PrivateEmptyMessage1: &testdatapb.EmptyMessage{},
				// PrivateEmptyMessage2: &testdatapb.EmptyMessage{},
				// GrantedEmptyMessage:  &testdatapb.EmptyMessage{},
				// PublicEmptyMessage:   &testdatapb.EmptyMessage{},
			},
		},
	}
	for _, tt := range nestedTests {
		t.Run(tt.name, func(t *testing.T) {
			// encode
			srcRec, err := codec.Encode(tt.args.msg)
			if !assert.NoError(t, err) {
				return
			}
			expRec, err := codec.Encode(tt.want)
			if !assert.NoError(t, err) {
				return
			}
			// scoped
			assert.Equal(t, "mtechnavi.visibility.testdata.NestedVisibilityTypes", srcRec.TypeName)
			assert.Len(t, srcRec.Fields, 24)
			got := visibility.NewScopedMessage(tt.args.scope, srcRec)
			assert.JSONEq(t, jsonStr(&expRec), jsonStr(got.(*dataproxypb.Record)))

			got = visibility.NewScopedMessage(tt.args.scope, got)
			assert.JSONEq(t, jsonStr(&expRec), jsonStr(got.(*dataproxypb.Record)))
			// decode
			var dst testdatapb.NestedVisibilityTypes
			if assert.NoError(t, codec.Decode(&dst, got.(*dataproxypb.Record))) {
				// omit zero value by encode
				tt.want.PrivateNestedMessage1.SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
				tt.want.PrivateNestedMessage2.SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
				tt.want.GrantedNestedMessage.SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
				tt.want.PublicNestedMessage.SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
				tt.want.SharedProperties = &sharelibpb.EmbeddedSharedProperties{}
				// 内部フィールドなどは比較したくないため、一度JSON化して比較する
				assert.JSONEq(t, jsonStr(&tt.want), jsonStr(&dst))
			}
		})
	}
}

func jsonStr(v interface{}) string {
	return string(jsonBytes(v))
}

func jsonBytes(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
