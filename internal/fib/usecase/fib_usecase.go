package usecase

import (
  "context"
  "fmt"
  "github.com/kreddevils18/fibonacci/internal/fib"
  "go.opentelemetry.io/otel"
  "go.opentelemetry.io/otel/attribute"
  "go.opentelemetry.io/otel/trace"
  "io"
  "log"
  "strconv"
)

const name = "fib"

type App struct {
  r io.Reader
  l *log.Logger
}

func NewApp(r io.Reader, l *log.Logger) *App {
  return &App{
    r: r,
    l: l,
  }
}

func (a *App) Run(ctx context.Context) error {
  for {
    var span trace.Span
    ctx, span = otel.Tracer(name).Start(ctx, "Run")

    n, err := a.Poll(ctx)
    if err != nil {
      span.End()
      return err
    }
    a.Write(ctx, n)
    span.End()
  }
}

func (a *App) Poll(ctx context.Context) (uint, error) {
  _, span := otel.Tracer(name).Start(ctx, "Poll")
  defer span.End()

  a.l.Print("What Fibonacci number would you like to know: ")

  var n uint
  _, err := fmt.Fscanf(a.r, "%d\n", &n)

  nStr := strconv.FormatUint(uint64(n), 10)
  span.SetAttributes(attribute.String("request.n", nStr))

  return n, err
}

func (a *App) Write(ctx context.Context, n uint) {
  var span trace.Span
  ctx, span = otel.Tracer(name).Start(ctx, "Write")
  defer span.End()

  f, err := func (ctx context.Context) (uint64, error) {
    _, span := otel.Tracer(name).Start(ctx, "Fibonacci")
    defer span.End()
    return fib.Fibonacci(n)
  }(ctx)

  if err != nil {
    a.l.Printf("Fibonacci(%d): %v\n", n, err)
  } else {
    a.l.Printf("Fibonacci(%d): %d\n", n, f)
  }
}
