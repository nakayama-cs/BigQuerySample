module mtechnavi/sharelib

go 1.19

require (
	cloud.google.com/go/compute/metadata v0.2.3
	cloud.google.com/go/firestore v1.9.0
	cloud.google.com/go/storage v1.29.0
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v1.12.0
	github.com/envoyproxy/protoc-gen-validate v0.9.1
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/chi/v5 v5.0.8
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/nicksnyder/go-i18n/v2 v2.2.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/sirupsen/logrus v1.9.0
	github.com/stoewer/go-strcase v1.2.1
	github.com/stretchr/testify v1.8.2
	github.com/tidwall/tinylru v1.1.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.40.0
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	golang.org/x/mod v0.7.0
	golang.org/x/net v0.7.0
	golang.org/x/text v0.7.0
	google.golang.org/api v0.111.0
	google.golang.org/genproto v0.0.0-20230227214838-9b19f0bdc514
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	mtechnavi/dataproxy v0.0.0
	mtechnavi/idp v0.0.0
	mtechnavi/programoption v0.0.0
)

require (
	cloud.google.com/go v0.110.0 // indirect
	cloud.google.com/go/compute v1.18.0 // indirect
	cloud.google.com/go/iam v0.12.0 // indirect
	cloud.google.com/go/longrunning v0.4.1 // indirect
	cloud.google.com/go/trace v1.8.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.36.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel/metric v0.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/oauth2 v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	mtechnavi/company v0.0.0 => ./_gomods/github.com/centillion-system/nittoku-mtechnavi-clientsdk-go/company/
	mtechnavi/dataproxy v0.0.0 => ./_gomods/github.com/centillion-system/nittoku-mtechnavi-clientsdk-go/dataproxy/
	mtechnavi/idp v0.0.0 => ./_gomods/github.com/centillion-system/nittoku-mtechnavi-clientsdk-go/idp/
	mtechnavi/programoption v0.0.0 => ./_gomods/github.com/centillion-system/nittoku-mtechnavi-clientsdk-go/programoption/
)
