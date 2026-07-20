package gorgonia

import (
	"math"

	"github.com/chewxy/math32"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/tensor"
)

const (
	maxFloat32 = math32.MaxFloat32
	maxFloat64 = math.MaxFloat64
)

func NodesToValueGrads(in Nodes) (out []ValueGrad) { _ = "STUB: not implemented"; return nil }

func graphNodeToNode(in graph.Nodes) (out Nodes) { _ = "STUB: not implemented"; return *new(Nodes) }

func sliceNodesToNodes(in []graph.Node) (out Nodes) { _ = "STUB: not implemented"; return *new(Nodes) }

func nodeToGraphNode(in []*Node) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func tensorInfo(t tensor.Tensor) (dt tensor.Dtype, dim int) {
	_ = "STUB: not implemented"
	return *new(tensor.Dtype), 0
}

func valueToInt(v Value) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func valuesToInts(values []Value) (retVal []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func valuesToTensors(values []Value) (retVal []tensor.Tensor, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func intRange(start, end int) []int { _ = "STUB: not implemented"; return nil }

func ones(dt tensor.Dtype, sizes ...int) (retVal Value) {
	_ = "STUB: not implemented"
	return *new(Value)
}

func hasInf(v Value, dev Device) bool { _ = "STUB: not implemented"; return false }

func hasNaN(v Value, dev Device) bool { _ = "STUB: not implemented"; return false }

func setZero(val Value) (retVal Value) { _ = "STUB: not implemented"; return *new(Value) }

func checkArity(op arityer, inputs int) error { _ = "STUB: not implemented"; return nil }

func maxInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func minInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func ceilDivInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func simpleHash(op hashWriter) uint32 { _ = "STUB: not implemented"; return 0 }

func getDV(x, y *Node) (xdv, ydv *dualValue) { _ = "STUB: not implemented"; return nil, nil }

func getDV3(x, y, z *Node) (xdv, ydv, zdv *dualValue) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getConst(x *Node, constant string) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scalarEquiv(s tensor.Shape) bool { _ = "STUB: not implemented"; return false }

func runInParallel(from, to int, cb func(i int)) { _ = "STUB: not implemented"; return }
