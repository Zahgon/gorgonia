package dot

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gorgonia.org/gorgonia"
	internalEncoding "gorgonia.org/gorgonia/internal/encoding"
)

func subGraphs() map[internalEncoding.Group]subgrapher { _ = "STUB: not implemented"; return nil }

type attributer []encoding.Attribute

func (a attributer) Attributes() []encoding.Attribute { _ = "STUB: not implemented"; return nil }

func sortedKeys(m map[internalEncoding.Group]subgrapher) (retVal internalEncoding.Groups) {
	_ = "STUB: not implemented"
	return *new(internalEncoding.Groups)
}

func generateDotGraph(g *gorgonia.ExprGraph) (graph.Graph, error) {
	_ = "STUB: not implemented"
	return *new(graph.Graph), nil
}
