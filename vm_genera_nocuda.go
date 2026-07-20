//go:build !cuda
// +build !cuda

package gorgonia

func (m *lispMachine) init() error {
	if err := m.prepGraph(); err != nil {
		return err
	}
	return nil
}

func (m *lispMachine) execDevTrans(op devTrans, n *Node, children Nodes) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func finalizeLispMachine(m *lispMachine) { _ = "STUB: not implemented"; return }

func (m *lispMachine) ForceCPU() { _ = "STUB: not implemented"; return }
