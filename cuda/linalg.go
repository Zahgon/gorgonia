package cuda

import (
	"gorgonia.org/tensor"
)

var (
	_ tensor.MatVecMuler = &Engine{}
	_ tensor.MatMuler    = &Engine{}
	_ tensor.OuterProder = &Engine{}
)

func (e *Engine) checkThreeFloat(a, b, ret tensor.Tensor) (ad, bd, retVal *tensor.Dense, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (e *Engine) MatVecMul(a, b, prealloc tensor.Tensor) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) MatMul(a, b, prealloc tensor.Tensor) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Outer(a, b, prealloc tensor.Tensor) (err error) {
	_ = "STUB: not implemented"
	return nil
}
