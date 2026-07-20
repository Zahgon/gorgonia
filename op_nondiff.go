package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type diagFlatOp struct{}

func (op diagFlatOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op diagFlatOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op diagFlatOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op diagFlatOp) Do(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op diagFlatOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op diagFlatOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op diagFlatOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op diagFlatOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op diagFlatOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op diagFlatOp) String() string { _ = "STUB: not implemented"; return "" }
