//go:build cuda
// +build cuda

package gorgonia

import (
	"log"
)

func (m *lispMachine) init() error {
	if err := m.prepGraph(); err != nil {
		return err
	}

	df := newdataflow()
	df.replaceWithSelf(m.sorted)
	df.buildIntervals(m.sorted)
	df.fixIntervalDevices(m.sorted)
	m.df = df

	if err := m.calcMemSize(); err != nil {
		log.Printf("err1")
		return err
	}

	if len(m.gpumem) == 0 {
		m.ForceCPU()
		return nil
	}

	if err := m.ExternMetadata.init(m.gpumem); err != nil {
		m.ExternMetadata.initFail()
		return err
	}
	m.loadStdLib()

	if len(m.engines) == 0 {
		m.ForceCPU()
	}
	return nil
}

func finalizeLispMachine(m *lispMachine) { _ = "STUB: not implemented"; return }

func (m *lispMachine) WorkAvailable() <-chan bool { _ = "STUB: not implemented"; return nil }

func (m *lispMachine) calcMemSize() (err error) { _ = "STUB: not implemented"; return nil }

func (m *lispMachine) execDevTrans(op devTrans, n *Node, children Nodes) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *lispMachine) loadStdLib() { _ = "STUB: not implemented"; return }

func (m *lispMachine) ForceCPU() { _ = "STUB: not implemented"; return }
