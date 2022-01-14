package jaeger

import (
  "go.opentelemetry.io/otel/attribute"
  "go.opentelemetry.io/otel/sdk/resource"
  "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// newResource returns a resource describing this application.
func NewResource() *resource.Resource {
  r, _ := resource.Merge(
    resource.Default(),
    resource.NewWithAttributes(
      semconv.SchemaURL,
      semconv.ServiceNameKey.String("fib"),
      semconv.ServiceVersionKey.String("v0.1.0"),
      attribute.String("environment", "demo"),
    ),
  )
  return r
}
