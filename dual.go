package gorgonia

import (
	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type dualValue struct {
	Value
	d Value
}

func (dv *dualValue) SetDeriv(d Value) error { _ = "STUB: not implemented"; return nil }

func (dv *dualValue) SetValue(v Value) error { _ = "STUB: not implemented"; return nil }

func (dv *dualValue) Clone() (retVal interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dv *dualValue) Type() hm.Type       { _ = "STUB: not implemented"; return *new(hm.Type) }
func (dv *dualValue) Dtype() tensor.Dtype { _ = "STUB: not implemented"; return *new(tensor.Dtype) }

func (dv *dualValue) ValueEq(a Value) bool { _ = "STUB: not implemented"; return false }

func (dv *dualValue) String() string { _ = "STUB: not implemented"; return "" }

func (dv *dualValue) sanity() error { _ = "STUB: not implemented"; return nil }

func (dv *dualValue) clone0() (retVal *dualValue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func constantDV(val Value) *dualValue { _ = "STUB: not implemented"; return nil }

func variableDV(val Value) *dualValue { _ = "STUB: not implemented"; return nil }

func dvUnit(v Value) *dualValue { _ = "STUB: not implemented"; return nil }

func dvUnitVar(v Value) *dualValue { _ = "STUB: not implemented"; return nil }

func dvUnit0(v Value) *dualValue { _ = "STUB: not implemented"; return nil }

func dvUnitManaged(v Value, op *ExternalOp) (*dualValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dvUnitVarManaged(v Value, op *ExternalOp) (*dualValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func idValue(inputs []*dualValue) (retVals []Value) { _ = "STUB: not implemented"; return nil }

func dvBind(op Op, inputs []*dualValue) (retVal *dualValue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dvBindVar(op Op, inputs []*dualValue) (retVal *dualValue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dvBind0(op Op, retVal *dualValue, inputs []*dualValue) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func dvBindVar0(op Op, retVal *dualValue, inputs []*dualValue) (err error) {
	_ = "STUB: not implemented"
	return nil
}
