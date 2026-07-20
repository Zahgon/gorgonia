package cuda

import (
	"gorgonia.org/cu"
)

var _ External = &Engine{}

const (
	memalign    = 32
	scalarAlign = 8
)

func (e *Engine) HasFunc(name string) bool { _ = "STUB: not implemented"; return false }

func (e *Engine) Sync() chan struct{} { _ = "STUB: not implemented"; return nil }

func (e *Engine) Signal() { _ = "STUB: not implemented"; return }

func (e *Engine) Context() *cu.BatchedContext { _ = "STUB: not implemented"; return nil }

func (e *Engine) CUDNNContext() *cudnn.Context { _ = "STUB: not implemented"; return nil }

func (e *Engine) BLASContext() *cublas.Standard { _ = "STUB: not implemented"; return nil }

func (e *Engine) Modules() map[string]cu.Module { _ = "STUB: not implemented"; return nil }

func (e *Engine) Functions() map[string]cu.Function { _ = "STUB: not implemented"; return nil }

func (e *Engine) ElemGridSize(n int) (gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0, 0
}

func (e *Engine) Init(device cu.Device, size int64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) doInit(size int64) (err error) { _ = "STUB: not implemented"; return nil }

func (e *Engine) Close() error { _ = "STUB: not implemented"; return nil }

func (e *Engine) DoWork() error { _ = "STUB: not implemented"; return nil }

func (e *Engine) Run() { _ = "STUB: not implemented"; return }

func (e *Engine) blockThread(n, dev int) (blocks, threads int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func calcBlocks(n, maxThreads int) int { _ = "STUB: not implemented"; return 0 }
