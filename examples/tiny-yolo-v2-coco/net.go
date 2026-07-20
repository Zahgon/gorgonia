package main

import (
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type TinyYOLOv2Net struct {
	g                                                                                                                                *gorgonia.ExprGraph
	classesNum, boxesPerCell                                                                                                         int
	convWeights0, convWeights2, convWeights4, convWeights6, convWeights8, convWeights10, convWeights12, convWeights13, convWeights14 *gorgonia.Node

	out *gorgonia.Node

	biases  map[string][]float32
	gammas  map[string][]float32
	means   map[string][]float32
	vars    map[string][]float32
	kernels map[string][]float32
}

func (tiny *TinyYOLOv2Net) GetOutput() *gorgonia.Node { _ = "STUB: not implemented"; return nil }

func NewTinyYOLOv2Net(g *gorgonia.ExprGraph, classesNumber int, boxesPerCell int, weightsFile string) *TinyYOLOv2Net {
	_ = "STUB: not implemented"
	return nil
}

func PrepareConv(g *gorgonia.ExprGraph, shape tensor.Shape, weights []float32, layerName string) *gorgonia.Node {
	_ = "STUB: not implemented"
	return nil
}

func PrepareBiases(g *gorgonia.ExprGraph, shape tensor.Shape, biases map[string][]float32, layerName string, biasName string) *gorgonia.Node {
	_ = "STUB: not implemented"
	return nil
}
