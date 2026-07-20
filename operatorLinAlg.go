package gorgonia

import (
	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type āBinaryOperator byte

const (
	matMulOperator āBinaryOperator = iota
	matVecMulOperator
	vecDotOperator
	outerProdOperator
	batchedMatMulOperator

	maxĀBinaryOperator
)

func (op āBinaryOperator) String() string { _ = "STUB: not implemented"; return "" }

func (op āBinaryOperator) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op āBinaryOperator) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func matMulDiffExpr(transA, transB bool, x, y, z, gradZ *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func matMulDiff(ctx ExecutionContext, transA, transB bool, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func matVecMulDiffExpr(transA, transB bool, x, y, z, gradZ *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func matVecMulDiff(ctx ExecutionContext, transA, transB bool, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func vecDotDiffExpr(transA, transB bool, x, y, z, gradZ *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func vecDotDiff(ctx ExecutionContext, transA, transB bool, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func outerProdDiffExpr(transA, transB bool, x, y, z, gradZ *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func outerProdDiff(ctx ExecutionContext, transA, transB bool, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func batchedMatMulDiffExpr(transA, transB bool, x, y, z, gradZ *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func batchedMatMulDiff(ctx ExecutionContext, transA, transB bool, x, y, z *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func reshape(name string, t tensor.Tensor, shape ...int) error {
	_ = "STUB: not implemented"
	return nil
}

func batchedMatMul(a, b, c tensor.Tensor, transA, transB, incr bool) (retVal tensor.Tensor, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func incrSlices(a []sli, shp tensor.Shape) (halt bool) { _ = "STUB: not implemented"; return false }
