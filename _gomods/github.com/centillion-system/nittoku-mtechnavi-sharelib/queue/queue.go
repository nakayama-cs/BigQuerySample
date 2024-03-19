package queue

import (
	"context"
	"errors"
	"io"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	TimeNow           = time.Now
	ErrNoMoreMessages = errors.New("queue has no more messages")
)

type Order struct {
	BucketName    string
	once          sync.Once
	initErr       error
	storageClient *storage.Client
}

func (o *Order) init(ctx context.Context) error {
	o.once.Do(func() {
		if c, err := storage.NewClient(ctx); err != nil {
			o.initErr = err
			return
		} else {
			o.storageClient = c
		}
		if o.BucketName == "" {
			o.initErr = errors.New("bucketName is mandatory")
			return
		}
	})
	return o.initErr
}

func (o *Order) Push(ctx context.Context, objectPrefix string, data proto.Message) error {
	logger := ctxlogrus.Extract(ctx)
	if err := o.init(ctx); err != nil {
		return err
	}

	objectName, err := o.push(ctx, objectPrefix, data)
	if err != nil {
		logger.WithError(err).Errorf("failed to upload queue object")
		return err
	}
	logger.Debugf("data object has been written: gs://%s/%s", o.BucketName, objectName)

	return nil
}

func (o *Order) Pull(ctx context.Context, objectPrefix string) (*Message, error) {
	logger := ctxlogrus.Extract(ctx)
	if err := o.init(ctx); err != nil {
		return nil, err
	}

	attrs, err := o.pull(ctx, objectPrefix)
	if errors.Is(err, ErrNoMoreMessages) {
		return nil, ErrNoMoreMessages
	} else if err != nil {
		logger.WithError(err).Errorf("failed to pull data by prefix: gs://%s/%s", o.BucketName, objectPrefix)
		return nil, err
	}

	// ロックを取得
	logger.Debugf("trying to lock pulled object: gs://%s/%s", attrs.Bucket, attrs.Name)
	object := o.storageClient.Bucket(attrs.Bucket).Object(attrs.Name)
	mu := &locker{
		LockId: uuid.NewString(),
		Object: object,
		Attrs:  attrs,
	}
	switch err := mu.Lock(ctx); {
	case err == nil:
		logger.Debugf("locked pulled object: gs://%s/%s", attrs.Bucket, attrs.Name)
	case errors.Is(err, ErrLockedByOther):
		return nil, err
	default:
		logger.WithError(err).Errorf("failed to lock pulled object: gs://%s/%s", attrs.Bucket, attrs.Name)
		return nil, err
	}

	data, err := func(rc *storage.Reader, err error) ([]byte, error) {
		if err != nil {
			return nil, err
		}
		defer rc.Close()
		return io.ReadAll(rc)
	}(object.NewReader(ctx))
	if err != nil {
		return nil, err
	}

	var val anypb.Any
	if err := proto.Unmarshal(data, &val); err != nil {
		return nil, err
	}
	out := &Message{
		logger: logger,
		mu:     mu,
	}
	if v, err := val.UnmarshalNew(); err != nil {
		return nil, err
	} else {
		out.Data = v
	}
	return out, nil
}

func (o *Order) push(ctx context.Context, objectPrefix string, data proto.Message) (string, error) {
	var val anypb.Any
	if err := val.MarshalFrom(data); err != nil {
		return "", err
	}
	b, err := proto.Marshal(&val)
	if err != nil {
		return "", err
	}

	// 同一時刻にオブジェクト発行しても被らないよう、uuidをサフィックスに付与
	objectName := strconv.FormatInt(TimeNow().UnixMicro(), 10) + "." + uuid.NewString()
	object := o.storageClient.Bucket(o.BucketName).Object(path.Join(objectPrefix, objectName))
	object = object.If(storage.Conditions{
		DoesNotExist: true,
	})

	writer := object.NewWriter(ctx)
	if _, err := writer.Write(b); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}
	return object.ObjectName(), nil
}

func (o *Order) pull(ctx context.Context, objectPrefix string) (*storage.ObjectAttrs, error) {
	logger := ctxlogrus.Extract(ctx)

	if !strings.HasSuffix(objectPrefix, "/") {
		objectPrefix += "/"
	}
	logger.Debugf("pull gs://%s/%s", o.BucketName, objectPrefix)

	bucket := o.storageClient.Bucket(o.BucketName)
	itr := bucket.Objects(ctx, &storage.Query{
		Prefix: objectPrefix,
	})
	// オブジェクト名（タイムスタンプ）順にイテレートされるため、
	// 最初にpushされたメッセージを取得するためには初回のオブジェクトのみ取得して終了させる
	// 削除遅延がある可能性を考慮し、論理削除状態なら無視する
	for {
		attrs, err := itr.Next()
		if errors.Is(err, iterator.Done) {
			return nil, ErrNoMoreMessages
		} else if err != nil {
			return nil, err
		}

		if deletedBy, ok := attrs.Metadata["deletedBy"]; ok && deletedBy != "" {
			logger.Warnf("detected deletion delayed (deletedBy=%q), skip: gs://%s/%s", deletedBy, attrs.Bucket, attrs.Name)
			continue
		}
		return attrs, nil
	}
}
