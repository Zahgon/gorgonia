package gorgonia

import (
	"gorgonia.org/gorgonia/execution/engines"
	"gorgonia.org/gorgonia/exprgraph"

	"gorgonia.org/gorgonia/internal/errors"
	gtu "gorgonia.org/gorgonia/internal/tensorutils"
	"gorgonia.org/gorgonia/ops"
	stdops "gorgonia.org/gorgonia/ops/std"

	"gorgonia.org/tensor"
)

func binopSymbolic[DT any](op ops.Desc, eng engines.Hybrid, a, b Tensor) (retVal Tensor, err error) {
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
	return
}

func Add[DT any](a, b Tensor) (retVal Tensor, err error) {
	// We check to see if the engines are Hybrid engine (i.e. they have graphs)
	eng, ok := a.Engine().(engines.Hybrid)
	if !ok {
		eng, ok = b.Engine().(engines.Hybrid)
	}
	ctx := gtu.CtxFromEngines(a.Engine(), b.Engine())

	op := stdops.Add[DT](a, b)
	if ok {
		// do symbolic stuff
		if retVal, err = binopSymbolic[DT](op, eng, a, b); err != nil {
			return nil, err
		}
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
		ct = any(ct).(tensor.BasicAliker[DT]).AlikeAsBasic(tensor.WithEngine(a.Engine()), tensor.WithShape(shp...))
	default:
		// one of a or b is not a value tensor
		return retVal, nil
	}
	if ct, err = op.PreallocDo(ctx, ct, at, bt); err != nil {
		return nil, err
	}
	if retVal == nil {
		retVal = ct // return not the Node, but the value.
	}

	// check if engine is backwards (i.e. requires a queue)
	// if not, return.
	var q engines.Queueer[DT]
	q, ok = a.Engine().Workhorse().(engines.Queueer[DT])
	if !ok {
		q, ok = b.Engine().Workhorse().(engines.Queueer[DT])
	}
	if ok && q != nil {
		// do queue stuff here
		err = q.Q(op, []Tensor{a, b}, retVal)
	}

	return
}
