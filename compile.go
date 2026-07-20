package gorgonia

import (
	"io"
)

func Compile(g *ExprGraph) (prog *program, locMap map[*Node]register, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CompileFunction(g *ExprGraph, inputs, outputs Nodes) (prog *program, locMap map[*Node]register, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type codegenerator struct {
	locMap     map[*Node]register
	lastWrites map[register]*Node
	flushed    map[int]struct{}
	allocated  map[register]struct{}
	freed      map[register]struct{}
	deferFree  map[register]struct{}
	instrMap   map[*Node]fragment
	queue      []int

	lastReads map[register]int

	cpumem int64
	gpumem []int64

	g              *ExprGraph
	inputs, sorted Nodes
	df             *dataflow
	instructions   fragment
}

func newCodeGenerator(inputs, sorted Nodes, df *dataflow) *codegenerator {
	_ = "STUB: not implemented"
	return nil
}

func (cg *codegenerator) addInstr(node *Node, instr tapeInstr) { _ = "STUB: not implemented"; return }

func (cg *codegenerator) updateLastWrites(reg register, n *Node) { _ = "STUB: not implemented"; return }

func (cg *codegenerator) flush() { _ = "STUB: not implemented"; return }

func (cg *codegenerator) addArg(node *Node, interv *interval) { _ = "STUB: not implemented"; return }

func (cg *codegenerator) addStmt(node *Node, interv *interval, i int) {
	_ = "STUB: not implemented"
	return
}

func (cg *codegenerator) addNode(node, replacement *Node, interv *interval, i int) {
	_ = "STUB: not implemented"
	return
}

func (cg *codegenerator) insertFree(instrID int, node *Node) { _ = "STUB: not implemented"; return }

func (cg *codegenerator) insertLastFrees() int { _ = "STUB: not implemented"; return 0 }

func (cg *codegenerator) gen() (*program, map[*Node]register) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compileState(w io.Writer, g *ExprGraph, df *dataflow) { _ = "STUB: not implemented"; return }
