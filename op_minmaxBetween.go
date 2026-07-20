package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type minBetween struct{}

func (op minBetween) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op minBetween) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op minBetween) InferShape(shps ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op minBetween) Do(vs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op minBetween) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op minBetween) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op minBetween) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op minBetween) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op minBetween) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op minBetween) String() string { _ = "STUB: not implemented"; return "" }

func (op minBetween) UsePreallocDo(prealloc Value, vs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op minBetween) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }
func (op minBetween) SymDiff(inputs Nodes, output, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op minBetween) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

type maxBetween struct{}

func (op maxBetween) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op maxBetween) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op maxBetween) InferShape(shps ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op maxBetween) Do(vs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op maxBetween) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op maxBetween) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op maxBetween) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op maxBetween) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op maxBetween) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op maxBetween) String() string { _ = "STUB: not implemented"; return "" }

func (op maxBetween) UsePreallocDo(prealloc Value, vs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op maxBetween) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }
func (op maxBetween) SymDiff(inputs Nodes, output, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op maxBetween) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func minmaxSymDiff(a, b *Node, out *Node, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func minmaxAutoDiff(ctx ExecutionContext, a, b *Node, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}
