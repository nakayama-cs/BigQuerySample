package bq

import (
	"context"
	"errors"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery"
)

var errCast = errors.New("invalid option")

func getOptionValue[T any](m map[string]interface{}, key string) (T, error) {
	if v, ok := m[key]; !ok || v == nil {
		var t T
		return t, fmt.Errorf("invalid option:%s", key)
	} else if cv, ok := v.(T); ok {
		return cv, nil
	}
	var t T
	return t, errCast
}

func ListMessages[T any, constraint ConstraintProtoMessage[T]](ctx context.Context, query string, options ...*Option) (*Iterator[T, constraint], error) {
	projectId, parameters, pageToken, maxSize, err := func() (string, []bigquery.QueryParameter, string, int, error) {
		mergedOption := map[string]interface{}{
			optProjectId:  os.Getenv("BQ_GOOGLE_CLOUD_PROJECT"),
			optParameters: []bigquery.QueryParameter{},
			optPageToken:  "",
			optMaxSize:    0,
		}
		for _, option := range options {
			if option != nil {
				for k, v := range option.values {
					mergedOption[k] = v
				}
			}
		}
		projectId, err := getOptionValue[string](mergedOption, optProjectId)
		if err != nil {
			return "", []bigquery.QueryParameter{}, "", 0, err
		}
		parameters, err := getOptionValue[[]bigquery.QueryParameter](mergedOption, optParameters)
		if err != nil {
			return "", []bigquery.QueryParameter{}, "", 0, err
		}
		pageToken, err := getOptionValue[string](mergedOption, optPageToken)
		if err != nil {
			return "", []bigquery.QueryParameter{}, "", 0, err
		}
		maxSize, err := getOptionValue[int](mergedOption, optMaxSize)
		if err != nil {
			return "", []bigquery.QueryParameter{}, "", 0, err
		}
		return projectId, parameters, pageToken, maxSize, nil
	}()
	if err != nil {
		return nil, err
	}

	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	q := client.Query(query)
	q.Parameters = parameters
	it, err := q.Read(ctx)
	if err != nil {
		return nil, err
	}
	if pageToken != "" {
		it.PageInfo().Token = pageToken
	}
	if 0 < maxSize {
		it.PageInfo().MaxSize = maxSize
	}
	return &Iterator[T, constraint]{
		it:      it,
		maxSize: maxSize,
	}, nil
}
