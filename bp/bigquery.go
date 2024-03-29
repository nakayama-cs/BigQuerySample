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

type Iterator struct {
	it *bigquery.RowIterator
}

func (i *Iterator) Next(m protoreflect.ProtoMessage) error {
	messageLoader := &protobq.MessageLoader{
		Message: m,
	}

	err := i.it.Next(messageLoader)
	if err != nil {
		return err
	}

	return nil
}

func ListMessages(ctx context.Context, query string, option Option) (*Iterator, error) {
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

	return &Iterator{
		it: it,
	}, nil
}
