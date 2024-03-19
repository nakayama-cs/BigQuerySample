package queue

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/googleapi"
	"google.golang.org/protobuf/proto"
)

const (
	initialDuration = 10 * time.Second

	extendDuration = 10 * time.Second
)

type ctxKey string

var (
	ErrNotLocked = errors.New("object is not locked")

	ErrLockedByOther = errors.New("object is locked by other")

	ctxKeyNow = ctxKey("time.Now()")
)

type Message struct {
	Data proto.Message

	logger *logrus.Entry

	mu *locker
}

func (v *Message) Ack() {
	v.logger.Tracef("ack gs://%s/%s", v.mu.Attrs.Bucket, v.mu.Attrs.Name)

	ctx := context.Background()
	ctx = ctxlogrus.ToContext(ctx, v.logger)
	if err := v.mu.Delete(ctx); err != nil {
		v.logger.WithError(err).Warnf("failed to ack gs://%s/%s", v.mu.Attrs.Bucket, v.mu.Attrs.Name)
	}
}

func (v *Message) Nack() {
	v.logger.Tracef("nack gs://%s/%s", v.mu.Attrs.Bucket, v.mu.Attrs.Name)

	ctx := context.Background()
	ctx = ctxlogrus.ToContext(ctx, v.logger)
	switch err := v.mu.Unlock(ctx); {
	case err == nil:
	case errors.Is(err, storage.ErrObjectNotExist):
		v.logger.WithError(err).Tracef("failed to nack gs://%s/%s", v.mu.Attrs.Bucket, v.mu.Attrs.Name)
	default:
		v.logger.WithError(err).Warnf("failed to nack gs://%s/%s", v.mu.Attrs.Bucket, v.mu.Attrs.Name)
	}
}

type locker struct {
	LockId string

	Object *storage.ObjectHandle

	Attrs *storage.ObjectAttrs
}

func (l *locker) Lock(ctx context.Context) error {
	now := nowFromContext(ctx)
	ctx = withNow(ctx, now)

	switch state := lockState(l.LockId, l.Attrs, now); state {
	case lockStateNone:
		return l.extend(ctx, initialDuration)
	case lockStateLockedByMe:
		return l.extend(ctx, extendDuration)
	case lockStateLockedByOther:
		return ErrLockedByOther
	default:
		panic(fmt.Sprintf("invalid lockState: %v", state))
	}
}

func (l *locker) Extend(ctx context.Context) error {
	now := nowFromContext(ctx)
	ctx = withNow(ctx, now)

	switch state := lockState(l.LockId, l.Attrs, now); state {
	case lockStateNone:
		return ErrNotLocked
	case lockStateLockedByMe:
		return l.extend(ctx, extendDuration)
	case lockStateLockedByOther:
		return ErrLockedByOther
	default:
		panic(fmt.Sprintf("invalid lockState: %v", state))
	}
}

func (l *locker) extend(ctx context.Context, d time.Duration) error {
	logger := ctxlogrus.Extract(ctx)
	now := nowFromContext(ctx)

	meta := l.Attrs.Metadata
	if meta == nil {
		meta = map[string]string{}
	}
	meta["lockId"] = l.LockId
	meta["expiry"] = now.Add(d).Format(time.RFC3339)

	object := l.Object.If(storage.Conditions{
		GenerationMatch:     l.Attrs.Generation,
		MetagenerationMatch: l.Attrs.Metageneration,
	})
	attrs, err := object.Update(ctx, storage.ObjectAttrsToUpdate{
		Metadata: meta,
	})
	var apiErr *googleapi.Error
	if errors.As(err, &apiErr) {
		switch apiErr.Code {
		case http.StatusPreconditionFailed:
			return ErrLockedByOther
		}
	}
	if errors.Is(err, storage.ErrObjectNotExist) {
		logger.WithError(err).Infof("fail to extend given object: %v", l.Attrs.Name)
		return err
	} else if err != nil {
		logger.WithError(err).Errorf("fail to extend given object: %v", l.Attrs.Name)
		return err
	}
	l.Attrs = attrs

	// 書き込み成功したら、再度チェック
	switch state := lockState(l.LockId, l.Attrs, now); state {
	case lockStateLockedByMe:
		return nil
	default:
		return ErrLockedByOther
	}
}

func (l *locker) Unlock(ctx context.Context) error {
	now := nowFromContext(ctx)

	switch state := lockState(l.LockId, l.Attrs, now); state {
	case lockStateNone:
		return nil
	case lockStateLockedByMe:
		// go next
	case lockStateLockedByOther:
		return ErrLockedByOther
	default:
		panic(fmt.Sprintf("invalid lockState: %v", state))
	}

	object := l.Object.If(storage.Conditions{
		GenerationMatch:     l.Attrs.Generation,
		MetagenerationMatch: l.Attrs.Metageneration,
	})
	attrs, err := object.Update(ctx, storage.ObjectAttrsToUpdate{
		Metadata: map[string]string{},
	})
	if err != nil {
		return err
	}
	l.Attrs = attrs
	return nil
}

func (l *locker) Delete(ctx context.Context) error {
	now := nowFromContext(ctx)

	attrs, err := l.Object.Attrs(ctx)
	if err != nil {
		return err
	}
	object := l.Object.If(storage.Conditions{
		GenerationMatch:     attrs.Generation,
		MetagenerationMatch: attrs.Metageneration,
	})
	switch state := lockState(l.LockId, attrs, now); state {
	case lockStateNone:
		return ErrNotLocked
	case lockStateLockedByMe:
		// do nothing
	case lockStateLockedByOther:
		return ErrLockedByOther
	default:
		panic(fmt.Sprintf("invalid lockState: %v", state))
	}
	// 削除が数秒遅延する可能性があるため、メタデータ更新して論理削除も合わせて実施
	removed := attrs.Metadata
	removed["deletedBy"] = l.LockId
	attrs, err = object.Update(ctx, storage.ObjectAttrsToUpdate{
		Metadata: removed,
	})
	if err != nil {
		return err
	}
	return object.If(storage.Conditions{
		GenerationMatch:     attrs.Generation,
		MetagenerationMatch: attrs.Metageneration,
	}).Delete(ctx)
}

const (
	lockStateNone = iota
	lockStateLockedByMe
	lockStateLockedByOther
)

func lockState(lockId string, attrs *storage.ObjectAttrs, now time.Time) int {
	// 有効期限切れチェック
	lockedBy := attrs.Metadata["lockId"]
	var expiresAt time.Time
	if s := attrs.Metadata["expiry"]; s != "" {
		if t, err := time.Parse(time.RFC3339, s); err != nil {
			logrus.WithError(err).Warnf("invalid storage object metadata: expiry=%q", s)
		} else {
			expiresAt = t
		}
	}

	// 未ロック
	if lockedBy == "" || expiresAt.IsZero() {
		return lockStateNone
	}

	if now.Before(expiresAt) {
		// 有効期限内
		if lockedBy != lockId {
			// 他者所有
			return lockStateLockedByOther
		}
	} else {
		// 有効期限切れ
		logrus.Debugf("lock expired! lockId:%s lockedBy:%s expiresAt:%v now:%v", lockId, lockedBy, expiresAt, now)
	}
	// ロック所有
	return lockStateLockedByMe
}

func withNow(parent context.Context, v time.Time) context.Context {
	return context.WithValue(parent, ctxKeyNow, v)
}

func nowFromContext(ctx context.Context) time.Time {
	if v, ok := ctx.Value(ctxKeyNow).(time.Time); ok {
		return v
	} else {
		return TimeNow()
	}
}
