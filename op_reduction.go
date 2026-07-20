package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

func reductionType(d int, along []int) hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func reductionInferShape(along []int, in tensor.Shape) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func reductionDo(op Op, s string, f func(*tensor.Dense, ...int) (*tensor.Dense, error), along []int, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

type maxOp struct {
	along axes
	d     int
}

func newMaxOp(along axes, dim int) *maxOp { _ = "STUB: not implemented"; return nil }

func (op maxOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op maxOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op maxOp) InferShape(dimsizers ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op maxOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }

func (op maxOp) SymDiff(inputs Nodes, output, gradNode *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op maxOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op maxOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op maxOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }
func (op maxOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }

func (op maxOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op maxOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op maxOp) String() string { _ = "STUB: not implemented"; return "" }
func (op maxOp) isUnary() bool  { _ = "STUB: not implemented"; return false }

type sumOp struct {
	along      axes
	d          int
	inputShape tensor.Shape
}

func newSumOp(along axes, s tensor.Shape, d int) sumOp {
	_ = "STUB: not implemented"
	return *new(sumOp)
}

func (op sumOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op sumOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op sumOp) InferShape(inputs ...DimSizer) (shape tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op sumOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }

func (op sumOp) SymDiff(inputs Nodes, output, gradNode *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op sumOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op sumOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op sumOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op sumOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op sumOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op sumOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }
func (op sumOp) Hashcode() uint32      { _ = "STUB: not implemented"; return 0 }
func (op sumOp) String() string        { _ = "STUB: not implemented"; return "" }
func (op sumOp) isUnary() bool         { _ = "STUB: not implemented"; return false }
