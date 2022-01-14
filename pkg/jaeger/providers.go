package jaeger

import (
  "go.opentelemetry.io/otel/sdk/resource"
  "go.opentelemetry.io/otel/sdk/trace"
)

func NewJaegerProvider(url string, resource *resource.Resource) (*trace.TracerProvider, error){
  jaegerExporter, err := NewJaegerExporter(url)

  if err != nil {
    return nil, err
  }

  tp := trace.NewTracerProvider(
    trace.WithBatcher(jaegerExporter),
    trace.WithResource(resource),
  )

  return tp, nil
}

