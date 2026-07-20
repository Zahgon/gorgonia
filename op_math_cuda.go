//go:build cuda
// +build cuda

package gorgonia

import (
	"gorgonia.org/gorgonia/cuda"
)

const (
	elemBinOpMod   = "elembinop"
	elemUnaryOpMod = "elemunaryop"
)

func (op elemUnaryOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op elemUnaryOp) CUDADo(extern External, dev Device, prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op elemBinOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op elemBinOp) CUDADo(extern External, dev Device, prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op elemBinOp) ssop(a, b, prealloc Value, e *cuda.Engine) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op linAlgBinOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op linAlgBinOp) CUDADo(extern External, dev Device, prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func NewAddOp(a, b *Node, ctx ExecutionContext) *ExternalOp { _ = "STUB: not implemented"; return nil }

func NewSubOp(a, b *Node, ctx ExecutionContext) *ExternalOp { _ = "STUB: not implemented"; return nil }

func NewHadamardProdOp(a, b *Node, ctx ExecutionContext) *ExternalOp {
	_ = "STUB: not implemented"
	return nil
}
