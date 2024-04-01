package bq

import "cloud.google.com/go/bigquery"

type Option struct {
	values map[string]interface{}
}

const (
	OptProjectId  = "projectId"
	OptParameters = "parameters"
	OptPageToken  = "pageToken"
	OptMaxSize    = "maxSize"
)

func WithProjectId(projectId string) *Option {
	return &Option{
		values: map[string]interface{}{
			OptProjectId: projectId,
		},
	}
}

func WithQueryParameter(parameters []bigquery.QueryParameter) *Option {
	return &Option{
		values: map[string]interface{}{
			OptParameters: parameters,
		},
	}
}

func WitPageToken(pageToken string) *Option {
	return &Option{
		values: map[string]interface{}{
			OptPageToken: pageToken,
		},
	}
}

func WitMaxSize(maxSize int) *Option {
	return &Option{
		values: map[string]interface{}{
			OptMaxSize: maxSize,
		},
	}
}
