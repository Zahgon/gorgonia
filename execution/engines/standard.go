package engines

import (
	"gorgonia.org/gorgonia/exprgraph"
	"gorgonia.org/gorgonia/internal/datatypes"
	"gorgonia.org/gorgonia/ops"
	"gorgonia.org/tensor"
)

// Hybrid is an engine that contains a *exprgraph.Graph
type Hybrid interface {
	tensor.Engine
	Workhorse() tensor.Engine

	Graph() *exprgraph.Graph
}

// H is an engine that creates symbolic nodes and also performs the operations immediately.
// However when it encounters a Symbolic node, the remaining operations are symbolic only.
type H[DT tensor.Num] struct {
	StandardEngine[DT]
	g *exprgraph.Graph
}

func (e *H[DT]) Graph() *exprgraph.Graph { return e.g }

func (e *H[DT]) SetGraph(g *exprgraph.Graph) { e.g = g }

// StandardEngine is a set of operations that must be supported by an engine in order to be used by Gorgonia.
type StandardEngine[DT any] interface {
	tensor.Engine
	tensor.FuncOptHandler[DT]
	tensor.BLA[DT]
	tensor.Adder[DT]
}

type Queueer[DT any] interface {
	Q(op ops.Op[DT], inputs []datatypes.Tensor, output datatypes.Tensor) error
}
