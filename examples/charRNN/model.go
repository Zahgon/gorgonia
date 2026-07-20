package main

import (
	. "gorgonia.org/gorgonia"
)

var hiddenSizes = []int{100}
var embeddingSize = 10

type layer struct {
	wix    Value
	wih    Value
	bias_i Value

	wfx    Value
	wfh    Value
	bias_f Value

	wox    Value
	woh    Value
	bias_o Value

	wcx    Value
	wch    Value
	bias_c Value
}

type lstm struct {
	wix    *Node
	wih    *Node
	bias_i *Node

	wfx    *Node
	wfh    *Node
	bias_f *Node

	wox    *Node
	woh    *Node
	bias_o *Node

	wcx    *Node
	wch    *Node
	bias_c *Node
}

func newLSTMLayer(g *ExprGraph, l *layer, name string) *lstm { _ = "STUB: not implemented"; return nil }

func (l *lstm) fwd(inputVector, prevHidden, prevCell *Node) (hidden, cell *Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

type model struct {
	ls []*layer

	whd    Value
	bias_d Value

	embedding Value

	inputSize, embeddingSize, outputSize int
	hiddenSizes                          []int

	prefix string
	free   bool
}

type lstmOut struct {
	hiddens Nodes
	cells   Nodes

	probs *Node
}

func NewLSTMModel(inputSize, embeddingSize, outputSize int, hiddenSizes []int) *model {
	_ = "STUB: not implemented"
	return nil
}

type charRNN struct {
	*model
	g  *ExprGraph
	ls []*lstm

	whd    *Node
	bias_d *Node

	embedding *Node

	inputVector *Node
	prevHiddens Nodes
	prevCells   Nodes
}

func newCharRNN(m *model) *charRNN { _ = "STUB: not implemented"; return nil }

func (r *charRNN) learnables() (retVal []ValueGrad) { _ = "STUB: not implemented"; return nil }

func (r *charRNN) fwd(srcIndex int, prev *lstmOut) (retVal *lstmOut, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *charRNN) costFn(sentence string) (cost, perplexity *Node, n int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

func (r *charRNN) predict() { _ = "STUB: not implemented"; return }

func (r *charRNN) cleanup() { _ = "STUB: not implemented"; return }

func run(r *charRNN, iter int, solver Solver) (retCost, retPerp float32, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
