package gorgonia

import (
	"gonum.org/v1/gonum/graph"
)

func WalkGraph(start *Node) <-chan *Node { _ = "STUB: not implemented"; return nil }

func walkGraph(start *Node, ch chan *Node, walked NodeSet) { _ = "STUB: not implemented"; return }

func Sort(g *ExprGraph) (sorted Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func UnstableSort(g *ExprGraph) (sorted Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func reverseNodes(sorted Nodes) { _ = "STUB: not implemented"; return }

type byID []graph.Node

func (ns byID) Len() int           { _ = "STUB: not implemented"; return 0 }
func (ns byID) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (ns byID) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func reverseLexical(a []graph.Node) { _ = "STUB: not implemented"; return }
