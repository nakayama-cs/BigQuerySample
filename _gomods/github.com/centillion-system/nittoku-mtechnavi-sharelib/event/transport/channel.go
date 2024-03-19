package transport

import (
	"context"
	"mtechnavi/sharelib/event"
	pb "mtechnavi/sharelib/event/protobuf"
	"sync"
)

type Channel struct {
	C chan *pb.Event

	mu sync.Mutex
}

func NewChannel() *Channel {
	return &Channel{
		C: make(chan *pb.Event, 1024),
	}
}

func (tp *Channel) Publish(ctx context.Context, evt *pb.Event) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	immediate := (cap(tp.C) - len(tp.C)) > 0
	if immediate {
		tp.C <- evt
	} else {
		go func() {
			tp.C <- evt
		}()
	}
	return nil
}

func (tp *Channel) Subscribe(ctx context.Context) (<-chan *pb.Event, <-chan error) {
	evtCh := make(chan *pb.Event)
	errCh := make(chan error)
	go func() {
		for evt := range tp.C {
			evtCh <- evt
		}
	}()
	go func() {
		<-ctx.Done()
		errCh <- ctx.Err()
	}()
	return evtCh, errCh
}

var _ event.Transport = (*Channel)(nil)
