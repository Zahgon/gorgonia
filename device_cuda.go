//go:build cuda
// +build cuda

package gorgonia

import (
	"gorgonia.org/cu"
	"gorgonia.org/tensor"
)

type Device cu.Device

const CPU = Device(cu.CPU)

func (d Device) String() string { _ = "STUB: not implemented"; return "" }

func (d Device) Alloc(extern External, size int64) (tensor.Memory, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Memory), nil
}

func (d Device) Free(extern External, mem tensor.Memory, size int64) (err error) {
	_ = "STUB: not implemented"
	return nil
}
