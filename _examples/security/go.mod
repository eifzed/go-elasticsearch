module github.com/eifzed/go-elasticsearch/v8/_examples/security

go 1.21

toolchain go1.21.8

replace github.com/eifzed/go-elasticsearch/v8 => ../..

require github.com/eifzed/go-elasticsearch/v8 v8.0.0-00010101000000-000000000000

require (
	github.com/elastic/elastic-transport-go/v8 v8.4.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
)
