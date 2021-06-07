module github.com/Ch1f/otel/exporters/trace/zipkin

go 1.13

replace github.com/Ch1f/otel => ../../..

require (
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/stretchr/testify v1.6.1
	github.com/Ch1f/otel v0.7.0
	google.golang.org/grpc v1.30.0
)
