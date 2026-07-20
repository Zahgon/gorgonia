package gorgonia

import (
	"hash"

	"gorgonia.org/tensor"

	"github.com/chewxy/hm"
)

type upsampleOp struct {
	stride int
}

func newUpsampleOp(inputShape tensor.Shape, stride int) *upsampleOp {
	_ = "STUB: not implemented"
	return nil
}

func Upsample2D(x *Node, scale int) (*Node, error) { _ = "STUB: not implemented"; return nil, nil }

func (op *upsampleOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *upsampleOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *upsampleOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *upsampleOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *upsampleOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *upsampleOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *upsampleOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *upsampleOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *upsampleOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *upsampleOp) checkInput(inputs ...Value) (tensor.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func (op *upsampleOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *upsampleOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op *upsampleOp) SymDiff(inputs Nodes, output, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

type upsampleDiffOp struct {
	upsampleOp
}

func (op *upsampleDiffOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *upsampleDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *upsampleDiffOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *upsampleDiffOp) checkInput(inputs ...Value) (in, pooled, pooledGrad tensor.Tensor, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), *new(tensor.Tensor), *new(tensor.Tensor), nil
}

func (op *upsampleDiffOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}
