//go:build !cuda
// +build !cuda

package gorgonia

import (
	"gorgonia.org/tensor"
)

func finalizeTapeMachine(m *tapeMachine) { _ = "STUB: not implemented"; return }

func UseCudaFor(ops ...string) VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func (m *tapeMachine) getEngine(dev Device) tensor.Engine {
	_ = "STUB: not implemented"
	return *new(tensor.Engine)
}

func (instr *execOp) exec(m *tapeMachine) (err error) { _ = "STUB: not implemented"; return nil }

func (instr deviceTransport) exec(m *tapeMachine) error { _ = "STUB: not implemented"; return nil }
