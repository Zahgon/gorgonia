package gorgonia

import (
	"bytes"
	"fmt"
	"log"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type tapeMachine struct {
	ExternMetadata

	p      *program
	locMap map[*Node]register

	cpumem []Value
	gpumem []Value

	pc int

	bindNodesDV  Nodes
	watchNodes   Nodes
	watchRegs    []register
	watchNodeIDs []NodeID
	logger       *log.Logger
	buf          *bytes.Buffer
	valueFmt     string
	tabcount     int
	logFlags     uint16
	closureQueue []func() error

	runFlags uint16
	evalMode bool
}

func NewTapeMachine(g *ExprGraph, opts ...VMOpt) *tapeMachine {
	_ = "STUB: not implemented"
	return nil
}

func (m *tapeMachine) logBwd() bool { _ = "STUB: not implemented"; return false }
func (m *tapeMachine) doLogBwd()    { _ = "STUB: not implemented"; return }
func (m *tapeMachine) dontLogBwd()  { _ = "STUB: not implemented"; return }

func (m *tapeMachine) logFwd() bool { _ = "STUB: not implemented"; return false }
func (m *tapeMachine) doLogFwd()    { _ = "STUB: not implemented"; return }
func (m *tapeMachine) dontLogFwd()  { _ = "STUB: not implemented"; return }

func (m *tapeMachine) watchNaN() bool { _ = "STUB: not implemented"; return false }
func (m *tapeMachine) doWatchNaN()    { _ = "STUB: not implemented"; return }
func (m *tapeMachine) dontWatchNaN()  { _ = "STUB: not implemented"; return }

func (m *tapeMachine) watchInf() bool { _ = "STUB: not implemented"; return false }
func (m *tapeMachine) doWatchInf()    { _ = "STUB: not implemented"; return }
func (m *tapeMachine) dontWatchInf()  { _ = "STUB: not implemented"; return }

func (m *tapeMachine) watchPointer() bool { _ = "STUB: not implemented"; return false }
func (m *tapeMachine) doWatchPointer()    { _ = "STUB: not implemented"; return }
func (m *tapeMachine) dontWatchPointer()  { _ = "STUB: not implemented"; return }

func (m *tapeMachine) watchAll() bool { _ = "STUB: not implemented"; return false }
func (m *tapeMachine) doWatchAll()    { _ = "STUB: not implemented"; return }
func (m *tapeMachine) dontWatchAll()  { _ = "STUB: not implemented"; return }

func (m *tapeMachine) alloc() bool { _ = "STUB: not implemented"; return false }
func (m *tapeMachine) doAlloc()    { _ = "STUB: not implemented"; return }
func (m *tapeMachine) dontAlloc()  { _ = "STUB: not implemented"; return }

func (m *tapeMachine) trace() bool { _ = "STUB: not implemented"; return false }
func (m *tapeMachine) doTrace()    { _ = "STUB: not implemented"; return }
func (m *tapeMachine) dontTrace()  { _ = "STUB: not implemented"; return }

func (m *tapeMachine) bindDV() bool { _ = "STUB: not implemented"; return false }
func (m *tapeMachine) doBindDV()    { _ = "STUB: not implemented"; return }
func (m *tapeMachine) dontBindDV()  { _ = "STUB: not implemented"; return }

func (m *tapeMachine) Reset() { _ = "STUB: not implemented"; return }

func (m *tapeMachine) Close() error { _ = "STUB: not implemented"; return nil }

func (m *tapeMachine) Prog() *program { _ = "STUB: not implemented"; return nil }

func (m *tapeMachine) LocMap() map[*Node]register { _ = "STUB: not implemented"; return nil }

func (m *tapeMachine) Let(n *Node, be interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *tapeMachine) Set(a, b *Node) (err error) { _ = "STUB: not implemented"; return nil }

func (m *tapeMachine) Run(frag fragment) (err error) { _ = "STUB: not implemented"; return nil }

func (m *tapeMachine) RunAll() (err error) { _ = "STUB: not implemented"; return nil }

func (m *tapeMachine) runall(errChan chan error, doneChan chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (m *tapeMachine) nodeFromInstr(instr tapeInstr) *Node { _ = "STUB: not implemented"; return nil }

func (m *tapeMachine) getValue(r register) Value { _ = "STUB: not implemented"; return *new(Value) }

func (m *tapeMachine) writeValue(r register, v Value) { _ = "STUB: not implemented"; return }

func (m *tapeMachine) watchedLogf(format string, attrs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (m *tapeMachine) logf(format string, attrs ...interface{}) { _ = "STUB: not implemented"; return }

func (m *tapeMachine) enterLogScope() { _ = "STUB: not implemented"; return }

func (m *tapeMachine) leaveLogScope() { _ = "STUB: not implemented"; return }

type program struct {
	instructions fragment
	args         int
	cpulocs      int
	gpulocs      int
	cpumem       int64
	gpumem       []int64
	g            *ExprGraph
	df           *dataflow
	m            map[*Node]fragment
	r            map[int64]*Node
	sorted       Nodes
}

func (p *program) String() string { _ = "STUB: not implemented"; return "" }

func (p *program) Graph() *ExprGraph { _ = "STUB: not implemented"; return nil }

func (p *program) CPUMemReq() int64 { _ = "STUB: not implemented"; return 0 }

func (p *program) GPUMemReq() []int64 { _ = "STUB: not implemented"; return nil }

type register struct {
	id     int
	device Device
}

func (r register) String() string { _ = "STUB: not implemented"; return "" }

type tapeInstr interface {
	ID() int64
	reads() []register
	writes() register
	exec(*tapeMachine) error
	fmt.Stringer
}

type fragment []tapeInstr

func (f fragment) String() string { _ = "STUB: not implemented"; return "" }

func (f fragment) has(want tapeInstr) bool { _ = "STUB: not implemented"; return false }

type alloc struct {
	id int64
	t  hm.Type
	s  tensor.Shape

	readFrom []register
	writeTo  register
}

func newAlloc(n *Node, writeTo register) alloc { _ = "STUB: not implemented"; return *new(alloc) }

func (instr alloc) ID() int64         { _ = "STUB: not implemented"; return 0 }
func (instr alloc) reads() []register { _ = "STUB: not implemented"; return nil }
func (instr alloc) writes() register  { _ = "STUB: not implemented"; return *new(register) }

func (instr alloc) exec(m *tapeMachine) (err error) { _ = "STUB: not implemented"; return nil }

func (instr alloc) String() string { _ = "STUB: not implemented"; return "" }

type free struct {
	readsFrom register
}

func (instr free) ID() int64                 { _ = "STUB: not implemented"; return 0 }
func (instr free) reads() []register         { _ = "STUB: not implemented"; return nil }
func (instr free) writes() register          { _ = "STUB: not implemented"; return *new(register) }
func (instr free) exec(m *tapeMachine) error { _ = "STUB: not implemented"; return nil }

func (instr free) String() string { _ = "STUB: not implemented"; return "" }

type loadArg struct {
	index   int64
	writeTo register
	name    string
}

func (instr loadArg) ID() int64         { _ = "STUB: not implemented"; return 0 }
func (instr loadArg) reads() []register { _ = "STUB: not implemented"; return nil }
func (instr loadArg) writes() register  { _ = "STUB: not implemented"; return *new(register) }

func (instr loadArg) exec(m *tapeMachine) error { _ = "STUB: not implemented"; return nil }

func (instr loadArg) String() string { _ = "STUB: not implemented"; return "" }

type execOp struct {
	op Op

	id int64

	readFrom []register
	writeTo  register
	size     int64

	preAllocated bool
	useUnsafe    bool
	useGPU       bool
}

func (instr *execOp) ID() int64         { _ = "STUB: not implemented"; return 0 }
func (instr *execOp) reads() []register { _ = "STUB: not implemented"; return nil }
func (instr *execOp) writes() register  { _ = "STUB: not implemented"; return *new(register) }

func newExecOp(n *Node) *execOp { _ = "STUB: not implemented"; return nil }

func (instr *execOp) String() string { _ = "STUB: not implemented"; return "" }

type flushInstr struct{}

func (instr flushInstr) exec(m *tapeMachine) error { _ = "STUB: not implemented"; return nil }

func (instr flushInstr) ID() int64         { _ = "STUB: not implemented"; return 0 }
func (instr flushInstr) reads() []register { _ = "STUB: not implemented"; return nil }
func (instr flushInstr) writes() register  { _ = "STUB: not implemented"; return *new(register) }
func (instr flushInstr) String() string    { _ = "STUB: not implemented"; return "" }

type letInstr struct {
	readFrom register
	writeTo  register
}

func (instr letInstr) ID() int64               { _ = "STUB: not implemented"; return 0 }
func (instr letInstr) reads() []register       { _ = "STUB: not implemented"; return nil }
func (instr letInstr) writes() register        { _ = "STUB: not implemented"; return *new(register) }
func (instr letInstr) exec(*tapeMachine) error { _ = "STUB: not implemented"; return nil }

func (instr letInstr) String() string { _ = "STUB: not implemented"; return "" }

type readInstr struct {
	readFrom register
	into     *Value

	t hm.Type
	s tensor.Shape
}

func (instr *readInstr) ID() int64                       { _ = "STUB: not implemented"; return 0 }
func (instr *readInstr) reads() []register               { _ = "STUB: not implemented"; return nil }
func (instr *readInstr) writes() register                { _ = "STUB: not implemented"; return *new(register) }
func (instr *readInstr) exec(m *tapeMachine) (err error) { _ = "STUB: not implemented"; return nil }

func (instr *readInstr) String() string { _ = "STUB: not implemented"; return "" }

type deviceTransport struct {
	from, to register
}

func (instr deviceTransport) ID() int64         { _ = "STUB: not implemented"; return 0 }
func (instr deviceTransport) reads() []register { _ = "STUB: not implemented"; return nil }

func (instr deviceTransport) writes() register { _ = "STUB: not implemented"; return *new(register) }

func (instr deviceTransport) String() string { _ = "STUB: not implemented"; return "" }
