package bq

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery"
)

func ListMessages[T any, constraint ConstraintProtoMessage[T]](ctx context.Context, query string, options ...*Option) (*Iterator[T, constraint], error) {
	mergedOption, err := mergeOption(options...)
	if err != nil {
		return nil, err
	}

	client, err := bigquery.NewClient(ctx, mergedOption.ProjectId)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	q := client.Query(query)
	q.Parameters = mergedOption.Parameters
	it, err := q.Read(ctx)
	if err != nil {
		return nil, err
	}
	if mergedOption.PageToken != "" {
		it.PageInfo().Token = mergedOption.PageToken
	}
	if 0 < mergedOption.MaxSize {
		it.PageInfo().MaxSize = mergedOption.MaxSize
	}
	return &Iterator[T, constraint]{
		it:      it,
		maxSize: mergedOption.MaxSize,
	}, nil
}

func mergeOption(options ...*Option) (*struct {
	ProjectId  string
	Parameters []bigquery.QueryParameter
	PageToken  string
	MaxSize    int
}, error) {

	mergedOption := map[string]interface{}{
		OptProjectId:  os.Getenv("BQ_GOOGLE_CLOUD_PROJECT"),
		OptParameters: []bigquery.QueryParameter{},
		OptPageToken:  "",
		OptMaxSize:    0,
	}
	for _, option := range options {
		if option != nil {
			for k, v := range option.values {
				mergedOption[k] = v
			}
		}
	}
	projectId, err := getOptionValue[string](mergedOption, OptProjectId)
	if err != nil {
		return nil, err
	}
	parameters, err := getOptionValue[[]bigquery.QueryParameter](mergedOption, OptParameters)
	if err != nil {
		return nil, err
	}
	pageToken, err := getOptionValue[string](mergedOption, OptPageToken)
	if err != nil {
		return nil, err
	}
	maxSize, err := getOptionValue[int](mergedOption, OptMaxSize)
	if err != nil {
		return nil, err
	}
	return &struct {
		ProjectId  string
		Parameters []bigquery.QueryParameter
		PageToken  string
		MaxSize    int
	}{
		ProjectId:  projectId,
		Parameters: parameters,
		PageToken:  pageToken,
		MaxSize:    maxSize,
	}, nil
}

func getOptionValue[T any](m map[string]interface{}, key string) (T, error) {
	if v, ok := m[key]; !ok || v == nil {
		var t T
		return t, fmt.Errorf("invalid option:%s", key)
	} else if cv, ok := v.(T); ok {
		return cv, nil
	}
	var t T
	return t, ErrCast
}
