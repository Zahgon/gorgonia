package gorgonia

import (
	"gorgonia.org/tensor"
)

type incrDoerBinOp interface {
	IncrDo(v Value, retSame bool, inputs ...Value) error
}
type usePreallocDoerBinOp interface {
	UsePreallocDo(v Value, retSame bool, inputs ...Value) (retVal Value, err error)
}
type unsafeDoerBinOp interface {
	UnsafeDo(retSame bool, inputs ...Value) (Value, error)
}

type ʘBinaryOperator interface {
	isArith() bool
	binOpType() ʘBinaryOperatorType
	Do(bool, ...Value) (Value, error)
	String() string
}

type scalarBinOp struct {
	ʘBinaryOperatorType
	t tensor.Dtype
}

func (o scalarBinOp) Arity() int { _ = "STUB: not implemented"; return 0 }
func (o scalarBinOp) binOpType() ʘBinaryOperatorType {
	_ = "STUB: not implemented"
	return *new(ʘBinaryOperatorType)
}
func (o scalarBinOp) isArith() bool  { _ = "STUB: not implemented"; return false }
func (o scalarBinOp) String() string { _ = "STUB: not implemented"; return "" }

func (o scalarBinOp) Do(same bool, vals ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

type tBinOp struct {
	ʘBinaryOperatorType
	tensorLeft bool
}

func (o tBinOp) Arity() int { _ = "STUB: not implemented"; return 0 }
func (o tBinOp) binOpType() ʘBinaryOperatorType {
	_ = "STUB: not implemented"
	return *new(ʘBinaryOperatorType)
}
func (o tBinOp) String() string { _ = "STUB: not implemented"; return "" }
func (o tBinOp) isArith() bool  { _ = "STUB: not implemented"; return false }

func (o tBinOp) Do(same bool, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (o tBinOp) UnsafeDo(retSame bool, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (o tBinOp) UsePreallocDo(v Value, retSame bool, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (o tBinOp) IncrDo(incr Value, retSame bool, inputs ...Value) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (o tBinOp) do(vals []Value, opts ...tensor.FuncOpt) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func addDiffExpr(x, y, z, gradZ *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func addDiff(ctx ExecutionContext, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func subDiffExpr(x, y, z, gradZ *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func subDiff(ctx ExecutionContext, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func hadamardProdDiffExpr(x, y, z, gradZ *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func hadamardProdDiff(ctx ExecutionContext, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func hadamardDivDiffExpr(x, y, z, gradZ *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func hadamardDivDiff(ctx ExecutionContext, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func hadamardPowDiffExpr(x, y, z, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func hadamardPowDiff(ctx ExecutionContext, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func nondiffBinOpExpr(x, y, z, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func nondiffBinOp(ctx ExecutionContext, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}
