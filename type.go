package gorgonia

import (
	"fmt"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

var (
	Float64 = tensor.Float64

	Float32 = tensor.Float32

	Int = tensor.Int

	Int64 = tensor.Int64

	Int32 = tensor.Int32

	Byte = tensor.Uint8

	Bool = tensor.Bool

	Ptr = tensor.UnsafePointer

	vecF64  = &TensorType{Dims: 1, Of: tensor.Float64}
	vecF32  = &TensorType{Dims: 1, Of: tensor.Float32}
	matF64  = &TensorType{Dims: 2, Of: tensor.Float64}
	matF32  = &TensorType{Dims: 2, Of: tensor.Float32}
	ten3F64 = &TensorType{Dims: 3, Of: tensor.Float64}
	ten3F32 = &TensorType{Dims: 3, Of: tensor.Float32}

	f64T = tensor.Float64
	f32T = tensor.Float32
)

var acceptableDtypes = [...]tensor.Dtype{tensor.Float64, tensor.Float32, tensor.Int, tensor.Int64, tensor.Int32, tensor.Byte, tensor.Bool}

type TensorType struct {
	Dims int

	Of hm.Type
}

func makeFromTensorType(t TensorType, tv hm.TypeVariable) TensorType {
	_ = "STUB: not implemented"
	return *new(TensorType)
}

func makeTensorType(dims int, typ hm.Type) TensorType {
	_ = "STUB: not implemented"
	return *new(TensorType)
}

func newTensorType(dims int, typ hm.Type) *TensorType { _ = "STUB: not implemented"; return nil }

func (t TensorType) Name() string { _ = "STUB: not implemented"; return "" }

func (t TensorType) Format(state fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (t TensorType) String() string { _ = "STUB: not implemented"; return "" }

func (t TensorType) Types() hm.Types { _ = "STUB: not implemented"; return *new(hm.Types) }

func (t TensorType) Normalize(k, v hm.TypeVarSet) (hm.Type, error) {
	_ = "STUB: not implemented"
	return *new(hm.Type), nil
}

func (t TensorType) Apply(sub hm.Subs) hm.Substitutable {
	_ = "STUB: not implemented"
	return *new(hm.Substitutable)
}

func (t TensorType) FreeTypeVar() hm.TypeVarSet {
	_ = "STUB: not implemented"
	return *new(hm.TypeVarSet)
}

func (t TensorType) Eq(other hm.Type) bool { _ = "STUB: not implemented"; return false }
