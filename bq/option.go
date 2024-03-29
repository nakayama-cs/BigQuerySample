package bq

import "cloud.google.com/go/bigquery"

type Option struct {
	projectId  string
	parameters []bigquery.QueryParameter
	pageToken  string
	maxSize    int
}

func (o *Option) WithProjectId(projectId string) *Option {
	o.projectId = projectId
	return o
}

func (o *Option) WithQueryParameter(parameters []bigquery.QueryParameter) *Option {
	o.parameters = parameters
	return o
}

func (o *Option) WitPageToken(pageToken string) *Option {
	o.pageToken = pageToken
	return o
}

func (o *Option) WitMaxSize(maxSize int) *Option {
	o.maxSize = maxSize
	return o
}
