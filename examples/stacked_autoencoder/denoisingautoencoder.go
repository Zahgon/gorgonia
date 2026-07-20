package main

import (
	. "gorgonia.org/gorgonia"
)

type DenoisingAutoencoder struct {
	Neuron
	LayerConfig

	h  Neuron
	af ActivationFunction

	input      *Node
	corruption *Node
	corrupted  *Node
	output     *Node
	hiddenOut  *Node
	g          *ExprGraph
}

func NewDATiedWeights(w, b *Node, corruption float64, opts ...LayerConsOpt) *DenoisingAutoencoder {
	_ = "STUB: not implemented"
	return nil
}

func (l *DenoisingAutoencoder) Activate() (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *DenoisingAutoencoder) Reconstruct() (err error) { _ = "STUB: not implemented"; return nil }

func (l *DenoisingAutoencoder) Corrupt() (err error) { _ = "STUB: not implemented"; return nil }

func (l *DenoisingAutoencoder) Cost(x *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
