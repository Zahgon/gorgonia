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
	_ G.Op       = &convolution{}
	_ G.CUDADoer = &convolution{}
)

type convolution struct {
	*cudnn.Convolution

	padding, stride, dilation []int
	inShape, filterShape      tensor.Shape

	xDesc, yDesc *cudnn.TensorDescriptor
	wDesc        *cudnn.Filter
}

func makeConvolutionOp(im, filter *G.Node, kernelShape tensor.Shape, pad, stride, dilation []int) (retVal *convolution, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *convolution) Arity() int { _ = "STUB: not implemented"; return 0 }

func (c *convolution) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (c *convolution) InferShape(inputs ...G.DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (c *convolution) Do(inputs ...G.Value) (retVal G.Value, err error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}

func (c *convolution) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (c *convolution) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (c *convolution) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (c *convolution) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (c *convolution) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (c *convolution) String() string { _ = "STUB: not implemented"; return "" }

func (c *convolution) CUDADo(extern G.External, dev G.Device, prealloc G.Value, inputs ...G.Value) (retVal G.Value, err error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}

func (c *convolution) DoDiff(ctx G.ExecutionContext, inputs G.Nodes, output *G.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *convolution) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (c *convolution) SymDiff(inputs G.Nodes, output *G.Node, grad *G.Node) (retVal G.Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(G.Nodes), nil
}

type convDiffIm struct {
	*convolution
	outputDesc *cudnn.TensorDescriptor
}

func (c *convDiffIm) Arity() int { _ = "STUB: not implemented"; return 0 }

func (c *convDiffIm) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (c *convDiffIm) InferShape(shps ...G.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (c *convDiffIm) Do(...G.Value) (G.Value, error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}

func (c *convDiffIm) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (c *convDiffIm) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (c *convDiffIm) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (c *convDiffIm) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (c *convDiffIm) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (c *convDiffIm) String() string { _ = "STUB: not implemented"; return "" }

func (c *convDiffIm) CUDADo(extern G.External, dev G.Device, prealloc G.Value, inputs ...G.Value) (retVal G.Value, err error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}

type convDiffFilter struct {
	*convolution
	outputDesc *cudnn.TensorDescriptor
}

func (c *convDiffFilter) Arity() int { _ = "STUB: not implemented"; return 0 }

func (c *convDiffFilter) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (c *convDiffFilter) InferShape(...G.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (c *convDiffFilter) Do(...G.Value) (G.Value, error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}

func (c *convDiffFilter) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (c *convDiffFilter) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (c *convDiffFilter) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (c *convDiffFilter) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (c *convDiffFilter) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (c *convDiffFilter) String() string { _ = "STUB: not implemented"; return "" }

func (c *convDiffFilter) CUDADo(extern G.External, dev G.Device, prealloc G.Value, inputs ...G.Value) (retVal G.Value, err error) {
	_ = "STUB: not implemented"
	return *new(G.Value), nil
}
