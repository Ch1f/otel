module github.com/Ch1f/otel/example/grpc

go 1.13

replace github.com/Ch1f/otel => ../..

require (
	github.com/golang/protobuf v1.4.2
	github.com/Ch1f/otel v0.7.0
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980
	google.golang.org/grpc v1.30.0
)
