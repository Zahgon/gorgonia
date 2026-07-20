package cuda

import (
	"gorgonia.org/tensor"
)

func getDenseTensor(t tensor.Tensor) (tensor.DenseTensor, error) {
	_ = "STUB: not implemented"
	return *new(tensor.DenseTensor), nil
}

func handleFuncOpts(expShape tensor.Shape, expType tensor.Dtype, o tensor.DataOrder, strict bool, opts ...tensor.FuncOpt) (reuse tensor.DenseTensor, safe, toReuse, incr, same bool, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.DenseTensor), false, false, false, false, nil
}

func binaryCheck(a, b tensor.Tensor) (err error) { _ = "STUB: not implemented"; return nil }

func unaryCheck(a tensor.Tensor) error { _ = "STUB: not implemented"; return nil }

func logicalSize(s tensor.Shape) int { _ = "STUB: not implemented"; return 0 }

func constructName2(a, b tensor.Tensor, fn string) (name string) {
	_ = "STUB: not implemented"
	return ""
}

func constructName1(a tensor.Tensor, leftTensor bool, fn string) (name string) {
	_ = "STUB: not implemented"
	return ""
}
