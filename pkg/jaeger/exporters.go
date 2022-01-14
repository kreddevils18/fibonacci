package jaeger

import "go.opentelemetry.io/otel/exporters/jaeger"

func NewJaegerExporter(url string) (*jaeger.Exporter, error) {
  return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}
