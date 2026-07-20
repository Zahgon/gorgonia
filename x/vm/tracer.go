package xvm

import (
	"context"
	"time"
)

type Trace struct {
	StateFunction string
	ID            int64
	Start         time.Time
	End           time.Time `json:",omitempty"`
}

type chanTracerContextKey int

const (
	globalTracerContextKey chanTracerContextKey = 0
)

func WithTracing(parent context.Context) (context.Context, <-chan Trace) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func CloseTracing(ctx context.Context) { _ = "STUB: not implemented"; return }

func extractTracingChannel(ctx context.Context) chan<- Trace { _ = "STUB: not implemented"; return nil }

var now = time.Now

func trace(ctx context.Context, t *Trace, n *node, state stateFn) *Trace {
	_ = "STUB: not implemented"
	return nil
}
