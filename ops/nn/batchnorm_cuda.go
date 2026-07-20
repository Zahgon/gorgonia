//go:build cuda
// +build cuda

package nnops

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type BatchNormOp struct {
	mode              cudnn.BatchNormMode
	momentum, epsilon float64

	xDesc     *cudnn.TensorDescriptor
	bnScratch *cudnn.TensorDescriptor

	training bool
}

func newBatchNormOp(momentum, epsilon float64) *BatchNormOp { _ = "STUB: not implemented"; return nil }

func (op *BatchNormOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *BatchNormOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *BatchNormOp) InferShape(inputs ...gorgonia.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *BatchNormOp) Do(...gorgonia.Value) (gorgonia.Value, error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}
func (op *BatchNormOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op *BatchNormOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op *BatchNormOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op *BatchNormOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *BatchNormOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }
func (op *BatchNormOp) String() string   { _ = "STUB: not implemented"; return "" }

func (op *BatchNormOp) CUDADo(extern gorgonia.External, dev gorgonia.Device, prealloc gorgonia.Value, inputs ...gorgonia.Value) (retVal gorgonia.Value, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}

func (op *BatchNormOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op *BatchNormOp) SymDiff(inputs gorgonia.Nodes, output *gorgonia.Node, grad *gorgonia.Node) (retVal gorgonia.Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Nodes), nil
}

func (op *BatchNormOp) DoDiff(ctx gorgonia.ExecutionContext, inputs gorgonia.Nodes, output *gorgonia.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *BatchNormOp) SetTraining() { _ = "STUB: not implemented"; return }
func (op *BatchNormOp) SetTesting()  { _ = "STUB: not implemented"; return }
func (op *BatchNormOp) Reset() error { _ = "STUB: not implemented"; return nil }

type batchNormDiffOp struct {
	*BatchNormOp
}

func (op *batchNormDiffOp) Do(...gorgonia.Value) (gorgonia.Value, error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}

func (op *batchNormDiffOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *batchNormDiffOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *batchNormDiffOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *batchNormDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *batchNormDiffOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *batchNormDiffOp) CUDADo(extern gorgonia.External, dev gorgonia.Device, prealloc gorgonia.Value, inputs ...gorgonia.Value) (retVal gorgonia.Value, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}
