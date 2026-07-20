package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type randomGenF func() float64

type dropoutOp struct {
	probability float64
	isTraining  bool
	rndGen      randomGenF
}

func newDropoutOp(probability float64, rndGen randomGenF) *dropoutOp {
	_ = "STUB: not implemented"
	return nil
}

func (op *dropoutOp) SetTraining(isTraining bool) error { _ = "STUB: not implemented"; return nil }

func (op *dropoutOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *dropoutOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *dropoutOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *dropoutOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *dropoutOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *dropoutOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *dropoutOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *dropoutOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *dropoutOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *dropoutOp) checkInput(inputs ...Value) (tensor.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func (op *dropoutOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *dropoutOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *dropoutOp) do(input, output interface{}) { _ = "STUB: not implemented"; return }

func (op *dropoutOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *dropoutOp) SymDiff(inputs Nodes, output, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *dropoutOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

type dropoutDiffOp struct {
	*dropoutOp
}

func (op dropoutDiffOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op dropoutDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op dropoutDiffOp) InferShape(ds ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *dropoutDiffOp) Do(values ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *dropoutDiffOp) UsePreallocDo(prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op dropoutDiffOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op dropoutDiffOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op dropoutDiffOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op dropoutDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op dropoutDiffOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op dropoutDiffOp) String() string { _ = "STUB: not implemented"; return "" }

var (
	_ Op          = &dropoutOp{}
	_ ADOp        = &dropoutOp{}
	_ SDOp        = &dropoutOp{}
	_ TrainModeOp = &dropoutOp{}
)
