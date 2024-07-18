package exprgraph_test

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/chewxy/hm"
	"github.com/pkg/errors"
	"gorgonia.org/gorgonia"
	"gorgonia.org/gorgonia/exprgraph"
	"gorgonia.org/gorgonia/ops"
	"gorgonia.org/gorgonia/values/dual"
	"gorgonia.org/shapes"
	"gorgonia.org/tensor"
	"gorgonia.org/tensor/dense"
)

var (
	_ ops.Op[float64] = matmul[float64]{}
)

type NoOp struct{}

func (NoOp) NoOp()         {}
func (NoOp) Error() string { return "NoOp" }

type GraphEngine interface {
	tensor.Engine
	Workhorse() tensor.Engine
	Graph() *exprgraph.Graph
}

type StandardEngine[DT any] interface {
	tensor.Engine
	tensor.FuncOptHandler[DT]
	tensor.BLA[DT]
	tensor.Adder[DT]
}

type ADOp[DT any] interface {
	ops.Op[DT]
	DoDiff(ctx context.Context, inputs []tensor.Basic[DT], output tensor.Basic[DT]) error
}

type Queueer[DT any] interface {
	Q(op ops.Op[DT], inputs []gorgonia.Tensor, output gorgonia.Tensor) error
}

type matmuler[T any] interface {
	MatMul(u T, opts ...tensor.FuncOpt) (T, error)
}

// matmul is an Op
type matmul[DT tensor.Num] struct{}

// Arity returns the number of inputs the Op expects. -1 indicates that it's n-ary and will be determined at runtime.
func (op matmul[DT]) Arity() int { return 2 }

// Type informs the type of the Op (not the node). This will be used by the type system to infer the final type of the node.
func (op matmul[DT]) Type() hm.Type {
	return hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('a'), hm.TypeVariable('a'))
}

// ShapeExpr informs the shape operations that the Op will do. A quick primer is given in the README of the shapes package.
func (op matmul[DT]) ShapeExpr() shapes.Expr {
	a := shapes.Var('a')
	b := shapes.Var('b')
	c := shapes.Var('c')
	return shapes.MakeArrow(
		shapes.Abstract{a, b},
		shapes.Abstract{b, c},
		shapes.Abstract{a, c},
	)
}

// Do executes the op.
func (op matmul[DT]) Do(ctx context.Context, vs ...tensor.Basic[DT]) (retVal tensor.Basic[DT], err error) {
	if ctx != nil {
		select {
		case <-ctx.Done():
			return retVal, NoOp{}
		default:
		}

	}
	a := vs[0]
	b := vs[1].(*dense.Dense[DT])
	mm, ok := any(a).(matmuler[*dense.Dense[DT]]) // because :)
	if !ok {
		return retVal, errors.Errorf("expected %T to have a MatMul method", a)
	}
	return mm.MatMul(b)
}

func (op matmul[DT]) String() string { return "×" }

func (op matmul[DT]) PreallocDo(ctx context.Context, prealloc tensor.Basic[DT], vs ...tensor.Basic[DT]) (retVal tensor.Basic[DT], err error) {
	if ctx != nil {
		select {
		case <-ctx.Done():
			return retVal, NoOp{}
		default:
		}

	}
	a := vs[0]
	b := vs[1]
	switch mm := any(a).(type) {
	case matmuler[*dense.Dense[DT]]:
		return mm.MatMul(b.(*dense.Dense[DT]), tensor.WithReuse(prealloc))
	default:
		var ret tensor.Basic[DT]
		if ret, err = tensor.MatMul[DT](a, b, tensor.WithReuse(prealloc)); err != nil {
			return retVal, err
		}
		return ret, nil
	}
}

