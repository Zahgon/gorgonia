//go:build cuda
// +build cuda

package nnops

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type activation struct {
	*cudnn.Activation
	xDesc, yDesc *cudnn.TensorDescriptor
}

func newRelu() (*activation, error) { _ = "STUB: not implemented"; return nil, nil }

func (op *activation) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *activation) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *activation) InferShape(inputs ...gorgonia.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *activation) Do(...gorgonia.Value) (gorgonia.Value, error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}

func (op *activation) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *activation) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *activation) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *activation) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *activation) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *activation) String() string { _ = "STUB: not implemented"; return "" }

func (op *activation) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op *activation) SymDiff(inputs gorgonia.Nodes, output *gorgonia.Node, grad *gorgonia.Node) (retVal gorgonia.Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Nodes), nil
}

func (op *activation) CUDADo(extern gorgonia.External, dev gorgonia.Device, prealloc gorgonia.Value, inputs ...gorgonia.Value) (retVal gorgonia.Value, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}

type activationDiff struct {
	*activation
	dyDesc, dxDesc *cudnn.TensorDescriptor
}

func (op *activationDiff) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *activationDiff) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *activationDiff) InferShape(inputs ...gorgonia.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *activationDiff) Do(...gorgonia.Value) (gorgonia.Value, error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}

func (op *activationDiff) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *activationDiff) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *activationDiff) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *activationDiff) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *activationDiff) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *activationDiff) String() string { _ = "STUB: not implemented"; return "" }

func (op *activationDiff) CUDADo(extern gorgonia.External, dev gorgonia.Device, prealloc gorgonia.Value, inputs ...gorgonia.Value) (retVal gorgonia.Value, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}
