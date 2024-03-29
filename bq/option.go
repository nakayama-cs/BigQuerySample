package bq

import "cloud.google.com/go/bigquery"

type Option struct {
	values map[string]interface{}
}

const (
	optProjectId  = "projectId"
	optParameters = "parameters"
	optPageToken  = "pageToken"
	optMaxSize    = "maxSize"
)

func WithProjectId(projectId string) *Option {
	return &Option{
		values: map[string]interface{}{
			optProjectId: projectId,
		},
	}
}

func WithQueryParameter(parameters []bigquery.QueryParameter) *Option {
	return &Option{
		values: map[string]interface{}{
			optParameters: parameters,
		},
	}
}

func WitPageToken(pageToken string) *Option {
	return &Option{
		values: map[string]interface{}{
			optPageToken: pageToken,
		},
	}
}

func WitMaxSize(maxSize int) *Option {
	return &Option{
		values: map[string]interface{}{
			optMaxSize: maxSize,
		},
	}
}
