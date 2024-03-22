package bq

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Option struct{}

func (Option) WithQueryParameter([]*bigquery.QueryParameter) {}
func (Option) WithLimit(limit int)                           {}
func (Option) WithPageToken(token string)                    {}

type ConstraintProtoMessage[X any] interface {
	protoreflect.ProtoMessage
	*X
}

type Iterator[T any, constraint ConstraintProtoMessage[T]] struct {
	it *bigquery.RowIterator
}

func (i Iterator[T, constraint]) Next() (*T, error) {
	var v T
	messageLoader := &protobq.MessageLoader{
		Message: constraint(&v),
	}

	err := i.it.Next(messageLoader)
	if err != nil {
		return nil, err
	}

	message := proto.Clone(messageLoader.Message)
	if testMessage, ok := message.(T); ok {
		log.Println(testMessage)
	}

	return &v, nil
}

func ListMessages[T any, constraint ConstraintProtoMessage[T]](ctx context.Context, query string, mtype T, options ...Option) (*Iterator[T, constraint], error) {
	// 環境変数からGCPのプロジェクトIDを取得する
	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT")

	// BigQueryのクライアントを取得
	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// クエリーの作成
	q := client.Query("call nkym_test.pr_sample()")

	// クエリーの実行
	it, err := q.Read(ctx)
	if err != nil {
		panic(err)
	}

	return &Iterator[T, constraint]{
		it: it,
	}, nil
}
