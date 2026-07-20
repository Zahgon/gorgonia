package gorgonia

import "github.com/chewxy/hm"

var āBinOpStrs = [maxĀBinaryOperator]string{
	"×",
	"×",
	"⋅",
	"⊗",
	"×××",
}

var āBinOpDiffExprs = [maxĀBinaryOperator]func(tA, tB bool, x, y, z, grad *Node) (Nodes, error){
	matMulDiffExpr,
	matVecMulDiffExpr,
	vecDotDiffExpr,
	outerProdDiffExpr,
	batchedMatMulDiffExpr,
}

var āBinOpDiffs = [maxĀBinaryOperator]func(ctx ExecutionContext, tA, tB bool, x, y, z *Node) error{
	matMulDiff,
	matVecMulDiff,
	vecDotDiff,
	outerProdDiff,
	batchedMatMulDiff,
}

var āBinOpTypes = [maxĀBinaryOperator]func() hm.Type{
	matMulType,
	matVecMulType,
	vecDotType,
	outerProdType,
	batchedMatMulType,
}

func matVecMulType() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func matMulType() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func vecDotType() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func outerProdType() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func batchedMatMulType() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }
