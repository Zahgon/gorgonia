package gorgonia

import (
	"gorgonia.org/tensor"
)

var scalarShape = tensor.ScalarShape()

type axes []int
type coordinates []int

func transpose2D(shape tensor.Shape) tensor.Shape {
	_ = "STUB: not implemented"
	return *new(tensor.Shape)
}

func transposeBatch2D(shape tensor.Shape) tensor.Shape {
	_ = "STUB: not implemented"
	return *new(tensor.Shape)
}

func calcBroadcastShape(a *Node, expectedDims int, broadcastAlong []int) (newShape tensor.Shape) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape)
}

func KeepDims(a *Node, expandLeft bool, fn func(a *Node) (*Node, error)) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
