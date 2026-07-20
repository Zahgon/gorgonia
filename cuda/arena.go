package cuda

import (
	"gorgonia.org/tensor"
)

var _ Arena = &Engine{}

func (e *Engine) Get(size int64) (tensor.Memory, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Memory), nil
}

func (e *Engine) Put(mem tensor.Memory, size int64) { _ = "STUB: not implemented"; return }

func (e *Engine) ResetAllocator() { _ = "STUB: not implemented"; return }
