package xvm

import (
	"context"
	"time"

	"gorgonia.org/gorgonia"
)

type Machine struct {
	nodes  []*node
	pubsub *pubsub
}

func NewMachine(g *gorgonia.ExprGraph) *Machine { _ = "STUB: not implemented"; return nil }

func createNetwork(ns []*node, g *gorgonia.ExprGraph) *pubsub {
	_ = "STUB: not implemented"
	return nil
}

func (m *Machine) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *Machine) Close() { _ = "STUB: not implemented"; return }

type nodeError struct {
	id  int64
	t   time.Time
	err error
}

type nodeErrors []nodeError

func (e nodeErrors) Error() string { _ = "STUB: not implemented"; return "" }

func (m *Machine) runAllNodes(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *Machine) GetResult(id int64) gorgonia.Value {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value)
}
