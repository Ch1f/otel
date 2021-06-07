module go.opentelemetry.go/otel/example/zipkin

go 1.13

replace (
	github.com/Ch1f/otel => ../..
	github.com/Ch1f/otel/exporters/trace/zipkin => ../../exporters/trace/zipkin
)

require (
	github.com/Ch1f/otel v0.7.0
	github.com/Ch1f/otel/exporters/trace/zipkin v0.7.0
)
