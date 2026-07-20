//go:build cuda
// +build cuda

package nnops

import (
	"hash"
	"unsafe"

	"github.com/chewxy/hm"
	"gorgonia.org/cu"
	cudnn "gorgonia.org/cu/dnn"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type dropout struct {
	*cudnn.Dropout
	seed  uint64
	xDesc *cudnn.TensorDescriptor
}

func newDropout(x *gorgonia.Node, prob float64) (*dropout, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *dropout) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *dropout) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *dropout) InferShape(inputs ...gorgonia.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *dropout) Do(...gorgonia.Value) (gorgonia.Value, error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}
func (op *dropout) ReturnsPtr() bool          { _ = "STUB: not implemented"; return false }
func (op *dropout) CallsExtern() bool         { _ = "STUB: not implemented"; return false }
func (op *dropout) OverwritesInput() int      { _ = "STUB: not implemented"; return 0 }
func (op *dropout) WriteHash(h hash.Hash)     { _ = "STUB: not implemented"; return }
func (op *dropout) Hashcode() uint32          { _ = "STUB: not implemented"; return 0 }
func (op *dropout) String() string            { _ = "STUB: not implemented"; return "" }
func (op *dropout) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op *dropout) SymDiff(inputs gorgonia.Nodes, output *gorgonia.Node, grad *gorgonia.Node) (retVal gorgonia.Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Nodes), nil
}

func (op *dropout) DoDiff(ctx gorgonia.ExecutionContext, inputs gorgonia.Nodes, output *gorgonia.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *dropout) CUDADo(extern gorgonia.External, dev gorgonia.Device, prealloc gorgonia.Value, inputs ...gorgonia.Value) (retVal gorgonia.Value, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}

type dropoutDiff struct {
	*dropout
}

func (op *dropoutDiff) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *dropoutDiff) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *dropoutDiff) InferShape(inputs ...gorgonia.DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *dropoutDiff) Do(...gorgonia.Value) (gorgonia.Value, error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}

func (op *dropoutDiff) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *dropoutDiff) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *dropoutDiff) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *dropoutDiff) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *dropoutDiff) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *dropoutDiff) String() string { _ = "STUB: not implemented"; return "" }

func (op *dropoutDiff) CUDADo(extern gorgonia.External, dev gorgonia.Device, prealloc gorgonia.Value, inputs ...gorgonia.Value) (retVal gorgonia.Value, err error) {
	_ = "STUB: not implemented"
	return *new(gorgonia.Value), nil
}

type tmpWrapper cu.DevicePtr

func (p tmpWrapper) Uintptr() uintptr { _ = "STUB: not implemented"; return 0 }

func (p tmpWrapper) Pointer() unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func (p tmpWrapper) IsNativelyAccessible() bool { _ = "STUB: not implemented"; return false }
