package gorgonia

import (
	"fmt"
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type DimSizer interface {
	DimSize(int) (int, error)
}

func ShapesToDimSizers(shapes []tensor.Shape) []DimSizer { _ = "STUB: not implemented"; return nil }

func DimSizersToShapes(ds []DimSizer) ([]tensor.Shape, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Op interface {
	Arity() int

	Type() hm.Type

	InferShape(...DimSizer) (tensor.Shape, error)

	Do(...Value) (Value, error)

	ReturnsPtr() bool

	CallsExtern() bool

	OverwritesInput() int

	WriteHash(h hash.Hash)
	Hashcode() uint32
	fmt.Stringer
}

type UnaryOp interface {
	Op

	IsUnary() bool
}

type BinaryOp interface {
	Op

	IsBinary() bool
}

type NoRetOp interface {
	Op

	ReturnsNothing() bool
}

type ADOp interface {
	Op

	DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error
}

type SDOp interface {
	Op

	DiffWRT(inputs int) []bool

	SymDiff(inputs Nodes, output, grad *Node) (retVal Nodes, err error)
}

type ReductionOp interface {
	Op

	IsReduction() bool
}

type TrainModeOp interface {
	SetTraining(isTraining bool) error
}

type IncrDoer interface {
	IncrDo(toIncr Value, inputs ...Value) error
}

type UsePreallocDoer interface {
	UsePreallocDo(prealloc Value, inputs ...Value) (Value, error)
}

type UnsafeDoer interface {
	UnsafeDo(inputs ...Value) (Value, error)
}

type CUDADoer interface {
	CUDADo(extern External, dev Device, prealloc Value, inputs ...Value) (retVal Value, err error)
}

type CLDoer interface {
	CLDo(inputs ...Value) (Value, error)
}

type CUDAADOp interface {
	ADOp
	CUDADoDiff(extern External, dev Device, inputs Nodes, output *Node) error
}

func ApplyOp(op Op, children ...*Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ApplyOpWithName(op Op, name string, children ...*Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type constant interface {
	Op

	isconstant() bool
	Value() Value
}

type constantScalar struct {
	v Scalar
}

func (c constantScalar) Arity() int    { _ = "STUB: not implemented"; return 0 }
func (c constantScalar) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }
func (c constantScalar) InferShape(...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}
func (c constantScalar) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (c constantScalar) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (c constantScalar) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }
func (c constantScalar) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }
func (c constantScalar) SymDiff(Nodes, *Node, *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (c constantScalar) Do(...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}
func (c constantScalar) String() string { _ = "STUB: not implemented"; return "" }

func (c constantScalar) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (c constantScalar) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (c constantScalar) isconstant() bool { _ = "STUB: not implemented"; return false }
func (c constantScalar) Value() Value     { _ = "STUB: not implemented"; return *new(Value) }

type constantTensor struct {
	v tensor.Tensor
}

func (c constantTensor) Arity() int    { _ = "STUB: not implemented"; return 0 }
func (c constantTensor) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }
func (c constantTensor) InferShape(...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (c constantTensor) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (c constantTensor) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }
func (c constantTensor) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (c constantTensor) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }
func (c constantTensor) SymDiff(Nodes, *Node, *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}
func (c constantTensor) Do(...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}
func (c constantTensor) String() string { _ = "STUB: not implemented"; return "" }

func (c constantTensor) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (c constantTensor) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (c constantTensor) isconstant() bool { _ = "STUB: not implemented"; return false }
func (c constantTensor) Value() Value     { _ = "STUB: not implemented"; return *new(Value) }

type Iop struct{}

func (i Iop) Arity() int { _ = "STUB: not implemented"; return 0 }

func (i Iop) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (i Iop) InferShape(ds ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (i Iop) Do(vs ...Value) (Value, error) { _ = "STUB: not implemented"; return *new(Value), nil }

func (i Iop) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (i Iop) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (i Iop) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (i Iop) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (i Iop) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (i Iop) String() string { _ = "STUB: not implemented"; return "" }
