module github.com/Ch1f/otel/example/jaeger

go 1.13

replace (
	github.com/Ch1f/otel => ../..
	github.com/Ch1f/otel/exporters/trace/jaeger => ../../exporters/trace/jaeger
)

require (
	github.com/Ch1f/otel v0.7.0
	github.com/Ch1f/otel/exporters/trace/jaeger v0.7.0
)
