package main

import (
	"gorgonia.org/gorgonia"
)

type convLayer struct {
	filters            int
	padding            int
	kernelSize         int
	stride             int
	activation         string
	activationReLUCoef float64
	batchNormalize     int
	bias               bool
	biases             []float32

	layerIndex int

	convNode *gorgonia.Node
	biasNode *gorgonia.Node
}

func (l *convLayer) String() string { _ = "STUB: not implemented"; return "" }

func (l *convLayer) Type() string { _ = "STUB: not implemented"; return "" }

func (l *convLayer) ToNode(g *gorgonia.ExprGraph, input ...*gorgonia.Node) (*gorgonia.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
