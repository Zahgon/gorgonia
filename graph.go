package gorgonia

import (
	"gonum.org/v1/gonum/graph"
)

type ExprGraph struct {
	name string

	all Nodes

	byID   map[int64]int
	byHash map[uint32]*Node
	evac   map[uint32]Nodes
	to     map[*Node]Nodes

	leaves    Nodes
	constants Nodes
	roots     Nodes
	counter   uint
}

type graphconopt func(g *ExprGraph)

func WithGraphName(name string) graphconopt { _ = "STUB: not implemented"; return *new(graphconopt) }

func NewGraph(opts ...graphconopt) *ExprGraph { _ = "STUB: not implemented"; return nil }

func (g *ExprGraph) Clone() interface{} { _ = "STUB: not implemented"; return nil }

func (g *ExprGraph) AddNode(n *Node) (retVal *Node) { _ = "STUB: not implemented"; return nil }

func (g *ExprGraph) addToAll(n *Node) { _ = "STUB: not implemented"; return }

func (g *ExprGraph) RemoveNode(node graph.Node) { _ = "STUB: not implemented"; return }

func (g *ExprGraph) SetEdge(e graph.Edge) { _ = "STUB: not implemented"; return }

func (g *ExprGraph) Roots() (retVal Nodes) { _ = "STUB: not implemented"; return *new(Nodes) }

func (g *ExprGraph) Inputs() (retVal Nodes) { _ = "STUB: not implemented"; return *new(Nodes) }

func (g *ExprGraph) UnbindAll() { _ = "STUB: not implemented"; return }

func (g *ExprGraph) UnbindAllNonInputs() { _ = "STUB: not implemented"; return }

func (g *ExprGraph) ByName(name string) (retVal Nodes) {
	_ = "STUB: not implemented"
	return *new(Nodes)
}

func (g *ExprGraph) Constant(v Value) *Node { _ = "STUB: not implemented"; return nil }

func (g *ExprGraph) String() string { _ = "STUB: not implemented"; return "" }

func (g *ExprGraph) ToDot() string { _ = "STUB: not implemented"; return "" }

func (g *ExprGraph) Edges() graph.Edges { _ = "STUB: not implemented"; return *new(graph.Edges) }

func (g *ExprGraph) removeAllEdgesFrom(n *Node) { _ = "STUB: not implemented"; return }

func (g *ExprGraph) Node(id int64) graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (g *ExprGraph) node(id int64) *Node { _ = "STUB: not implemented"; return nil }

func (g *ExprGraph) Has(nodeid int64) bool { _ = "STUB: not implemented"; return false }

func (g *ExprGraph) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *ExprGraph) AllNodes() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (g *ExprGraph) From(nodeid int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (g *ExprGraph) HasEdgeBetween(x, y int64) bool { _ = "STUB: not implemented"; return false }

func (g *ExprGraph) Edge(u, v int64) graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }

func (g *ExprGraph) HasEdgeFromTo(u, v int64) bool { _ = "STUB: not implemented"; return false }

func (g *ExprGraph) To(nid int64) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (g *ExprGraph) subgraph(ns Nodes, findMissing bool, opts ...Nodes) *ExprGraph {
	_ = "STUB: not implemented"
	return nil
}

func (g *ExprGraph) Subgraph(ns ...*Node) *ExprGraph { _ = "STUB: not implemented"; return nil }

func (g *ExprGraph) SubgraphRoots(ns ...*Node) *ExprGraph { _ = "STUB: not implemented"; return nil }

func (g *ExprGraph) ExactSubgraphRoots(ns ...*Node) *ExprGraph {
	_ = "STUB: not implemented"
	return nil
}

func (g *ExprGraph) walkFromRoots(ns ...*Node) Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

type edge struct {
	from, to graph.Node
	weight   float64
}

func (e edge) From() graph.Node         { _ = "STUB: not implemented"; return *new(graph.Node) }
func (e edge) To() graph.Node           { _ = "STUB: not implemented"; return *new(graph.Node) }
func (e edge) ReversedEdge() graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }
func (e edge) Weight() float64          { _ = "STUB: not implemented"; return 0 }
