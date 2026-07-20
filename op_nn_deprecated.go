//go:build ignore
// +build ignore

package gorgonia

import (
	"gorgonia.org/tensor"
)

func (op *BatchNormOp) f64sOld(input, output *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *BatchNormOp) f32sOld(input, output *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *batchnormDiffOp) f64sOld(input, inGrad, outGrad *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *batchnormDiffOp) f32sOld(input, inGrad, outGrad *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *BatchNormOp) mul64(prealloc Value, scale Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *BatchNormOp) add64(prealloc, bias Value) error { _ = "STUB: not implemented"; return nil }

func (op *BatchNormOp) mul32(prealloc Value, scale Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *BatchNormOp) add32(prealloc, bias Value) error { _ = "STUB: not implemented"; return nil }
