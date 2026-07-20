package dot

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	gonumDot "gonum.org/v1/gonum/graph/encoding/dot"
	internalEncoding "gorgonia.org/gorgonia/internal/encoding"
)

type operatorSubGraph struct {
	name string
	id   int
	graph.DirectedBuilder
	subs map[internalEncoding.Group]operatorSubGraph
}

func (g operatorSubGraph) DOTID() string { _ = "STUB: not implemented"; return "" }

func (g operatorSubGraph) DOTAttributers() (graph, node, edge encoding.Attributer) {
	_ = "STUB: not implemented"
	return *new(encoding.Attributer), *new(encoding.Attributer), *new(encoding.Attributer)
}

func (g operatorSubGraph) Structure() []gonumDot.Graph { _ = "STUB: not implemented"; return nil }
