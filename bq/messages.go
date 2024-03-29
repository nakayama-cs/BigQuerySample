package bq

import (
	"context"
	"os"

	"cloud.google.com/go/bigquery"
)

func ListMessages[T any, constraint ConstraintProtoMessage[T]](ctx context.Context, query string, option Option) (*Iterator[T, constraint], error) {
	projectId := option.projectId
	if projectId == "" {
		projectId = os.Getenv("BQ_GOOGLE_CLOUD_PROJECT")
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
	if option.pageToken != "" {
		it.PageInfo().Token = option.pageToken
	}
	if 0 < option.maxSize {
		it.PageInfo().MaxSize = option.maxSize
	}
	return &Iterator[T, constraint]{
		it:      it,
		maxSize: option.maxSize,
	}, nil
}
