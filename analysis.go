package gorgonia

type dataflow struct {
	uniques map[uint32]*Node

	replacements map[*Node]*Node
	intervals    map[*Node]*interval

	devTransChildren map[*Node]Nodes
	devTransRepl     map[*Node]*Node
}

func newdataflow() *dataflow { _ = "STUB: not implemented"; return nil }

func (df *dataflow) vn(n *Node) (retVal *Node, unique bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (df *dataflow) analyzeDevice(n *Node) { _ = "STUB: not implemented"; return }

func (df *dataflow) replaceWithSelf(sorted Nodes) { _ = "STUB: not implemented"; return }

func (df *dataflow) fixIntervalDevices(sorted Nodes) { _ = "STUB: not implemented"; return }

func analyze(g *ExprGraph, sorted Nodes) *dataflow { _ = "STUB: not implemented"; return nil }

func newDevTransNode(read, write *Node, from, to Device) *Node {
	_ = "STUB: not implemented"
	return nil
}

func (df *dataflow) insertDeviceInstr(sorted Nodes) Nodes {
	_ = "STUB: not implemented"
	return *new(Nodes)
}

func (df *dataflow) buildIntervals(sorted Nodes) { _ = "STUB: not implemented"; return }
