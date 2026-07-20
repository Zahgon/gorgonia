package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type Reduction uint

const (
	ReductionMean Reduction = iota
	ReductionSum
)

func CTCLoss(logProbs, targets, inputLengths, targetLengths *Node, reduction Reduction) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ctcLossOp struct {
	dtype      tensor.Dtype
	targetDims int
	reduction  Reduction

	logAlpha         *tensor.Dense
	negLogLikelihood *tensor.Dense
}

func newCTCLossOp(dtype tensor.Dtype, targetDims int, reduction Reduction) *ctcLossOp {
	_ = "STUB: not implemented"
	return nil
}

func (op *ctcLossOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *ctcLossOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *ctcLossOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *ctcLossOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *ctcLossOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *ctcLossOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *ctcLossOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *ctcLossOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *ctcLossOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *ctcLossOp) getPrimeTarget(targets []int, offset, stride, idx int) int {
	_ = "STUB: not implemented"
	return 0
}

func (op *ctcLossOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *ctcLossOp) f64s(logProbsT, prealloc, targetsT, inputLengthsT, targetLengthsT *tensor.Dense) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *ctcLossOp) f32s(logProbsT, prealloc, targetsT, inputLengthsT, targetLengthsT *tensor.Dense) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *ctcLossOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *ctcLossOp) SymDiff(inputs Nodes, output, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *ctcLossOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

type ctcLossDiffOp struct {
	*ctcLossOp
}

func (op *ctcLossDiffOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *ctcLossDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *ctcLossDiffOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *ctcLossDiffOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *ctcLossDiffOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *ctcLossDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *ctcLossDiffOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *ctcLossDiffOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *ctcLossDiffOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *ctcLossDiffOp) f64s(logProbsT, targetsT, inputLengthsT, targetLengthsT, gradT *tensor.Dense, gradOutT *F64) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *ctcLossDiffOp) f32s(logProbsT, targetsT, inputLengthsT, targetLengthsT, gradT *tensor.Dense, gradOutT *F32) error {
	_ = "STUB: not implemented"
	return nil
}

func (op ctcLossDiffOp) getOrPanic(view tensor.View, coords ...int) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (op ctcLossDiffOp) getOrPanicF64(view tensor.View, coords ...int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (op ctcLossDiffOp) getOrPanicF32(view tensor.View, coords ...int) float32 {
	_ = "STUB: not implemented"
	return 0
}

var (
	_ Op = &ctcLossDiffOp{}
)
