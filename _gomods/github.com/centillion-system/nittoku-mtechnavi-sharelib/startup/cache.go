package startup

import (
	"context"
	"fmt"
	"mtechnavi/sharelib/cache"
	"mtechnavi/sharelib/cloudruntime"
	"mtechnavi/sharelib/event"
	"mtechnavi/sharelib/event/transport"

	"github.com/sirupsen/logrus"
)

func initCache(ctx context.Context) (*cache.LRU, error) {
	var lru cache.LRU
	lru.Propagator = &event.Propagator{
		Transport: &transport.Firestore{
			ProjectID:      cloudruntime.ProjectID(ctx),
			CollectionPath: "__cache",
		},
	}
	lru.Logger = &cacheLogger{logrus.WithFields(nil)}
	if err := lru.Init(ctx); err != nil {
		return nil, fmt.Errorf("failed to create cache: %w", err)
	}
	go func() {
		ctx := context.Background()
		if err := lru.RunMaintenance(ctx); err != nil {
			panic(err)
		}
	}()
	return &lru, nil
}

type cacheLogger struct {
	*logrus.Entry
}

func (l *cacheLogger) WithError(err error) cache.Logger {
	return &cacheLogger{l.Entry.WithError(err)}
}

func (l *cacheLogger) Warnf(format string, args ...interface{}) {
	l.Entry.Warnf(format, args...)
}

var _ cache.Logger = (*cacheLogger)(nil)
