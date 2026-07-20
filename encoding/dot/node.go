package dot

import (
	"gonum.org/v1/gonum/graph/encoding"
	"gorgonia.org/gorgonia"
	internalEncoding "gorgonia.org/gorgonia/internal/encoding"
)

type node struct {
	n *gorgonia.Node
}

func (n *node) ID() int64 { _ = "STUB: not implemented"; return 0 }

func (n *node) DOTID() string { _ = "STUB: not implemented"; return "" }

func (n *node) Attributes() []encoding.Attribute { _ = "STUB: not implemented"; return nil }

func (n *node) Groups() internalEncoding.Groups {
	_ = "STUB: not implemented"
	return *new(internalEncoding.Groups)
}
