package mnist

import (
	"gorgonia.org/tensor"
)

type RawImage []byte

type Label uint8

func Load(typ, loc string, as tensor.Dtype) (inputs, targets tensor.Tensor, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), *new(tensor.Tensor), nil
}

func pixelWeight(px byte) float64 { _ = "STUB: not implemented"; return 0 }

func reversePixelWeight(px float64) byte { _ = "STUB: not implemented"; return 0 }

func prepareX(M []RawImage, dt tensor.Dtype) (retVal tensor.Tensor) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor)
}

func prepareY(N []Label, dt tensor.Dtype) (retVal tensor.Tensor) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor)
}
