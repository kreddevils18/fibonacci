package main

import (
  "context"
  "github.com/kreddevils18/fibonacci/internal/fib/usecase"
  jaeger "github.com/kreddevils18/fibonacci/pkg/jaeger"
  "go.opentelemetry.io/otel"
  "log"
  "os"
  "os/signal"
  "time"
)

func main() {
	l := log.New(os.Stdout, "", 0)

	res := jaeger.NewResource()
	tp, err := jaeger.NewJaegerProvider("http://localhost:14268/api/traces", res)

	if err != nil {
		l.Fatal(err)
	}
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(ctx context.Context) {
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		if err := tp.Shutdown(ctx); err != nil {
			l.Fatal(err)
		}
	}(ctx)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	app := usecase.NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(ctx)
	}()

	select {
	case <- sigCh:
		l.Println("\nGoodbye")
		return
	case err := <- errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}

