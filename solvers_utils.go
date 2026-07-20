package gorgonia

import (
	"gorgonia.org/tensor"
)

func doL1Reg(w, g tensor.Tensor, l1reg interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func doL2Reg(w, g tensor.Tensor, l2reg interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func computeRecip(x float64, as tensor.Dtype) (retVal interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func divBatch(g tensor.Tensor, batch float64) (err error) { _ = "STUB: not implemented"; return nil }

func clipGrad(g tensor.Tensor, clip, negClip interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}
