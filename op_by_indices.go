package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type byIndicesOp struct {
	axis int
}

func newByIndicesOp(axis int) *byIndicesOp { _ = "STUB: not implemented"; return nil }

func ByIndices(x *Node, indices *Node, axis int) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *byIndicesOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *byIndicesOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *byIndicesOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *byIndicesOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *byIndicesOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *byIndicesOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *byIndicesOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *byIndicesOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *byIndicesOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *byIndicesOp) checkInput(inputs ...Value) (x, indices tensor.Tensor, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), *new(tensor.Tensor), nil
}

func (op *byIndicesOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *byIndicesOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *byIndicesOp) SymDiff(inputs Nodes, output, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *byIndicesOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

type byIndicesOpDiffOp struct {
	*byIndicesOp
}

func (op *byIndicesOpDiffOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *byIndicesOpDiffOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *byIndicesOpDiffOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *byIndicesOpDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *byIndicesOpDiffOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *byIndicesOpDiffOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *byIndicesOpDiffOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *byIndicesOpDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *byIndicesOpDiffOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *byIndicesOpDiffOp) checkInput(inputs ...Value) (in, indices, gradient *tensor.Dense, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (op *byIndicesOpDiffOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

var (
	_ Op = &byIndicesOpDiffOp{}

	_ Op   = &byIndicesOp{}
	_ SDOp = &byIndicesOp{}
	_ ADOp = &byIndicesOp{}
)
