package gorgonia

import (
	"bytes"
	"log"
)

type lispMachine struct {
	ExternMetadata
	g *ExprGraph
	q []adInstr

	cpumem int64
	gpumem []int64

	sorted Nodes
	df     *dataflow
	fwd    int
	bwd    int

	watchlist Nodes
	logger    *log.Logger
	buf       *bytes.Buffer
	valueFmt  string
	tabcount  int
	logFlags  uint16

	runFlags     uint16
	checkedRoots bool
	evalMode     bool
}

func NewLispMachine(g *ExprGraph, opts ...VMOpt) *lispMachine {
	_ = "STUB: not implemented"
	return nil
}

func (m *lispMachine) logBwd() bool { _ = "STUB: not implemented"; return false }
func (m *lispMachine) doLogBwd()    { _ = "STUB: not implemented"; return }
func (m *lispMachine) dontLogBwd()  { _ = "STUB: not implemented"; return }
func (m *lispMachine) runBwd() bool { _ = "STUB: not implemented"; return false }
func (m *lispMachine) doExecBwd()   { _ = "STUB: not implemented"; return }
func (m *lispMachine) dontExecBwd() { _ = "STUB: not implemented"; return }

func (m *lispMachine) logFwd() bool { _ = "STUB: not implemented"; return false }
func (m *lispMachine) doLogFwd()    { _ = "STUB: not implemented"; return }
func (m *lispMachine) dontLogFwd()  { _ = "STUB: not implemented"; return }
func (m *lispMachine) runFwd() bool { _ = "STUB: not implemented"; return false }
func (m *lispMachine) doExecFwd()   { _ = "STUB: not implemented"; return }
func (m *lispMachine) dontExecFwd() { _ = "STUB: not implemented"; return }

func (m *lispMachine) watchNaN() bool { _ = "STUB: not implemented"; return false }
func (m *lispMachine) doWatchNaN()    { _ = "STUB: not implemented"; return }
func (m *lispMachine) dontWatchNaN()  { _ = "STUB: not implemented"; return }

func (m *lispMachine) watchInf() bool { _ = "STUB: not implemented"; return false }
func (m *lispMachine) doWatchInf()    { _ = "STUB: not implemented"; return }
func (m *lispMachine) dontWatchInf()  { _ = "STUB: not implemented"; return }

func (m *lispMachine) watchAll() bool { _ = "STUB: not implemented"; return false }
func (m *lispMachine) doWatchAll()    { _ = "STUB: not implemented"; return }
func (m *lispMachine) dontWatchAll()  { _ = "STUB: not implemented"; return }

func (m *lispMachine) dealloc() bool { _ = "STUB: not implemented"; return false }
func (m *lispMachine) doDealloc()    { _ = "STUB: not implemented"; return }
func (m *lispMachine) dontDealloc()  { _ = "STUB: not implemented"; return }

func (m *lispMachine) setRootGrad() bool    { _ = "STUB: not implemented"; return false }
func (m *lispMachine) allowSetRootGrad()    { _ = "STUB: not implemented"; return }
func (m *lispMachine) disallowSetRootGrad() { _ = "STUB: not implemented"; return }

func (m *lispMachine) Reset() { _ = "STUB: not implemented"; return }

func (m *lispMachine) Close() error { _ = "STUB: not implemented"; return nil }

func (m *lispMachine) RunAll() (err error) { _ = "STUB: not implemented"; return nil }

func (m *lispMachine) UnbindAll() { _ = "STUB: not implemented"; return }

func (m *lispMachine) LastRun() (n *Node, backprop bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *lispMachine) checkRoots() (err error) { _ = "STUB: not implemented"; return nil }

func (m *lispMachine) prepGraph() (err error) { _ = "STUB: not implemented"; return nil }

func (m *lispMachine) runall(errChan chan error, doneChan chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (m *lispMachine) forward() (err error) { _ = "STUB: not implemented"; return nil }

func (m *lispMachine) backward() (err error) { _ = "STUB: not implemented"; return nil }

func (m *lispMachine) watchedLogf(format string, attrs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (m *lispMachine) logf(format string, attrs ...interface{}) { _ = "STUB: not implemented"; return }

func (m *lispMachine) enterLogScope() { _ = "STUB: not implemented"; return }

func (m *lispMachine) leaveLogScope() { _ = "STUB: not implemented"; return }

type adInstr struct {
	ADOp
	ctx ExecutionContext

	inputs Nodes
	output *Node
}

func (instr adInstr) do() error { _ = "STUB: not implemented"; return nil }
