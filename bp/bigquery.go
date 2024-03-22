package bq

import (
	"context"
	"os"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Option struct {
	projectId  string
	parameters []bigquery.QueryParameter
}

func (o *Option) WithProjectId(projectId string) {
	o.projectId = projectId
}

func (o *Option) WithQueryParameter(parameters []bigquery.QueryParameter) {
	o.parameters = parameters
}

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

	return &v, nil
}

func ListMessages[T any, constraint ConstraintProtoMessage[T]](ctx context.Context, query string, mtype T, option Option) (*Iterator[T, constraint], error) {
	projectId := option.projectId
	if projectId == "" {
		projectId = os.Getenv("GOOGLE_CLOUD_PROJECT")
	}
	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	q := client.Query(query)
	q.Parameters = option.parameters
	it, err := q.Read(ctx)
	if err != nil {
		return nil, err
	}

	return &Iterator[T, constraint]{
		it: it,
	}, nil
}
