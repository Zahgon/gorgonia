package main

import (
	"gorgonia.org/gorgonia"
)

type yoloLayer struct {
	masks          []int
	anchors        [][2]int
	flattenAnchors []int
	inputSize      int
	classesNum     int
	ignoreThresh   float32
}

func (l *yoloLayer) String() string { _ = "STUB: not implemented"; return "" }

func (l *yoloLayer) Type() string { _ = "STUB: not implemented"; return "" }

func (l *yoloLayer) ToNode(g *gorgonia.ExprGraph, input ...*gorgonia.Node) (*gorgonia.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
