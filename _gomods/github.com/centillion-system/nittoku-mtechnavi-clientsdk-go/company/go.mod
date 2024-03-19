module mtechnavi/company

go 1.19

require (
	github.com/envoyproxy/protoc-gen-validate v0.9.1
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	mtechnavi/sharelib v0.0.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230227214838-9b19f0bdc514 // indirect
)

replace mtechnavi/sharelib v0.0.0 => ../_gomods/github.com/centillion-system/nittoku-mtechnavi-sharelib/