func (op matmul[DT]) DoDiff(ctx context.Context, inputs []tensor.Basic[DT], output tensor.Basic[DT]) (err error) {
	adv := exprgraph.T2B[DT](inputs[0]).(dual.Value[DT])
	bdv := exprgraph.T2B[DT](inputs[1]).(dual.Value[DT])
	cdv := exprgraph.T2B[DT](output).(dual.Value[DT])

	advd := adv.DVal()
	bdvd := bdv.DVal()

	// temporary transpose
	var bdvT, advT tensor.Basic[DT]
	if bdvT, err = bdv.Val().(tensor.BasicOperable[DT]).TAsBasic(); err != nil {
		return err
	}
	if advT, err = adv.Val().(tensor.BasicOperable[DT]).TAsBasic(); err != nil {
		return err
	}

	// dA = C×B'
	if _, err := op.PreallocDo(ctx, advd, cdv.Val(), bdvT); err != nil {
		return err
	}

	// dB = A'×C
	if _, err := op.PreallocDo(ctx, bdvd, advT, cdv.Val()); err != nil {
		return err
	}
	return nil
}

func MatMul[DT tensor.Num](a, b gorgonia.Tensor) (retVal gorgonia.Tensor, err error) {
	eng, ok := a.Engine().(GraphEngine)
	if !ok {
		eng, ok = b.Engine().(GraphEngine)
	}

	op := matmul[DT]{}
	if ok {
		// do symbolic stuff
		g := eng.Graph()

		var aname, bname string
		var anode, bnode exprgraph.Node
		if aname, err = g.NameOf(a); err != nil {
			// create a node
			aname = randomName(a)
			anode = exprgraph.New[DT](g, aname, tensor.WithBacking(a))
		}

		if bname, err = g.NameOf(b); err != nil {
			// create b node
			bname = randomName(b)
			bnode = exprgraph.New[DT](g, bname, tensor.WithBacking(b))
		}

		cname := aname + op.String() + bname
		err = nil

		// construct node
		if anode == nil {
			if anode = g.NodeOf(a); anode == nil {
				return nil, errors.Errorf("MatMul: Cannot find Node a of %v", a)
			}
		}
		if bnode == nil {
			if bnode = g.NodeOf(b); bnode == nil {
				return nil, errors.Errorf("MatMul: Cannot find Node b of %v", b)
			}
		}

		// shape checks are done here
		cnode, err := exprgraph.Apply[DT](g, op, cname, anode, bnode)
		if err != nil {
			return nil, err
		}
		retVal = cnode
	}

	// check if engine supports MatMul. If not, return
	_, aok := a.Engine().Workhorse().(tensor.BLA[DT])
	_, bok := b.Engine().Workhorse().(tensor.BLA[DT])
	switch {
	case !aok && !bok:
		_, aok = a.Engine().Workhorse().(tensor.BLA[DT])
		_, bok = b.Engine().Workhorse().(tensor.BLA[DT])
		if !aok && !bok {
			return
		}
	default:

	}
	// do the values stuff
	at := exprgraph.T2B[DT](a)
	bt := exprgraph.T2B[DT](b)
	aok, bok = at != nil, bt != nil

	var ct tensor.Basic[DT]
	switch {
	case aok && bok && retVal != nil:
		// both a and b  are values, so we can "materialize" c
		rv := exprgraph.SymToVal[DT, tensor.Basic[DT]](retVal.(*exprgraph.Symbolic[DT])) // turn a Symbolic into a Value
		retVal = rv
		ct = rv.Value()
	case aok && bok && retVal == nil:
		// we'd have to create one ourselves
		shp := tensor.Shape{a.Shape()[0], b.Shape()[1]}
		ct = at.AlikeAsBasic(tensor.WithEngine(a.Engine()), tensor.WithShape(shp...))
	default:
		// one of a or b is not a value tensor
		log.Printf("One of a or b is not a value tensor a %T b %T", a, b)
		return retVal, nil
	}

	if ct, err = op.PreallocDo(context.Background(), ct, at, bt); err != nil {
		return nil, err
	}
	if retVal == nil {
		retVal = ct // return not the Node, but the value.
	}

	// check if engine is backwards (i.e. requires a queue)
	// if not, return.
	var q Queueer[DT]
	q, ok = a.Engine().Workhorse().(Queueer[DT])
	if !ok {
		q, ok = b.Engine().Workhorse().(Queueer[DT])
	}
	if q != nil {
		// do queue stuff here
		err = q.Q(op, []gorgonia.Tensor{a, b}, retVal)
	}
	return
}

