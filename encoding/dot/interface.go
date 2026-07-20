package dot

import gonumDot "gonum.org/v1/gonum/graph/encoding/dot"

type subgrapher interface {
	gonumDot.Graph
	gonumDot.Attributers
	gonumDot.Structurer
}
