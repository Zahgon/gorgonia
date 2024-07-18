package stdops

import (
	"context"
	"fmt"
	"runtime/trace"

	"github.com/chewxy/hm"
	"gorgonia.org/gorgonia/exprgraph"
	"gorgonia.org/gorgonia/internal"
	"gorgonia.org/gorgonia/internal/errors"
	"gorgonia.org/tensor"
	"gorgonia.org/tensor/dense"

	"gorgonia.org/gorgonia/ops"
	"gorgonia.org/gorgonia/types"
	"gorgonia.org/shapes"
)

// In general all reductions will reduce the dims
// { a → b | (D b = D a - 1) }
func reductionShapeExpr(along shapes.Axes) shapes.Expr {
	var reducts shapes.ReductOf
	if len(along) == 0 {
		reducts.A = shapes.Var('a')
		reducts.Along = shapes.AllAxes
	} else {
		reducts = shapes.Reduce(shapes.Var('a'), along)

	}
	return shapes.Arrow{
		shapes.Var('a'),
		reducts,
	}
}

func reductionTypeExpr(along shapes.Axes) hm.Type {
	a := hm.TypeVariable('a')
	d := types.MakeReduct(a, along)
	return hm.NewFnType(a, d)
}

func denseReduction[DT any](task *trace.Task, ctx context.Context, f func(t *dense.Dense[DT], opts ...tensor.FuncOpt) (*dense.Dense[DT], error), along []int, input *dense.Dense[DT]) (retVal *dense.Dense[DT], err error) {
	defer task.End()
	// TODO: put ctx into input.Engine somehow
	var ret *dense.Dense[DT]
	if ret, err = f(input, tensor.Along(along...)); err != nil {
		return nil, errors.Wrapf(err, "Failed to perform reduction of %v", errors.ThisFn())
	}
	return ret, err
}

type Reduction[DT any] struct {
	op    ops.Op[DT]
	along shapes.Axes
	//def   T
}

// Arity returns the number of inputs the Op expects. -1 indicates that it's n-ary and will be determined at runtime.
func (op *Reduction[DT]) Arity() int { return 1 }

// Type informs the type of the Op (not the node). This will be used by the type system to infer the final type of the node.
func (op *Reduction[DT]) Type() hm.Type { return reductionTypeExpr(op.along) }

// ShapeExpr informs the shape operations that the Op will do. A quick primer is given in the README of the shapes package.
func (op *Reduction[DT]) ShapeExpr() shapes.Expr { return reductionShapeExpr(op.along) }

// Do executes the op.
func (op *Reduction[DT]) Do(ctx context.Context, vs ...tensor.Basic[DT]) (retVal tensor.Basic[DT], err error) {
	panic("not implemented") // TODO: Implement
}

func (op *Reduction[DT]) String() string { return fmt.Sprintf("%v/", op.op) }

// Sum is an op that performs a reduction with + along the given axes
type Sum[DT tensor.Num] struct {
	along shapes.Axes
}

// Arity returns the number of inputs the Op expects. -1 indicates that it's n-ary and will be determined at runtime.
func (op *Sum[DT]) Arity() int { return 1 }

// Type informs the type of the Op (not the node). This will be used by the type system to infer the final type of the node.
func (op *Sum[DT]) Type() hm.Type { return reductionTypeExpr(op.along) }

// ShapeExpr informs the shape operations that the Op will do. A quick primer is given in the README of the shapes package.
func (op *Sum[DT]) ShapeExpr() shapes.Expr { return reductionShapeExpr(op.along) }

// Do executes the op.
func (op *Sum[DT]) Do(ctx context.Context, vs ...tensor.Basic[DT]) (retVal tensor.Basic[DT], err error) {
	if err := internal.HandleCtx(ctx); err != nil {
		return retVal, err
	}

	switch t := any(vs[0]).(type) {
	case *dense.Dense[DT]:
		ctx2, task := trace.NewTask(ctx, op.String())
		retVal, err = denseReduction(task, ctx2, dense.Sum[DT], axesToInts(op.along), t)
		if err != nil {
			return retVal, err
		}
		return
	default:
		return retVal, errors.NYI(t)
	}
}

func (op *Sum[DT]) String() string { return "∑" }

func (op *Sum[DT]) DiffWRT(inputs int) []bool { return onetrue }

func (op *Sum[DT]) SymDiff(g *exprgraph.Graph, inputs []*exprgraph.Node, output *exprgraph.Node, grad *exprgraph.Node) (retVal []*exprgraph.Node, err error) {
	panic("not implemented") // TODO: Implement
}
