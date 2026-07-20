//go:build cuda
// +build cuda

package nnops

import (
	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func Conv2d(im, filter *G.Node, kernelShape tensor.Shape, pad, stride, dilation []int) (retVal *G.Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Conv1d(in, filter *G.Node, kernel, pad, stride, dilation int) (*G.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MaxPool2D(x *G.Node, kernel tensor.Shape, pad, stride []int) (retVal *G.Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Dropout(x *G.Node, prob float64) (retVal *G.Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Rectify(x *G.Node) (retVal *G.Node, err error) { _ = "STUB: not implemented"; return nil, nil }

func BatchNorm(x, scale, bias *G.Node, momentum, epsilon float64) (retVal, γ, β *G.Node, op *BatchNormOp, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}
