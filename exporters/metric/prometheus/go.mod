module github.com/Ch1f/otel/exporters/metric/prometheus

go 1.13

replace github.com/Ch1f/otel => ../../..

require (
	github.com/prometheus/client_golang v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/Ch1f/otel v0.7.0
)
