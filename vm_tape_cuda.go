//go:build cuda
// +build cuda

package gorgonia

import (
	"gorgonia.org/tensor"
)

func finalizeTapeMachine(m *tapeMachine) { _ = "STUB: not implemented"; return }

func (m *tapeMachine) init() {
	var initCUDA bool
	cudaLogf("instructions %v", len(m.p.instructions))
	for _, instr := range m.p.instructions {
		if eo, ok := instr.(*execOp); ok {
			if _, ok := eo.op.(CUDADoer); ok {
				initCUDA = true
				break
			}
		}
	}

	if !initCUDA {
		cudaLogf("No CUDA ops")
		return
	}

	if err := m.ExternMetadata.init(m.p.gpumem); err != nil {
		m.ExternMetadata.initFail()
		panic(err)
	}
	m.loadStdLib()

}

func (m *tapeMachine) loadStdLib() { _ = "STUB: not implemented"; return }

func (m *tapeMachine) getEngine(dev Device) tensor.Engine {
	_ = "STUB: not implemented"
	return *new(tensor.Engine)
}

func (instr *execOp) exec(m *tapeMachine) (err error) { _ = "STUB: not implemented"; return nil }

func (instr deviceTransport) exec(m *tapeMachine) (err error) {
	_ = "STUB: not implemented"
	return nil
}
