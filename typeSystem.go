package gorgonia

import (
	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

func inferType(expr interface{}) (retVal hm.Type, err error) {
	_ = "STUB: not implemented"
	return *new(hm.Type), nil
}

func inferNodeType(op Op, children ...*Node) (retVal hm.Type, err error) {
	_ = "STUB: not implemented"
	return *new(hm.Type), nil
}

func isScalarType(t hm.Type) bool { _ = "STUB: not implemented"; return false }

func dtypeOf(t hm.Type) (retVal tensor.Dtype, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Dtype), nil
}
