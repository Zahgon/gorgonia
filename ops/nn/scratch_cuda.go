//go:build cuda
// +build cuda

package nnops

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type gpuScratchOp struct {
	scratchOp
}

func (op *gpuScratchOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *gpuScratchOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *gpuScratchOp) InferShape(...gorgonia.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}
func (op *gpuScratchOp) Do(...gorgonia.Value) (gorgonia.Value, error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}
func (op *gpuScratchOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op *gpuScratchOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op *gpuScratchOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op *gpuScratchOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *gpuScratchOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }
func (op *gpuScratchOp) String() string   { _ = "STUB: not implemented"; return "" }

func (op *gpuScratchOp) CUDADo(extern gorgonia.External, dev gorgonia.Device, prealloc gorgonia.Value, inputs ...gorgonia.Value) (retVal gorgonia.Value, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}
