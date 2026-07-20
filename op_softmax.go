package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

func SoftMax(x *Node, axes ...int) (*Node, error) { _ = "STUB: not implemented"; return nil, nil }

func LogSoftMax(x *Node, axes ...int) (*Node, error) { _ = "STUB: not implemented"; return nil, nil }

type softmaxOp struct {
	shape tensor.Shape
	axis  int
	isLog bool
}

func newSoftmaxOp(inputShape tensor.Shape, axes ...int) *softmaxOp {
	_ = "STUB: not implemented"
	return nil
}

func (op *softmaxOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *softmaxOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *softmaxOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *softmaxOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *softmaxOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *softmaxOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *softmaxOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *softmaxOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *softmaxOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *softmaxOp) checkInput(inputs ...Value) (tensor.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func (op *softmaxOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *softmaxOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *softmaxOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *softmaxOp) SymDiff(inputs Nodes, output, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *softmaxOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

type softmaxDiffOp struct {
	*softmaxOp
}

func (op *softmaxDiffOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *softmaxDiffOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *softmaxDiffOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *softmaxDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *softmaxDiffOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *softmaxDiffOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *softmaxDiffOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *softmaxDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *softmaxDiffOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *softmaxDiffOp) checkInput(inputs ...Value) (tensor.Tensor, tensor.Tensor, tensor.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), *new(tensor.Tensor), *new(tensor.Tensor), nil
}

func (op *softmaxDiffOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *softmaxDiffOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

var (
	_ Op   = &softmaxOp{}
	_ ADOp = &softmaxOp{}
	_ SDOp = &softmaxOp{}

	_ Op = &softmaxDiffOp{}
)
