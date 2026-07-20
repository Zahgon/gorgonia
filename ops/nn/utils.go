package nnops

import (
	"unsafe"

	"github.com/chewxy/hm"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func simpleHash(op gorgonia.Op) uint32 { _ = "STUB: not implemented"; return 0 }

func checkArity(op gorgonia.Op, inputs int) error { _ = "STUB: not implemented"; return nil }

type nomem struct{}

func (nomem) Uintptr() uintptr           { _ = "STUB: not implemented"; return 0 }
func (nomem) Pointer() unsafe.Pointer    { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }
func (nomem) IsNativelyAccessible() bool { _ = "STUB: not implemented"; return false }

func calcMemSize(dt tensor.Dtype, s tensor.Shape) uintptr { _ = "STUB: not implemented"; return 0 }

func dtypeOf(t hm.Type) (retVal tensor.Dtype, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Dtype), nil
}

func CheckConvolutionParams(pad, stride, dilation []int) error {
	_ = "STUB: not implemented"
	return nil
}
