package fib

import "context"

type App interface {
  Run(ctx context.Context) error
  Poll(ctx context.Context) (uint, error)
  Write(ctx context.Context, n uint)
}
