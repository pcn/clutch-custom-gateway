module github.com/pcn/clutch-custom-gateway/backend

go 1.13

require (
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/golang/protobuf v1.3.5
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/lyft/clutch/backend v0.0.0-20200729214654-90db01f0aa54
	github.com/lyft/clutch/tools v0.0.0-20200729214654-90db01f0aa54 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/uber-go/tally v3.3.17+incompatible
	go.uber.org/zap v1.15.0
	google.golang.org/genproto v0.0.0-20200513103714-09dca8ec2884
	google.golang.org/grpc v1.30.0
)
