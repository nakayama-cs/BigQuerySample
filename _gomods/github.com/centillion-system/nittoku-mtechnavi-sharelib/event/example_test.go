package event_test

import (
	"context"
	"errors"
	"fmt"
	"mtechnavi/sharelib/event"
	"mtechnavi/sharelib/event/transport"
	"time"
)

type Item struct {
	Message string
}

func (v *Item) String() string {
	return "[" + v.Message + "]"
}

func Example() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tp := transport.NewChannel()
	sender := &event.Propagator{
		Transport: tp,
	}
	receiver := &event.Propagator{
		Transport: &transport.Channel{
			C: tp.C,
		},
	}

	go func() {
		events := []struct {
			Name  string
			Value interface{}
		}{
			{"a", &Item{
				Message: "first message",
			}},
			{"b", &Item{
				Message: "second message",
			}},
		}
		for _, evt := range events {
			fmt.Printf("SEND: %v\n", evt.Name)
			if err := sender.Publish(ctx, evt.Name, evt.Value); err != nil {
				panic(err)
			}
		}
		time.Sleep(time.Second)
		cancel()
	}()

	itr := receiver.Iterate(ctx)
loop:
	for {
		var value *Item
		name, err := itr.Next(&value)
		switch {
		case err == nil:
			// OK
		case errors.Is(err, context.Canceled):
			// Stopped
			break loop
		default:
			panic(err)
		}
		fmt.Printf("RECV: %v = %v\n", name, value)
	}
	// Output:
	// SEND: a
	// SEND: b
	// RECV: a = [first message]
	// RECV: b = [second message]
}
