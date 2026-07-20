package cuda

import "C"

import (
	"sync"

	"gorgonia.org/cu"
	"gorgonia.org/tensor"
)

var (
	_ tensor.Adder = &Engine{}
	_ tensor.Suber = &Engine{}
	_ tensor.Muler = &Engine{}
	_ tensor.Diver = &Engine{}
	_ tensor.Power = &Engine{}
	_ tensor.Moder = &Engine{}

	_ tensor.MatMuler    = &Engine{}
	_ tensor.MatVecMuler = &Engine{}
	_ tensor.OuterProder = &Engine{}

	_ tensor.Lter   = &Engine{}
	_ tensor.Lteer  = &Engine{}
	_ tensor.Gter   = &Engine{}
	_ tensor.Gteer  = &Engine{}
	_ tensor.ElEqer = &Engine{}
)

type Engine struct {
	tensor.Engine
	sync.Mutex

	a bfc
	b cublas.Standard
	c cu.BatchedContext
	d cu.Device
	f map[string]cu.Function
	m map[string]cu.Module
	n cudnn.Context

	warp int
	mtpb int
	mgdx int
	mgdy int
	mgdz int
	mbdx int
	mbdy int
	mbdz int

	freeMem  int64
	totalMem int64

	syncChan      chan struct{}
	finishChan    chan struct{}
	finishChan2   chan struct{}
	workAvailable chan bool
	err           error
	initialized   bool
	running       bool
}

func (e *Engine) AllocAccessible() bool { _ = "STUB: not implemented"; return false }

func (e *Engine) Alloc(size int64) (tensor.Memory, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Memory), nil
}

func (e *Engine) AllocFlags() (tensor.MemoryFlag, tensor.DataOrder) {
	_ = "STUB: not implemented"
	return *new(tensor.MemoryFlag), *new(tensor.DataOrder)
}

func (e *Engine) Free(mem tensor.Memory, size int64) error { _ = "STUB: not implemented"; return nil }

func (e *Engine) Memset(mem tensor.Memory, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Memclr(mem tensor.Memory) { _ = "STUB: not implemented"; return }

func (e *Engine) Memcpy(dst tensor.Memory, src tensor.Memory) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) memcpy(dst cu.DevicePtr, src cu.DevicePtr, size int64) {
	_ = "STUB: not implemented"
	return
}

func (e *Engine) Accessible(mem tensor.Memory) (tensor.Memory, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Memory), nil
}

func (e *Engine) WorksWith(order tensor.DataOrder) bool { _ = "STUB: not implemented"; return false }

func (e *Engine) NonStdAlloc() { _ = "STUB: not implemented"; return }

func (e *Engine) Errors() error { _ = "STUB: not implemented"; return nil }

func (e *Engine) HasNaN(a tensor.Tensor) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Engine) HasInf(a tensor.Tensor) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
