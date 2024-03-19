package event

import (
	"context"
	pb "mtechnavi/sharelib/event/protobuf"
)

type Transport interface {
	Publish(context.Context, *pb.Event) error
	Subscribe(context.Context) (<-chan *pb.Event, <-chan error)
}
