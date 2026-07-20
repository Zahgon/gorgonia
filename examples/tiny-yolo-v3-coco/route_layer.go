package main

import (
	"gorgonia.org/gorgonia"
)

type routeLayer struct {
	firstLayerIdx  int
	secondLayerIdx int
}

func (l *routeLayer) String() string { _ = "STUB: not implemented"; return "" }

func (l *routeLayer) Type() string { _ = "STUB: not implemented"; return "" }

func (l *routeLayer) ToNode(g *gorgonia.ExprGraph, input ...*gorgonia.Node) (*gorgonia.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
