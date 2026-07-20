package gorgonia

import (
	"gorgonia.org/tensor"
)

func Must(n *Node, err error, opts ...NodeConsOpt) *Node { _ = "STUB: not implemented"; return nil }

func NodeFromAny(g *ExprGraph, any interface{}, opts ...NodeConsOpt) *Node {
	_ = "STUB: not implemented"
	return nil
}

func NewScalar(g *ExprGraph, t tensor.Dtype, opts ...NodeConsOpt) *Node {
	_ = "STUB: not implemented"
	return nil
}

func NewVector(g *ExprGraph, t tensor.Dtype, opts ...NodeConsOpt) *Node {
	_ = "STUB: not implemented"
	return nil
}

func NewMatrix(g *ExprGraph, t tensor.Dtype, opts ...NodeConsOpt) *Node {
	_ = "STUB: not implemented"
	return nil
}

func NewTensor(g *ExprGraph, t tensor.Dtype, dims int, opts ...NodeConsOpt) *Node {
	_ = "STUB: not implemented"
	return nil
}

func NewConstant(v interface{}, opts ...NodeConsOpt) *Node { _ = "STUB: not implemented"; return nil }

func UniformRandomNode(g *ExprGraph, dt tensor.Dtype, low, high float64, shape ...int) *Node {
	_ = "STUB: not implemented"
	return nil
}

func GaussianRandomNode(g *ExprGraph, dt tensor.Dtype, mean, stdev float64, shape ...int) *Node {
	_ = "STUB: not implemented"
	return nil
}

func BinomialRandomNode(g *ExprGraph, dt tensor.Dtype, trials, prob float64, shape ...int) *Node {
	_ = "STUB: not implemented"
	return nil
}

func OneHotVector(id, classes int, t tensor.Dtype, opts ...NodeConsOpt) *Node {
	_ = "STUB: not implemented"
	return nil
}

func Grad(cost *Node, WRTs ...*Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func Let(n *Node, be interface{}) error { _ = "STUB: not implemented"; return nil }

func UnsafeLet(n *Node, be interface{}) error { _ = "STUB: not implemented"; return nil }

func Set(a, b *Node) (retVal *Node) { _ = "STUB: not implemented"; return nil }

func Read(n *Node, into *Value) (retVal *Node) { _ = "STUB: not implemented"; return nil }
