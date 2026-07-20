//go:build !cuda
// +build !cuda

package gorgonia

import "gorgonia.org/tensor"

type Device int

const (
	CPU Device = 0
)

func (d Device) String() string { _ = "STUB: not implemented"; return "" }

func (d Device) IsGPU() bool { _ = "STUB: not implemented"; return false }

func (d Device) Alloc(extern External, size int64) (tensor.Memory, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Memory), nil
}

func (d Device) Free(extern External, mem tensor.Memory, sie uint) error {
	_ = "STUB: not implemented"
	return nil
}
