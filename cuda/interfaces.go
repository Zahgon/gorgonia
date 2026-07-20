package cuda

import (
	"gorgonia.org/cu"
	"gorgonia.org/tensor"
)

type Arena interface {
	Get(size int64) (tensor.Memory, error)

	Put(mem tensor.Memory, size int64)

	ResetAllocator()
}

type External interface {
	Arena

	tensor.Engine

	HasFunc(string) bool

	Sync() chan struct{}

	Signal()

	Context() *cu.BatchedContext

	CUDNNContext() *cudnn.Context

	BLASContext() *cublas.Standard

	Modules() map[string]cu.Module

	Functions() map[string]cu.Function

	ElemGridSize(n int) (gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ int)

	Init(device cu.Device, size int64) error

	Close() error

	DoWork() error
}
