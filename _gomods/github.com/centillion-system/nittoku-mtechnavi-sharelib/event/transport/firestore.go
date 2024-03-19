package transport

import (
	"context"
	"errors"
	"mtechnavi/sharelib/event"
	pb "mtechnavi/sharelib/event/protobuf"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

type Firestore struct {
	ProjectID string

	CollectionPath string

	cli *firestore.Client
}

func (tp *Firestore) Init(ctx context.Context) error {
	if cli, err := firestore.NewClient(ctx, tp.ProjectID); err != nil {
		return err
	} else {
		tp.cli = cli
	}
	return nil
}

func (tp *Firestore) Publish(ctx context.Context, evt *pb.Event) error {
	timestamp := time.Now()
	col := tp.cli.Collection(tp.CollectionPath)
	doc := col.Doc(uuid.NewString())
	_, err := doc.Set(ctx, map[string]interface{}{
		"timestamp": timestamp.UnixMicro(),
		"name":      evt.Name,
		"data":      evt.Data,
	})
	return err
}

func (tp *Firestore) Subscribe(ctx context.Context) (<-chan *pb.Event, <-chan error) {
	timestamp := time.Now()
	evtCh := make(chan *pb.Event, 1024)
	errCh := make(chan error, 1)
	go func() {
		col := tp.cli.Collection(tp.CollectionPath)
		itr := col.Query.
			// indexを利用しないため、1つの条件のみ
			Where("timestamp", ">=", timestamp.UnixMicro()).
			Snapshots(ctx)
		defer itr.Stop()
		for {
			qs, err := itr.Next()
			if errors.Is(err, iterator.Done) {
				break
			} else if err != nil {
				errCh <- err
				continue
			}
			for _, change := range qs.Changes {
				if change.Kind != firestore.DocumentAdded {
					continue
				}
				data := change.Doc.Data()
				var evt pb.Event
				if v, ok := data["name"].(string); ok {
					evt.Name = v
				}
				if v, ok := data["data"].([]byte); ok {
					evt.Data = v
				}
				evtCh <- &evt
			}
		}
	}()
	return evtCh, errCh
}

var _ event.Transport = (*Firestore)(nil)
