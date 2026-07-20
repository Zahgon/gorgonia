package main

import (
	"gorgonia.org/gorgonia"
)

type upsampleLayer struct {
	scale int
}

func (l *upsampleLayer) String() string { _ = "STUB: not implemented"; return "" }

func (l *upsampleLayer) Type() string { _ = "STUB: not implemented"; return "" }

func (l *upsampleLayer) ToNode(g *gorgonia.ExprGraph, input ...*gorgonia.Node) (*gorgonia.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
