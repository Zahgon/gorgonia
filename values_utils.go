package gorgonia

import (
	"unsafe"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

func TypeOf(v Value) hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func typeCheckTypeOf(v Value) hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func ValueEq(a, b Value) bool { _ = "STUB: not implemented"; return false }

func ValueClose(a, b Value) bool { _ = "STUB: not implemented"; return false }

func CloneValue(v Value) (Value, error) { _ = "STUB: not implemented"; return *new(Value), nil }

func ZeroValue(v Value) Value { _ = "STUB: not implemented"; return *new(Value) }

func Copy(dest, src Value) (Value, error) { _ = "STUB: not implemented"; return *new(Value), nil }

func setEngine(v Value, e tensor.Engine) { _ = "STUB: not implemented"; return }

func valueToPointer(v Value) unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }
