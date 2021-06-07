module github.com/Ch1f/otel/exporters/trace/jaeger

go 1.13

replace github.com/Ch1f/otel => ../../..

require (
	github.com/apache/thrift v0.13.0
	github.com/google/go-cmp v0.5.0
	github.com/stretchr/testify v1.6.1
	github.com/Ch1f/otel v0.7.0
	google.golang.org/api v0.29.0
	google.golang.org/grpc v1.30.0
)
