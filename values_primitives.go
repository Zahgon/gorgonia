package gorgonia

import (
	"fmt"
	"unsafe"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

type Scalar interface {
	Value
	isScalar() bool
}

type F64 float64

type F32 float32

type I int

type I64 int64

type I32 int32

type U8 byte

type B bool

func NewF64(v float64) *F64 { _ = "STUB: not implemented"; return nil }
func NewF32(v float32) *F32 { _ = "STUB: not implemented"; return nil }
func NewI(v int) *I         { _ = "STUB: not implemented"; return nil }
func NewI64(v int64) *I64   { _ = "STUB: not implemented"; return nil }
func NewI32(v int32) *I32   { _ = "STUB: not implemented"; return nil }
func NewU8(v byte) *U8      { _ = "STUB: not implemented"; return nil }
func NewB(v bool) *B        { _ = "STUB: not implemented"; return nil }

func (v *F64) Shape() tensor.Shape { _ = "STUB: not implemented"; return *new(tensor.Shape) }

func (v *F32) Shape() tensor.Shape { _ = "STUB: not implemented"; return *new(tensor.Shape) }

func (v *I) Shape() tensor.Shape { _ = "STUB: not implemented"; return *new(tensor.Shape) }

func (v *I64) Shape() tensor.Shape { _ = "STUB: not implemented"; return *new(tensor.Shape) }

func (v *I32) Shape() tensor.Shape { _ = "STUB: not implemented"; return *new(tensor.Shape) }

func (v *U8) Shape() tensor.Shape { _ = "STUB: not implemented"; return *new(tensor.Shape) }

func (v *B) Shape() tensor.Shape { _ = "STUB: not implemented"; return *new(tensor.Shape) }

func (v *F64) Size() int { _ = "STUB: not implemented"; return 0 }

func (v *F32) Size() int { _ = "STUB: not implemented"; return 0 }

func (v *I) Size() int { _ = "STUB: not implemented"; return 0 }

func (v *I64) Size() int { _ = "STUB: not implemented"; return 0 }

func (v *I32) Size() int { _ = "STUB: not implemented"; return 0 }

func (v *U8) Size() int { _ = "STUB: not implemented"; return 0 }

func (v *B) Size() int { _ = "STUB: not implemented"; return 0 }

func (v *F64) Data() interface{} { _ = "STUB: not implemented"; return nil }

func (v *F32) Data() interface{} { _ = "STUB: not implemented"; return nil }

func (v *I) Data() interface{} { _ = "STUB: not implemented"; return nil }

func (v *I64) Data() interface{} { _ = "STUB: not implemented"; return nil }

func (v *I32) Data() interface{} { _ = "STUB: not implemented"; return nil }

func (v *U8) Data() interface{} { _ = "STUB: not implemented"; return nil }

func (v *B) Data() interface{} { _ = "STUB: not implemented"; return nil }

func (v *F64) any() float64 { _ = "STUB: not implemented"; return 0 }
func (v *F32) any() float32 { _ = "STUB: not implemented"; return 0 }
func (v *I) any() int       { _ = "STUB: not implemented"; return 0 }
func (v *I64) any() int64   { _ = "STUB: not implemented"; return 0 }
func (v *I32) any() int32   { _ = "STUB: not implemented"; return 0 }
func (v *U8) any() byte     { _ = "STUB: not implemented"; return 0 }
func (v *B) any() bool      { _ = "STUB: not implemented"; return false }

func (v *F64) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (v *F32) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (v *I) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (v *I64) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (v *I32) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (v *U8) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (v *B) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (v *F64) Dtype() tensor.Dtype { _ = "STUB: not implemented"; return *new(tensor.Dtype) }

func (v *F32) Dtype() tensor.Dtype { _ = "STUB: not implemented"; return *new(tensor.Dtype) }

func (v *I) Dtype() tensor.Dtype { _ = "STUB: not implemented"; return *new(tensor.Dtype) }

func (v *I64) Dtype() tensor.Dtype { _ = "STUB: not implemented"; return *new(tensor.Dtype) }

func (v *I32) Dtype() tensor.Dtype { _ = "STUB: not implemented"; return *new(tensor.Dtype) }

func (v *U8) Dtype() tensor.Dtype { _ = "STUB: not implemented"; return *new(tensor.Dtype) }

func (v *B) Dtype() tensor.Dtype { _ = "STUB: not implemented"; return *new(tensor.Dtype) }

func (v *F64) isScalar() bool { _ = "STUB: not implemented"; return false }
func (v *F32) isScalar() bool { _ = "STUB: not implemented"; return false }
func (v *I) isScalar() bool   { _ = "STUB: not implemented"; return false }
func (v *I64) isScalar() bool { _ = "STUB: not implemented"; return false }
func (v *I32) isScalar() bool { _ = "STUB: not implemented"; return false }
func (v *U8) isScalar() bool  { _ = "STUB: not implemented"; return false }
func (v *B) isScalar() bool   { _ = "STUB: not implemented"; return false }

func (v *F64) Uintptr() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *F32) Uintptr() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *I) Uintptr() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *I64) Uintptr() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *I32) Uintptr() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *U8) Uintptr() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *B) Uintptr() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *F64) MemSize() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *F32) MemSize() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *I) MemSize() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *I64) MemSize() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *I32) MemSize() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *U8) MemSize() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *B) MemSize() uintptr { _ = "STUB: not implemented"; return 0 }

func (v *F64) Pointer() unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func (v *F32) Pointer() unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func (v *I) Pointer() unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func (v *I64) Pointer() unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func (v *I32) Pointer() unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func (v *U8) Pointer() unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func (v *B) Pointer() unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func formatScalar(v Scalar, s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func anyToScalar(any interface{}) (Scalar, tensor.Dtype) {
	_ = "STUB: not implemented"
	return *new(Scalar), *new(tensor.Dtype)
}

func anyToValue(any interface{}) (val Value, t hm.Type, dt tensor.Dtype, err error) {
	_ = "STUB: not implemented"
	return *new(Value), *new(hm.Type), *new(tensor.Dtype), nil
}

func one(dt tensor.Dtype) Scalar { _ = "STUB: not implemented"; return *new(Scalar) }

func zero(dt tensor.Dtype) Scalar { _ = "STUB: not implemented"; return *new(Scalar) }
