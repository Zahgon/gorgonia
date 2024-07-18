package stdops

import (
	"bytes"
	"context"
	"fmt"
	"runtime/trace"

	"github.com/chewxy/hm"
	"github.com/pkg/errors"
	"gorgonia.org/gorgonia/exprgraph"
	"gorgonia.org/gorgonia/internal"
	"gorgonia.org/gorgonia/internal/datatypes"
	"gorgonia.org/gorgonia/internal/encoding"
	"gorgonia.org/gorgonia/types"
	"gorgonia.org/gorgonia/values/dual"
	"gorgonia.org/shapes"
	"gorgonia.org/tensor"
)

type transposeOp[DT any] struct {
	pattern shapes.Axes
}

func Transpose[DT any](pattern []int) transposeOp[DT] {
	// TODO: pattern checks
	return transposeOp[DT]{ints2axes(pattern)}
}

// Arity returns the number of inputs the Op expects. -1 indicates that it's n-ary and will be determined at runtime.
func (op transposeOp[DT]) Arity() int { return 1 }

// Type returns Tensor-d a → Tensor-d a.
func (op transposeOp[DT]) Type() hm.Type {
	a := hm.TypeVariable('a')
	d := op.pattern.Dims()
	t := types.MakeTensorType(d, a)
	return types.NewFunc(t, t)
}

// ShapeExpr returns { a → T X[b] a | (D X[b] = D a) },
func (op transposeOp[DT]) ShapeExpr() shapes.Expr {
	expr := shapes.Arrow{
		shapes.Var('a'),
		shapes.TransposeOf{
			op.pattern,
			shapes.Var('a'),
		},
	}

	st := shapes.SubjectTo{
		shapes.Eq,
		shapes.UnaryOp{shapes.Dims, op.pattern},
		shapes.UnaryOp{shapes.Dims, shapes.Var('a')},
	}
	return shapes.Compound{Expr: expr, SubjectTo: st}

	return expr
}

// Do executes the op.
func (op transposeOp[DT]) Do(ctx context.Context, vs ...tensor.Basic[DT]) (retVal tensor.Basic[DT], err error) {
	if err := internal.HandleCtx(ctx); err != nil {
		return retVal, err
	}

	a := vs[0]
	t := any(a).(tensor.BasicOperable[DT])
	_, task := trace.NewTask(ctx, op.String())
	pattern := make([]int, op.pattern.Dims())
	copy(pattern, op.pattern.AsInts())
	retVal, err = t.TransposeAsBasic(pattern...)
	task.End()
	return retVal, err
}

// String returns Aᵀ{...} where `...` is the transposition pattern.
func (op transposeOp[DT]) String() string {
	var buf bytes.Buffer
	buf.WriteString("Aᵀ{")
	for i, ax := range op.pattern {
		fmt.Fprintf(&buf, "%d", ax)
		if i < len(op.pattern)-1 {
			buf.WriteString(", ")
		}
	}

	buf.WriteString("}")
	return buf.String()
}

/* DIFFERENTIATION */

// diffAxes computes the backwards pass transposition pattern.
func (op transposeOp[DT]) diffAxes() shapes.Axes {
	newPattern := make(shapes.Axes, len(op.pattern))
	for i, p := range op.pattern {
		newPattern[p] = shapes.Axis(i)
	}
	return newPattern
}

/* transposeOp implements symdiff.Op */

// DiffWRT returns []bool{true}.
func (op transposeOp[DT]) DiffWRT(i int) []bool { return []bool{true} }

// SymDiff performs the symbolic differentiation of `transposeOp`.
func (op transposeOp[DT]) SymDiff(g *exprgraph.Graph, inputs []exprgraph.Node, output, grad exprgraph.Node) (retVal []exprgraph.Node, err error) {
	newPattern := op.diffAxes()
	op2 := transposeOp[DT]{pattern: newPattern}
	retVal = make([]exprgraph.Node, 1)
	if retVal[0], err = apply[DT](g, op2, grN(inputs[0]), grad); err == nil {
		setGroup(g, encoding.GradientCluster, retVal...)
	}
	return
}

/* transposeOp implements ADOp */

// DoDiff allows transposeOp to be automatically differentiated.
func (op transposeOp[DT]) DoDiff(ctx context.Context, inputs []datatypes.Tensor, output datatypes.Tensor) (err error) {
	newPattern := op.diffAxes()

	adv := exprgraph.T2B[DT](inputs[0]).(dual.Value[DT])
	bdv := exprgraph.T2B[DT](output).(dual.Value[DT])
	advd := adv.DVal()
	bdvd := bdv.DVal()

	// set up new context
	ctx2, task := trace.NewTask(ctx, op.String())

	op2 := transposeOp[DT]{pattern: newPattern}

	d, err := op2.Do(ctx2, bdvd)
	if err != nil {
		return errors.Wrap(err, "Failed to perform transposeOp.DoDiff()")
	}
	// get an addOp
	add := Add[DT](adv, d)
	if advd, err = add.PreallocDo(ctx2, advd, advd, d); err != nil {
		return errors.Wrap(err, "Failed to perform addition of gradients in transposeOp.DoDiff()")
	}
	task.End()
	return nil
}
