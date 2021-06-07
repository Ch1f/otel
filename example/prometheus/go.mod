module github.com/Ch1f/otel/example/prometheus

go 1.13

replace (
	github.com/Ch1f/otel => ../..
	github.com/Ch1f/otel/exporters/metric/prometheus => ../../exporters/metric/prometheus
)

require (
	github.com/Ch1f/otel v0.7.0
	github.com/Ch1f/otel/exporters/metric/prometheus v0.7.0
)
