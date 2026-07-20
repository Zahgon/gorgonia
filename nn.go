package gorgonia

import (
	"gorgonia.org/tensor"
)

func BinaryXent(output, target *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Dropout(x *Node, dropProb float64) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LeakyRelu(x *Node, alpha float64) (*Node, error) { _ = "STUB: not implemented"; return nil, nil }

func Rectify(x *Node) (retVal *Node, err error) { _ = "STUB: not implemented"; return nil, nil }

func Im2Col(n *Node, kernel, pad, stride, dilation tensor.Shape) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Conv2d(im, filter *Node, kernelShape tensor.Shape, pad, stride, dilation []int) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Conv1d(in, filter *Node, kernel, pad, stride, dilation int) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MaxPool2D(x *Node, kernel tensor.Shape, pad, stride []int) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MaxPool1D(x *Node, kernel, pad, stride int) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BatchNorm(x, scale, bias *Node, momentum, epsilon float64) (retVal, γ, β *Node, op *BatchNormOp, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

func GlobalAveragePool2D(x *Node) (*Node, error) { _ = "STUB: not implemented"; return nil, nil }
