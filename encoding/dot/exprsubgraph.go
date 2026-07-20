package dot

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	gonumDot "gonum.org/v1/gonum/graph/encoding/dot"
	internalEncoding "gorgonia.org/gorgonia/internal/encoding"
)

type exprSubGraph struct {
	name string
	subs map[internalEncoding.Group]operatorSubGraph
	graph.DirectedBuilder
}

func (g exprSubGraph) DOTID() string { _ = "STUB: not implemented"; return "" }

func (g exprSubGraph) DOTAttributers() (graph, node, edge encoding.Attributer) {
	_ = "STUB: not implemented"
	return *new(encoding.Attributer), *new(encoding.Attributer), *new(encoding.Attributer)
}

func (g exprSubGraph) Structure() []gonumDot.Graph { _ = "STUB: not implemented"; return nil }
