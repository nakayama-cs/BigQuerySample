package to

import (
	"time"

	pb "mtechnavi/sharelib/protobuf"
)

func ToEmbeddedUser(src interface {
	GetDisplayName() string
	GetEmail() string
}) *pb.EmbeddedUser {
	return &pb.EmbeddedUser{
		DisplayName: src.GetDisplayName(),
		Email:       src.GetEmail(),
	}
}

func ToEmbeddedUpdatedProperties(t time.Time, user interface {
	GetDisplayName() string
	GetEmail() string
}) *pb.EmbeddedUpdatedProperties {
	return &pb.EmbeddedUpdatedProperties{
		UpdatedAt: t.UnixMicro(),
		UpdatedBy: ToEmbeddedUser(user),
	}
}
