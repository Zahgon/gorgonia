package main

import (
	. "gorgonia.org/gorgonia"
)

type Neuron struct {
	w *Node
	b *Node

	g *ExprGraph
}

type initFn func() InitWFn

func MakeNeuron(inputs, outputs, batchSize int, g *ExprGraph, fn initFn) Neuron {
	_ = "STUB: not implemented"
	return *new(Neuron)
}

func (n *Neuron) GobEncode() (p []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (n *Neuron) GobDecode(p []byte) (err error) { _ = "STUB: not implemented"; return nil }
