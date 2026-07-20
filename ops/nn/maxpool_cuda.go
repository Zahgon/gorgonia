//go:build cuda
// +build cuda

package nnops

import (
	"hash"

	"github.com/chewxy/hm"
	cudnn "gorgonia.org/cu/dnn"
	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

var (
	_ G.Op       = &maxpool{}
	_ G.CUDADoer = &maxpool{}
	_ G.Op       = &maxpoolDiff{}
	_ G.CUDADoer = &maxpoolDiff{}
)

type maxpool struct {
	*cudnn.Pooling

	xDesc *cudnn.TensorDescriptor
	yDesc *cudnn.TensorDescriptor
}

func newMaxPoolOp(x *G.Node, kernel, pad, stride []int) (*maxpool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *maxpool) Arity() int { _ = "STUB: not implemented"; return 0 }

func (p *maxpool) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (p *maxpool) InferShape(inputs ...G.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (p *maxpool) Do(...G.Value) (G.Value, error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}

func (p *maxpool) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (p *maxpool) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (p *maxpool) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (p *maxpool) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (p *maxpool) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (p *maxpool) String() string { _ = "STUB: not implemented"; return "" }

func (p *maxpool) CUDADo(extern G.External, dev G.Device, prealloc G.Value, inputs ...G.Value) (retVal G.Value, err error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}

func (p *maxpool) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (p *maxpool) SymDiff(inputs G.Nodes, output *G.Node, grad *G.Node) (retVal G.Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(G.Nodes), nil
}

func (p *maxpool) DoDiff(ctx G.ExecutionContext, inputs G.Nodes, output *G.Node) error {
	_ = "STUB: not implemented"
	return nil
}

type maxpoolDiff maxpool

func (op *maxpoolDiff) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *maxpoolDiff) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *maxpoolDiff) InferShape(inputs ...G.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *maxpoolDiff) Do(...G.Value) (G.Value, error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}

func (op *maxpoolDiff) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *maxpoolDiff) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *maxpoolDiff) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *maxpoolDiff) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *maxpoolDiff) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *maxpoolDiff) String() string { _ = "STUB: not implemented"; return "" }

func (op *maxpoolDiff) CUDADo(extern G.External, dev G.Device, prealloc G.Value, inputs ...G.Value) (retVal G.Value, err error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}
