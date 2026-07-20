package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type atOp struct {
	coordinates coordinates
	d           int
}

func (op atOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op atOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op atOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op atOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }
func (op atOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (op atOp) InferShape(...DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}
func (op atOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }
func (op atOp) SymDiff(Nodes, *Node, *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}
func (op atOp) String() string { _ = "STUB: not implemented"; return "" }

func (op atOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op atOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op atOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op atOp) isStmt() bool { _ = "STUB: not implemented"; return false }

type sizeOp struct {
	axis, d int
	val     int
}

func (op sizeOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op sizeOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op sizeOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op sizeOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }
func (op sizeOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (op sizeOp) InferShape(...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}
func (op sizeOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }
func (op sizeOp) String() string       { _ = "STUB: not implemented"; return "" }

func (op sizeOp) SymDiff(inputs Nodes, output, gradNode *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op sizeOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op sizeOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op sizeOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op sizeOp) DimSize(d int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type repeatOp struct {
	along      int
	inputShape tensor.Shape
}

func newRepeatOp(along int, a *Node) *repeatOp { _ = "STUB: not implemented"; return nil }

func repeatedApply(along []int, children Nodes) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op repeatOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op repeatOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op repeatOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op repeatOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }
func (op repeatOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }

func (op repeatOp) InferShape(inputs ...DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op repeatOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }

func (op repeatOp) SymDiff(inputs Nodes, output, gradNode *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op repeatOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op repeatOp) String() string { _ = "STUB: not implemented"; return "" }

func (op repeatOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op repeatOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op repeatOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op repeatOp) UsePreallocDo(prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

type sliceOp struct {
	tensor.Slice

	along int

	a int
	d int
}

func (op *sliceOp) IsSlice() tensor.Slice { _ = "STUB: not implemented"; return *new(tensor.Slice) }

func newSliceOp(s tensor.Slice, along, d int) *sliceOp { _ = "STUB: not implemented"; return nil }

func (op *sliceOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *sliceOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *sliceOp) InferShape(inputs ...DimSizer) (s tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *sliceOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }

func (op *sliceOp) SymDiff(inputs Nodes, outputNode, gradNode *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *sliceOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *sliceOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *sliceOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op *sliceOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (op *sliceOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }
func (op sliceOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op sliceOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op sliceOp) String() string { _ = "STUB: not implemented"; return "" }

func (op sliceOp) all() bool { _ = "STUB: not implemented"; return false }

type sliceIncrOp struct {
	*sliceOp
}

func (op sliceIncrOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op sliceIncrOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op sliceIncrOp) InferShape(inputs ...DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op sliceIncrOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }

func (op sliceIncrOp) SymDiff(inputs Nodes, outputNode, gradNode *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op sliceIncrOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op sliceIncrOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op sliceIncrOp) UsePreallocDo(prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op sliceIncrOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op sliceIncrOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op sliceIncrOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op sliceIncrOp) String() string { _ = "STUB: not implemented"; return "" }

type transposeOp struct {
	pattern []int
	d       int
}

func (op transposeOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op transposeOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op transposeOp) InferShape(inputs ...DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op transposeOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }

func (op transposeOp) SymDiff(inputs Nodes, outputNode, gradNode *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op transposeOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op transposeOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op transposeOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op transposeOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (op transposeOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op transposeOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op transposeOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op transposeOp) String() string { _ = "STUB: not implemented"; return "" }

type concatOp struct {
	axis     int
	d        int
	children int
}

func (op concatOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op concatOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op concatOp) InferShape(ds ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op concatOp) Do(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op concatOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op concatOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (op concatOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op concatOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op concatOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op concatOp) String() string { _ = "STUB: not implemented"; return "" }

func (op concatOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op concatOp) SymDiff(inputs Nodes, output *Node, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op concatOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

type reshapeOp struct {
	from, to tensor.Shape
}

func (op reshapeOp) Arity() int    { _ = "STUB: not implemented"; return 0 }
func (op reshapeOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op reshapeOp) InferShape(ds ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op reshapeOp) Do(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op reshapeOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op reshapeOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op reshapeOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op reshapeOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op reshapeOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op reshapeOp) String() string { _ = "STUB: not implemented"; return "" }

func (op reshapeOp) UnsafeDo(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op reshapeOp) CUDADo(extern External, dev Device, prealloc Value, vals ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op reshapeOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }

func (op reshapeOp) SymDiff(inputs Nodes, output *Node, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op reshapeOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func contains(slice []int, value int) int { _ = "STUB: not implemented"; return 0 }

func sortUniqueIntWithImitator(toBeSorted, imitator []int) { _ = "STUB: not implemented"; return }
