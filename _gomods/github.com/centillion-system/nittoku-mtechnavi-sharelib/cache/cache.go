package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"mtechnavi/sharelib/event"
	"net/http"
	"time"

	"github.com/tidwall/tinylru"
	"google.golang.org/api/option"
)

type LRU struct {
	Size int

	Propagator *event.Propagator

	Logger Logger

	store *tinylru.LRU
}

type Cache interface {
	Set(string, any)
	Get(string, any) bool
	Delete(string)

	NewHTTPClient(context.Context, ...option.ClientOption) *http.Client
}

var _ Cache = (*LRU)(nil)

func (c *LRU) Init(ctx context.Context) error {
	size := c.Size
	if size == 0 {
		size = 1024
	}
	c.store = &tinylru.LRU{}
	c.store.Resize(size)
	return nil
}

func (c *LRU) RunMaintenance(ctx context.Context) error {
	itr := c.Propagator.Iterate(ctx)
	for {
		var key string
		action, err := itr.Next(&key)
		if err != nil {
			return err
		}
		switch action {
		case "cache:delete":
			c.store.Delete(key)
		}
	}
}

func (c *LRU) Set(key string, v interface{}) {
	b, err := marshal(v)
	if err != nil {
		c.Logger.WithError(err).Warnf("failed to marshal binary: key=%q, v=%T", key, v)
		return
	}
	c.set(key, b)
}

func (c *LRU) Get(key string, v interface{}) bool {
	b, ok := c.get(key)
	if !ok {
		return false
	}
	if err := unmarshal(b, v); err != nil {
		c.Logger.WithError(err).Warnf("invalid cache: key=%q, v=%T", key, v)
		c.Delete(key)
		return false
	}
	return true
}

func (c *LRU) Delete(key string) {
	c.store.Delete(key)

	if c.Propagator == nil {
		return
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	if err := c.Propagator.Publish(ctx, "cache:delete", key); err != nil {
		c.Logger.WithError(err).Warnf("failed to propagate: key=%q", key)
	}
}

func (c *LRU) set(key string, value []byte) {
	c.store.Set(key, value)
}

func (c *LRU) get(key string) ([]byte, bool) {
	v, ok := c.store.Get(key)
	if !ok {
		return nil, false
	}
	b, ok := v.([]byte)
	return b, ok
}

func marshal(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func unmarshal(data []byte, v interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}
