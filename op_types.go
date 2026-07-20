package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

func ConvType(x *Node, from, to tensor.Dtype) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type dtConvOp struct {
	inshape  tensor.Shape
	from, to tensor.Dtype
}

func (op *dtConvOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *dtConvOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *dtConvOp) InferShape(_ ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *dtConvOp) Do(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *dtConvOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *dtConvOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *dtConvOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *dtConvOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *dtConvOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *dtConvOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *dtConvOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op *dtConvOp) SymDiff(inputs Nodes, output *Node, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *dtConvOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}
