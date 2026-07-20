package main

import (
	. "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type StackedDA struct {
	DeepConfig

	autoencoders []*DenoisingAutoencoder
	hiddenLayers []*FC
	final        *SoftmaxLayer

	input *Node
	g     *ExprGraph
}

func NewStackedDA(g *ExprGraph, batchSize, size, inputs, outputs, layers int, hiddenSizes []int, corruptions []float64) *StackedDA {
	_ = "STUB: not implemented"
	return nil
}

func (sda *StackedDA) Pretrain(x tensor.Tensor, epoch int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (sda *StackedDA) Finetune(x tensor.Tensor, y []int, epoch int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (sda *StackedDA) Forwards(x tensor.Tensor) (res tensor.Tensor, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func (sda *StackedDA) Save(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func (sda *StackedDA) Load(filename string) (err error) { _ = "STUB: not implemented"; return nil }
