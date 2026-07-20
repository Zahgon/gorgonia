package gorgonia

import (
	"hash"
	"unsafe"

	"gorgonia.org/tensor"
)

type Tensor interface {
	Shape() tensor.Shape
	Strides() []int
	Dtype() tensor.Dtype
	Dims() int
	Size() int
	DataSize() int

	IsScalar() bool
	ScalarValue() interface{}

	Engine() tensor.Engine
	MemSize() uintptr
	Uintptr() uintptr
	Pointer() unsafe.Pointer
	IsNativelyAccessible() bool
	IsManuallyManaged() bool
}

type hashWriter interface {
	WriteHash(hash.Hash)
}

type arityer interface {
	Arity() int
}
