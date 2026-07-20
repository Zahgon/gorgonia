package gorgonia

import (
	"gorgonia.org/tensor"
)

type ʘUnaryOperator interface {
	unaryOpType() ʘUnaryOperatorType
	String() string
}

type sf32UnaryOperator func(float32) float32
type sf64UnaryOperator func(float64) float64

func unaryCheckApply(op ʘUnaryOperator, t tensor.Tensor, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func nondiffUnaryOpExpr(x, y, gradY *Node) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nondiffUnaryOp(x, y *Node) error { _ = "STUB: not implemented"; return nil }

func absDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func absDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func sinDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sinDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func cosDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cosDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func expDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func lnDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func lnDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func log2DiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func log2Diff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func negDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func negDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func squareDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func squareDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func sqrtDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sqrtDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func inverseDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inverseDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func inverseSqrtDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inverseSqrtDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func cubeDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cubeDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func tanhDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tanhDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func sigmoidDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sigmoidDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func log1pDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func log1pDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func expm1DiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expm1Diff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }

func softplusDiffExpr(x, y, gradY *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func softplusDiff(x, y *Node) (err error) { _ = "STUB: not implemented"; return nil }