type adder[DT any] interface {
	Add(tensor.Basic[DT], ...tensor.FuncOpt) (tensor.Basic[DT], error)
	AddScalar(s DT, scalarOnLeft bool, opts ...tensor.FuncOpt) (tensor.Basic[DT], error)
}

// add is addition with a scalar on the right
type add[DT tensor.Num] struct{}

// Arity returns the number of inputs the Op expects. -1 indicates that it's n-ary and will be determined at runtime.
func (op add[DT]) Arity() int { return 2 }

// Type informs the type of the Op (not the node). This will be used by the type system to infer the final type of the node.
func (op add[DT]) Type() hm.Type {
	return hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('b'), hm.TypeVariable('a'))
}

// ShapeExpr informs the shape operations that the Op will do. A quick primer is given in the README of the shapes package.
func (op add[DT]) ShapeExpr() shapes.Expr {
	a := shapes.Var('a')
	return shapes.MakeArrow(a, shapes.ScalarShape(), a)
}

// Do executes the op.
func (op add[DT]) Do(ctx context.Context, vs ...tensor.Basic[DT]) (retVal tensor.Basic[DT], err error) {
	a := vs[0]
	b := vs[1]
	mm, ok := any(a).(adder[DT])
	if !ok {
		return retVal, errors.Errorf("expected %T to have a Add method", a)
	}
	return mm.Add(b)
}

func (op add[DT]) String() string { return "+" }

func (op add[DT]) PreallocDo(ctx context.Context, prealloc tensor.Basic[DT], vs ...tensor.Basic[DT]) (retVal tensor.Basic[DT], err error) {
	if ctx != nil {
		select {
		case <-ctx.Done():
			return retVal, NoOp{}
		default:
		}

	}

	a := vs[0]
	b := vs[1]
	switch mm := any(a).(type) {
	case adder[DT]:
		return mm.AddScalar(b.Data()[0], true, tensor.WithReuse(prealloc))
	default:
		var ret tensor.Basic[DT]

		if ret, err = tensor.Add[DT](a, b, tensor.WithReuse(prealloc)); err != nil {
			return retVal, err
		}
		return ret, nil
	}

}

func (op add[DT]) DoDiff(ctx context.Context, inputs []gorgonia.Tensor, output gorgonia.Tensor) error {
	adv := exprgraph.T2B[DT](inputs[0]).(dual.Value[DT])
	bdv := exprgraph.T2B[DT](inputs[1]).(dual.Value[DT])

	advd := adv.DVal()
	bdvd := bdv.DVal()
	// this should be replaced with a kernel call somewhere
	data := advd.Data()
	for i := range data {
		data[i] += 1
	}

	data = bdvd.Data()
	for i := range data {
		data[i] += 1
	}

	return nil
}

