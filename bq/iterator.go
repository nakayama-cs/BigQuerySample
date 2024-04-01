package bq

import (
	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	"google.golang.org/api/iterator"
)

type Iterator[T any, constraint ConstraintProtoMessage[T]] struct {
	it      *bigquery.RowIterator
	maxSize int
	count   int
}

func (i *Iterator[T, constraint]) Next() (*T, error) {
	if i.maxSize != 0 && i.maxSize <= i.count {
		return nil, Done
	}
	var v T
	messageLoader := &protobq.MessageLoader{
		Message: constraint(&v),
	}
	err := i.it.Next(messageLoader)
	if err == iterator.Done {
		return nil, Done
	}
	if err != nil {
		return nil, err
	}
	i.count++
	return &v, nil
}

func (i *Iterator[T, confident]) PageInfo() *iterator.PageInfo {
	return i.it.PageInfo()
}
