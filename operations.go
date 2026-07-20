package gorgonia

import (
	"gorgonia.org/tensor"
)

func binOpNode(op BinaryOp, a, b *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Mul(a, b *Node) (retVal *Node, err error) { _ = "STUB: not implemented"; return nil, nil }

func BatchedMatMul(a, b *Node, transes ...bool) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OuterProd(a, b *Node) (retVal *Node, err error) { _ = "STUB: not implemented"; return nil, nil }

func Div(a, b *Node) (retVal *Node, err error) { _ = "STUB: not implemented"; return nil, nil }

func Auto(op func(a, b *Node, leftPattern, rightPattern []byte) (*Node, error), a, b *Node) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unaryOpNode(op Op, a *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LogSumExp(a *Node, axis int) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func At(a *Node, coords ...int) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Max(a *Node, along ...int) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Mean(a *Node, along ...int) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Sum(a *Node, along ...int) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Norm(a *Node, axis, p int) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReduceAdd(nodes Nodes, opts ...NodeConsOpt) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReduceMul(nodes Nodes, opts ...NodeConsOpt) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SizeOf(axis int, x *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Slice(n *Node, slices ...tensor.Slice) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Transpose(n *Node, axes ...int) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Concat(axis int, ns ...*Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Unconcat(a *Node, along int, n int) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func Reshape(n *Node, to tensor.Shape) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Ravel(n *Node) (retVal *Node, err error) { _ = "STUB: not implemented"; return nil, nil }

func Tensordot(aAxes []int, bAxes []int, a, b *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Mish(a *Node) (retVal *Node, err error) { _ = "STUB: not implemented"; return nil, nil }

func MinBetween(a, b *Node) (retVal *Node, err error) { _ = "STUB: not implemented"; return nil, nil }

func MaxBetween(a, b *Node) (retVal *Node, err error) { _ = "STUB: not implemented"; return nil, nil }

func containsDuplicate(slice []int) bool { _ = "STUB: not implemented"; return false }
