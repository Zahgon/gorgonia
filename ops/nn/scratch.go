package nnops

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type scratchOp struct {
	shape tensor.Shape
	dt    tensor.Dtype
	name  string
}

func (op *scratchOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *scratchOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *scratchOp) InferShape(...gorgonia.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}
func (op *scratchOp) Do(...gorgonia.Value) (gorgonia.Value, error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}
func (op *scratchOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op *scratchOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op *scratchOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op *scratchOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *scratchOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }
func (op *scratchOp) String() string   { _ = "STUB: not implemented"; return "" }

func (op *scratchOp) UsePreallocDo(prealloc gorgonia.Value, inputs ...gorgonia.Value) (gorgonia.Value, error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}
