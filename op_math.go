package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type elemBinOp struct {
	ʘBinaryOperator
	arg0, arg1 hm.Type
	retSame    bool
}

func newEBOByType(ot ʘBinaryOperatorType, at, bt hm.Type) elemBinOp {
	_ = "STUB: not implemented"
	return *new(elemBinOp)
}

func newElemBinOp(ot ʘBinaryOperatorType, a, b *Node) elemBinOp {
	_ = "STUB: not implemented"
	return *new(elemBinOp)
}

func (op elemBinOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op elemBinOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op elemBinOp) InferShape(inputs ...DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op elemBinOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op elemBinOp) SymDiff(inputs Nodes, output, gradNode *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op elemBinOp) Do(values ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op elemBinOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op elemBinOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op elemBinOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op elemBinOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op elemBinOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op elemBinOp) UsePreallocDo(prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op elemBinOp) UnsafeDo(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op elemBinOp) IncrDo(incr Value, inputs ...Value) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op elemBinOp) String() string { _ = "STUB: not implemented"; return "" }

func (op elemBinOp) IsBinary() bool { _ = "STUB: not implemented"; return false }

type elemUnaryOp struct {
	ʘUnaryOperator

	argTensor     bool
	numericResult bool
}

func newElemUnaryOp(op ʘUnaryOperatorType, a *Node) elemUnaryOp {
	_ = "STUB: not implemented"
	return *new(elemUnaryOp)
}

func (op elemUnaryOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op elemUnaryOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op elemUnaryOp) InferShape(inputs ...DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op elemUnaryOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op elemUnaryOp) SymDiff(inputs Nodes, output, gradNode *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op elemUnaryOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op elemUnaryOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op elemUnaryOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op elemUnaryOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op elemUnaryOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op elemUnaryOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op elemUnaryOp) UnsafeDo(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op elemUnaryOp) isUnary() bool { _ = "STUB: not implemented"; return false }

func (op elemUnaryOp) do(a Value, opts ...tensor.FuncOpt) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

type linAlgBinOp struct {
	āBinaryOperator
	transA, transB bool
}

func (op linAlgBinOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op linAlgBinOp) InferShape(inputs ...DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op linAlgBinOp) SymDiff(inputs Nodes, output, gradNode *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op linAlgBinOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op linAlgBinOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}
func (op linAlgBinOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op linAlgBinOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op linAlgBinOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op linAlgBinOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op linAlgBinOp) String() string { _ = "STUB: not implemented"; return "" }

func (op linAlgBinOp) IncrDo(incr Value, inputs ...Value) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op linAlgBinOp) UsePreallocDo(prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op linAlgBinOp) IsBinary() bool { _ = "STUB: not implemented"; return false }

func (op linAlgBinOp) do(inputs []Value, opts ...tensor.FuncOpt) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op linAlgBinOp) preallocBatchMatMul(incr bool, prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

type tensordotOp struct {
	aAxes   []int
	bAxes   []int
	aDims   int
	bDims   int
	retDims int
}

func makeTensordotOp(a, b *Node, aAxes, bAxes []int) tensordotOp {
	_ = "STUB: not implemented"
	return *new(tensordotOp)
}

func (op tensordotOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op tensordotOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op tensordotOp) InferShape(ds ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op tensordotOp) Do(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op tensordotOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op tensordotOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op tensordotOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op tensordotOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op tensordotOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op tensordotOp) String() string { _ = "STUB: not implemented"; return "" }

func (op tensordotOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op tensordotOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op tensordotOp) SymDiff(inputs Nodes, output *Node, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}
