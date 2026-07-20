package gorgonia

import (
	"gorgonia.org/tensor"
)

type StandardEngine struct {
	tensor.StdEng
}

func (e StandardEngine) Transpose(a tensor.Tensor, expStrides []int) error {
	_ = "STUB: not implemented"
	return nil
}
