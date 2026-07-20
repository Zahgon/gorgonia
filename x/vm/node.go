package xvm

import (
	"context"

	"gorgonia.org/gorgonia"
)

type Doer interface {
	Do(...gorgonia.Value) (gorgonia.Value, error)
}

type node struct {
	id             int64
	op             Doer
	output         gorgonia.Value
	outputC        chan gorgonia.Value
	receivedValues int
	err            error
	inputValues    []gorgonia.Value
	inputC         chan ioValue
}

type ioValue struct {
	pos int
	v   gorgonia.Value
}

type stateFn func(context.Context, *node) stateFn

func defaultState(_ context.Context, n *node) stateFn {
	_ = "STUB: not implemented"
	return *new(stateFn)
}

func receiveInput(ctx context.Context, n *node) stateFn {
	_ = "STUB: not implemented"
	return *new(stateFn)
}

func computeFwd(_ context.Context, n *node) stateFn {
	_ = "STUB: not implemented"
	return *new(stateFn)
}

func emitOutput(ctx context.Context, n *node) stateFn {
	_ = "STUB: not implemented"
	return *new(stateFn)
}

func computeBackward(_ context.Context, _ *node) stateFn {
	_ = "STUB: not implemented"
	return *new(stateFn)
}

func (n *node) Compute(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func newOp(n *gorgonia.Node, hasOutputChan bool) *node { _ = "STUB: not implemented"; return nil }

func newInput(n *gorgonia.Node) *node { _ = "STUB: not implemented"; return nil }
