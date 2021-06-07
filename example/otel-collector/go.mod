module github.com/Ch1f/otel/example/otel-collector

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector v0.3.0
	github.com/Ch1f/otel v0.7.0
	github.com/Ch1f/otel/exporters/otlp v0.7.0
	google.golang.org/grpc v1.30.0
)

replace (
	github.com/Ch1f/otel => ../..
	github.com/Ch1f/otel/exporters/otlp => ../../exporters/otlp
)
