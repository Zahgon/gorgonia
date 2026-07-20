package dot

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	gonumDot "gonum.org/v1/gonum/graph/encoding/dot"
	internalEncoding "gorgonia.org/gorgonia/internal/encoding"
)

type constantSubGraph struct {
	name string
	graph.DirectedBuilder
	subs map[internalEncoding.Group]operatorSubGraph
}

func (g constantSubGraph) DOTID() string { _ = "STUB: not implemented"; return "" }

func (g constantSubGraph) DOTAttributers() (graph, node, edge encoding.Attributer) {
	_ = "STUB: not implemented"
	return *new(encoding.Attributer), *new(encoding.Attributer), *new(encoding.Attributer)
}

func (g constantSubGraph) Structure() []gonumDot.Graph { _ = "STUB: not implemented"; return nil }
