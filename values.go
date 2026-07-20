package gorgonia

import (
	"fmt"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type Value interface {
	Shape() tensor.Shape
	Size() int
	Data() interface{}
	Dtype() tensor.Dtype

	tensor.Memory
	fmt.Formatter
}

type Valuer interface {
	Value() Value
}

type Zeroer interface {
	Value
	Zero()
}

type ZeroValuer interface {
	Value
	ZeroValue() Value
}

type Dtyper interface {
	Dtype() tensor.Dtype
}

type Typer interface {
	Type() hm.Type
}

type ValueEqualer interface {
	ValueEq(Value) bool
}

type ValueCloser interface {
	ValueClose(interface{}) bool
}

type Cloner interface {
	Clone() interface{}
}

type CloneErrorer interface {
	Clone() (interface{}, error)
}

type CopierTo interface {
	CopyTo(dest interface{}) error
}

type CopierFrom interface {
	CopyFrom(src interface{}) error
}

func makeValue(t hm.Type, s tensor.Shape) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func makeValueFromMem(t hm.Type, s tensor.Shape, mem tensor.Memory) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func makeScalarFromMem(dt tensor.Dtype, mem tensor.Memory) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func logicalSize(s tensor.Shape) int { _ = "STUB: not implemented"; return 0 }

func calcMemSize(dt tensor.Dtype, s tensor.Shape) int64 { _ = "STUB: not implemented"; return 0 }

func ScalarAsTensor(v Value, dims int, e tensor.Engine) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}
