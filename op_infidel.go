package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type stmtOp interface {
	Op
	isStmt() bool
}

type letOp struct{}

func (op letOp) Arity() int           { _ = "STUB: not implemented"; return 0 }
func (op letOp) Type() hm.Type        { _ = "STUB: not implemented"; return *new(hm.Type) }
func (op letOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op letOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }
func (op letOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (op letOp) InferShape(...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}
func (op letOp) DiffWRT(int) []bool { _ = "STUB: not implemented"; return nil }
func (op letOp) SymDiff(inputs Nodes, outputNode, gradNode *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}
func (op letOp) Do(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}
func (op letOp) String() string        { _ = "STUB: not implemented"; return "" }
func (op letOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }
func (op letOp) Hashcode() uint32      { _ = "STUB: not implemented"; return 0 }

func (op letOp) isStmt() bool { _ = "STUB: not implemented"; return false }

type readOp struct {
	into *Value
}

func (op readOp) Arity() int           { _ = "STUB: not implemented"; return 0 }
func (op readOp) Type() hm.Type        { _ = "STUB: not implemented"; return *new(hm.Type) }
func (op readOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op readOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }
func (op readOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (op readOp) InferShape(...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}
func (op readOp) DiffWRT(int) []bool { _ = "STUB: not implemented"; return nil }
func (op readOp) SymDiff(inputs Nodes, outputNode, gradNode *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}
func (op readOp) Do(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}
func (op readOp) String() string        { _ = "STUB: not implemented"; return "" }
func (op readOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }
func (op readOp) Hashcode() uint32      { _ = "STUB: not implemented"; return 0 }

func (op readOp) isStmt() bool { _ = "STUB: not implemented"; return false }

type devTrans struct {
	from, to Device
	toNode   *Node
}

func (op devTrans) Arity() int    { _ = "STUB: not implemented"; return 0 }
func (op devTrans) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }
func (op devTrans) InferShape(...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}
func (op devTrans) Do(...Value) (Value, error) { _ = "STUB: not implemented"; return *new(Value), nil }
func (op devTrans) ReturnsPtr() bool           { _ = "STUB: not implemented"; return false }
func (op devTrans) CallsExtern() bool          { _ = "STUB: not implemented"; return false }
func (op devTrans) OverwritesInput() int       { _ = "STUB: not implemented"; return 0 }
func (op devTrans) WriteHash(h hash.Hash)      { _ = "STUB: not implemented"; return }
func (op devTrans) Hashcode() uint32           { _ = "STUB: not implemented"; return 0 }

func (op devTrans) String() string { _ = "STUB: not implemented"; return "" }
func (op devTrans) isStmt() bool   { _ = "STUB: not implemented"; return false }

func (op devTrans) CUDADo(extern External, dev Device, prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op devTrans) CUDAFuncName() string { _ = "STUB: not implemented"; return "" }
