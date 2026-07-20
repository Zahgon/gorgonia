package dot

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	gonumDot "gonum.org/v1/gonum/graph/encoding/dot"
)

type dotGraph struct {
	graph.Directed
	subs []gonumDot.Graph
}

func (g dotGraph) DOTAttributers() (graph, node, edge encoding.Attributer) {
	_ = "STUB: not implemented"
	return *new(encoding.Attributer), *new(encoding.Attributer), *new(encoding.Attributer)
}

func (g dotGraph) Structure() []gonumDot.Graph { _ = "STUB: not implemented"; return nil }
