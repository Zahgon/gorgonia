package main

import (
	. "gorgonia.org/gorgonia"
)

type ActivationFunction func(*Node) (*Node, error)

type Layer interface {
	Activate() (*Node, error)
}

type LayerConsOpt func(l Layer)

func WithConf(inputs, outputs, batchSize int) LayerConsOpt {
	_ = "STUB: not implemented"
	return *new(LayerConsOpt)
}

func WithActivationFunction(af ActivationFunction) LayerConsOpt {
	_ = "STUB: not implemented"
	return *new(LayerConsOpt)
}

func WithGraph(g *ExprGraph) LayerConsOpt { _ = "STUB: not implemented"; return *new(LayerConsOpt) }

type FC struct {
	Neuron
	LayerConfig

	af ActivationFunction

	input  *Node
	output *Node
	g      *ExprGraph
}

func NewFC(opts ...LayerConsOpt) *FC { _ = "STUB: not implemented"; return nil }

func (l *FC) Activate() (retVal *Node, err error) { _ = "STUB: not implemented"; return nil, nil }

type SoftmaxLayer struct {
	Neuron
	LayerConfig

	input  *Node
	output *Node
	g      *ExprGraph
}

func NewSoftmaxLayer(opts ...LayerConsOpt) *SoftmaxLayer { _ = "STUB: not implemented"; return nil }

func (l *SoftmaxLayer) Activate() (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
