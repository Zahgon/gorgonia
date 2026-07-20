package main

import (
	"gorgonia.org/gorgonia"
)

type maxPoolingLayer struct {
	size   int
	stride int
}

func (l *maxPoolingLayer) String() string { _ = "STUB: not implemented"; return "" }

func (l *maxPoolingLayer) Type() string { _ = "STUB: not implemented"; return "" }

func (l *maxPoolingLayer) ToNode(g *gorgonia.ExprGraph, input ...*gorgonia.Node) (*gorgonia.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
