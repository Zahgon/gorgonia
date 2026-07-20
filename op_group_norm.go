package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

const (
	groupNormChunkSize = 16
	groupNormVecSize   = 8
)

func GroupNorm(x, scale, bias *Node, numGroups, numChannels int, epsilon float64) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GroupNormOp struct {
	numGroups, numChannels int
	epsilon                float64

	mean, rstd *tensor.Dense
}

func (op *GroupNormOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *GroupNormOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *GroupNormOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *GroupNormOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *GroupNormOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *GroupNormOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *GroupNormOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *GroupNormOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *GroupNormOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *GroupNormOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *GroupNormOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *GroupNormOp) f64s(xT, prealloc *tensor.Dense, batchSize, channels, imageSize int) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *GroupNormOp) rowwiseMomentsF64(x []float64, n int, ddof int) (mean float64, variance float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (op *GroupNormOp) addMomentsF64(m0add int, m1add, m2add float64, m0 *int, m1, m2 *float64) {
	_ = "STUB: not implemented"
	return
}

func (op *GroupNormOp) addMomentsVecF64(m0add int, m1add, m2add []float64, m0 *int, m1, m2 []float64) {
	_ = "STUB: not implemented"
	return
}

func (op *GroupNormOp) ceilLog2F64(x int) int { _ = "STUB: not implemented"; return 0 }

func (op *GroupNormOp) f32s(xT, prealloc *tensor.Dense, batchSize, channels, imageSize int) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *GroupNormOp) rowwiseMomentsF32(x []float32, n int, ddof int) (mean float32, variance float32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (op *GroupNormOp) addMomentsF32(m0add int, m1add, m2add float32, m0 *int, m1, m2 *float32) {
	_ = "STUB: not implemented"
	return
}

func (op *GroupNormOp) addMomentsVecF32(m0add int, m1add, m2add []float32, m0 *int, m1, m2 []float32) {
	_ = "STUB: not implemented"
	return
}

func (op *GroupNormOp) ceilLog2F32(x int) int { _ = "STUB: not implemented"; return 0 }

func (op *GroupNormOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *GroupNormOp) SymDiff(inputs Nodes, output, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *GroupNormOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

type groupNormDiffOp struct {
	*GroupNormOp
}

func (op *groupNormDiffOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *groupNormDiffOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *groupNormDiffOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *groupNormDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *groupNormDiffOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *groupNormDiffOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *groupNormDiffOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *groupNormDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *groupNormDiffOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *groupNormDiffOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *groupNormDiffOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *groupNormDiffOp) f64s(input, prealloc, outGrad *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *groupNormDiffOp) computeInternalGradientsF64(batchSize, channels, imageSize int, input, dyT *tensor.Dense) ([]float64, []float64) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *groupNormDiffOp) f32s(input, prealloc, outGrad *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *groupNormDiffOp) computeInternalGradientsF32(batchSize, channels, imageSize int, input, dyT *tensor.Dense) ([]float32, []float32) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *groupNormDiffOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *groupNormDiffOp) SymDiff(inputs Nodes, output, grad *Node) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *groupNormDiffOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

var (
	_ Op   = &GroupNormOp{}
	_ ADOp = &GroupNormOp{}
	_ SDOp = &GroupNormOp{}

	_ Op = &groupNormDiffOp{}
)