func Add[DT tensor.Num](a, b gorgonia.Tensor) (retVal gorgonia.Tensor, err error) {
	eng, ok := a.Engine().(GraphEngine)
	if !ok {
		eng, ok = b.Engine().(GraphEngine)
	}

	op := add[DT]{}
	if ok {
		// do symbolic stuff
		g := eng.Graph()

		var aname, bname string
		var anode, bnode exprgraph.Node

		if aname, err = g.NameOf(a); err != nil {
			// create a node
			aname = randomName(a)
			anode = exprgraph.New[DT](g, aname, tensor.WithBacking(a))
		}
		if bname, err = g.NameOf(b); err != nil {
			// create b node
			bname = randomName(b)
			bnode = exprgraph.New[DT](g, bname, tensor.WithBacking(b))
		}
		cname := aname + op.String() + bname
		err = nil

		// construct node
		if anode == nil {
			if anode = g.NodeOf(a); anode == nil {
				return nil, errors.Errorf("Add: Cannot find Node a of %v", a)
			}
		}
		if bnode == nil {
			if bnode = g.NodeOf(b); bnode == nil {
				return nil, errors.Errorf("Add: Cannot find Node b of %v", b)
			}
		}

		// shape checks are done here
		cnode, err := exprgraph.Apply[DT](g, op, cname, anode, bnode)
		if err != nil {
			return nil, err
		}
		retVal = cnode
	}

	// check if engine supports Add. If not, return
	_, aok := a.Engine().Workhorse().(tensor.Adder[DT])
	_, bok := b.Engine().Workhorse().(tensor.Adder[DT])
	switch {
	case !aok && !bok:
		_, aok = a.Engine().Workhorse().(tensor.Adder[DT])
		_, bok = b.Engine().Workhorse().(tensor.Adder[DT])
		if !aok && !bok {
			return
		}
	default:
	}
	// do the values stuff'
	at := exprgraph.T2B[DT](a)
	bt := exprgraph.T2B[DT](b)
	aok, bok = at != nil, bt != nil

	var ct tensor.Basic[DT]
	switch {
	case aok && bok && retVal != nil:
		// both a and b  are values, so we can "materialize" c
		rv := exprgraph.SymToVal[DT, tensor.Basic[DT]](retVal.(*exprgraph.Symbolic[DT])) // turn a Symbolic into a Value
		retVal = rv
		ct = rv.Value()
	case aok && bok && retVal == nil:
		// we'd have to create one ourselves
		// NOTICE: This example assumes that `Add` adds a matrix to a scalar.
		shp := a.Shape()
		ct = at.AlikeAsBasic(tensor.WithEngine(a.Engine()), tensor.WithShape(shp...))
	default:
		// one of a or b is not a value tensor
		log.Printf("One of a or b is not a value tensor a %T b %T", a, b)
		return retVal, nil
	}
	if ct, err = op.PreallocDo(nil, ct, at, bt); err != nil {
		return nil, err
	}
	if retVal == nil {
		retVal = ct // return not the Node, but the value.
	}

	// check if engine is backwards (i.e. requires a queue)
	// if not, return.
	var q Queueer[DT]
	q, ok = a.Engine().Workhorse().(Queueer[DT])
	if !ok {
		q, ok = b.Engine().Workhorse().(Queueer[DT])
	}
	if q != nil {
		// do queue stuff here
		err = q.Q(op, []gorgonia.Tensor{a, b}, retVal)
	}

	return
}

// ExampleOperations is a placeholder to display documentation of
// the MatMul and Add functions used in the other examples.
func Example_operations() {
	// See other examples for usage
}

var rndCounter int
var rndLock sync.Mutex

func randomName(a gorgonia.Tensor) string {
	rndLock.Lock()
	defer rndLock.Unlock()
	rndCounter++
	return fmt.Sprintf("Random_%d", rndCounter)
}

func resetRnd() {
	rndLock.Lock()
	rndCounter = 0
	rndLock.Unlock()
}

// getDeriv is a utility function
func getDeriv[DT tensor.Num, T tensor.Tensor[DT, T]](t gorgonia.Tensor) T {
	switch n := t.(type) {
	case *exprgraph.Value[DT, T]:
		return n.Basic.(*dual.Dual[DT, T]).Deriv()
	case *exprgraph.Value[DT, *dual.Dual[DT, T]]:
		return n.Basic.(*dual.Dual[DT, T]).Deriv()
	case *exprgraph.Value[DT, tensor.Basic[DT]]:
		return n.DV().(T)
	}
	panic(fmt.Sprintf("NYI %T", t))

}
