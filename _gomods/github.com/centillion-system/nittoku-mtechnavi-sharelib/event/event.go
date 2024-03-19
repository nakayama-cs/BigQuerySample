package event

import (
	"context"
	"errors"
	"fmt"
	"mtechnavi/sharelib/event/codec"
	pb "mtechnavi/sharelib/event/protobuf"
	"sync"
)

type Propagator struct {
	Transport Transport

	Codec Codec

	once sync.Once

	initErr error
}

func (p *Propagator) init(ctx context.Context) error {
	p.once.Do(func() {
		if p.Transport == nil {
			p.initErr = errors.New("Transport field is mandatory")
			return
		}
		if p.Codec == nil {
			p.Codec = &codec.Gob{}
		}

		if tp, ok := p.Transport.(interface {
			Init(context.Context) error
		}); ok {
			if err := tp.Init(ctx); err != nil {
				p.initErr = fmt.Errorf("failed to initialize transport: %w", err)
				return
			}
		}
	})
	return p.initErr
}

func (p *Propagator) Publish(ctx context.Context, name string, value interface{}) error {
	if err := p.init(ctx); err != nil {
		return err
	}

	var evt pb.Event
	evt.Name = name
	if data, err := p.Codec.Encode(value); err != nil {
		return err
	} else {
		evt.Data = data
	}
	return p.Transport.Publish(ctx, &evt)
}

func (p *Propagator) Iterate(ctx context.Context) Iterator {
	if err := p.init(ctx); err != nil {
		return &errIterator{
			err: err,
		}
	}

	evtCh, errCh := p.Transport.Subscribe(ctx)
	return &iterator{
		codec: p.Codec,
		evtCh: evtCh,
		errCh: errCh,
	}
}

type Iterator interface {
	Next(interface{}) (string, error)
}

type errIterator struct {
	err error
}

func (itr *errIterator) Next(interface{}) (string, error) {
	return "", itr.err
}

type iterator struct {
	codec Codec
	evtCh <-chan *pb.Event
	errCh <-chan error
}

func (itr *iterator) Next(value interface{}) (string, error) {
	var evt *pb.Event
	select {
	case evt = <-itr.evtCh:
	case err := <-itr.errCh:
		return "", err
	}
	if err := itr.codec.Decode(evt.Data, value); err != nil {
		return "", err
	} else {
		return evt.Name, nil
	}
}
