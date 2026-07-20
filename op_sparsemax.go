package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type sparsemaxOp struct {
	axis int
}

func newSparsemaxOp(axes ...int) *sparsemaxOp { _ = "STUB: not implemented"; return nil }

func Sparsemax(x *Node, axes ...int) (*Node, error) { _ = "STUB: not implemented"; return nil, nil }

func (op *sparsemaxOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *sparsemaxOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *sparsemaxOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *sparsemaxOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *sparsemaxOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *sparsemaxOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *sparsemaxOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *sparsemaxOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *sparsemaxOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *sparsemaxOp) checkInput(inputs ...Value) (tensor.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func (op *sparsemaxOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *sparsemaxOp) float32sparseMax(inputTensor tensor.Tensor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *sparsemaxOp) float64sparseMax(inputTensor tensor.Tensor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *sparsemaxOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *sparsemaxOp) SymDiff(inputs Nodes, output, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *sparsemaxOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

type sparsemaxDiffOp struct {
}

func newSparsemaxOpDiff() *sparsemaxDiffOp { _ = "STUB: not implemented"; return nil }

func (op *sparsemaxDiffOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *sparsemaxDiffOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *sparsemaxDiffOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *sparsemaxDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *sparsemaxDiffOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *sparsemaxDiffOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *sparsemaxDiffOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *sparsemaxDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *sparsemaxDiffOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *sparsemaxDiffOp) checkInput(inputs ...Value) (*tensor.Dense, *tensor.Dense, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (op *sparsemaxDiffOp) mul(a tensor.Tensor, b tensor.Tensor) (tensor.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func (op *sparsemaxDiffOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

var (
	_ Op = &sparsemaxDiffOp{}

	_ Op   = &sparsemaxOp{}
	_ SDOp = &sparsemaxOp{}
	_ ADOp = &sparsemaxOp{}
)
