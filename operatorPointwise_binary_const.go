package gorgonia

import "gorgonia.org/tensor"

var (
	tadd = denseBinOp(tensor.Add)
	tsub = denseBinOp(tensor.Sub)
	tmul = denseBinOp(tensor.Mul)
	tdiv = denseBinOp(tensor.Div)
	tpow = denseBinOp(tensor.Pow)

	tlt  = denseCmpOp(tensor.Lt)
	tgt  = denseCmpOp(tensor.Gt)
	tlte = denseCmpOp(tensor.Lte)
	tgte = denseCmpOp(tensor.Gte)
	teq  = denseCmpOp(tensor.ElEq)
	tne  = denseCmpOp(tensor.ElNe)
)

type denseBinOp func(a, b interface{}, opts ...tensor.FuncOpt) (tensor.Tensor, error)
type denseCmpOp func(a, b interface{}, opts ...tensor.FuncOpt) (tensor.Tensor, error)

type ʘBinaryOperatorType byte

const (
	addOpType ʘBinaryOperatorType = iota
	subOpType
	mulOpType
	divOpType
	powOpType

	ltOpType
	gtOpType
	lteOpType
	gteOpType
	eqOpType
	neOpType

	maxʘBinaryOpType
)

func (op ʘBinaryOperatorType) String() string { _ = "STUB: not implemented"; return "" }

var ʘBinOpStrs = [maxʘBinaryOpType]string{

	"+",
	"-",
	"⊙",
	"÷",
	"^",

	"<",
	">",
	"<=",
	">=",
	"==",
	"!=",
}

var ʘBinOpNames = [maxʘBinaryOpType]string{

	"add",
	"sub",
	"mul",
	"div",
	"pow",

	"lt",
	"gt",
	"lte",
	"gte",
	"eq",
	"ne",
}

var ʘBinOpCommutative = [maxʘBinaryOpType]bool{
	true, false, true, false, false,
	false, false, false, false, true, true,
}

var ʘBinOpDiffExprs = [maxʘBinaryOpType]func(x, y, z, gradZ *Node) (Nodes, error){
	addDiffExpr, subDiffExpr, hadamardProdDiffExpr, hadamardDivDiffExpr, hadamardPowDiffExpr,
	nondiffBinOpExpr, nondiffBinOpExpr, nondiffBinOpExpr, nondiffBinOpExpr, nondiffBinOpExpr, nondiffBinOpExpr,
}

var ʘBinOpDiffFns = [maxʘBinaryOpType]func(ctx ExecutionContext, x, y, z *Node) error{
	addDiff, subDiff, hadamardProdDiff, hadamardDivDiff, hadamardPowDiff,
	nondiffBinOp, nondiffBinOp, nondiffBinOp, nondiffBinOp, nondiffBinOp, nondiffBinOp,
}

func (op ʘBinaryOperatorType) isCommutative() bool { _ = "STUB: not implemented"; return false }

func (op ʘBinaryOperatorType) diffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op ʘBinaryOperatorType) isArith() bool { _ = "STUB: not implemented"; return false }

var binOps = [maxʘBinaryOpType]*denseBinOp{
	&tadd,
	&tsub,
	&tmul,
	&tdiv,
	&tpow,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
}

var cmpOps = [maxʘBinaryOpType]*denseCmpOp{
	nil,
	nil,
	nil,
	nil,
	nil,
	&tlt,
	&tgt,
	&tlte,
	&tgte,
	&teq,
	&tne,
}
