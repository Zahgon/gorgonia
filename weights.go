package gorgonia

import (
	"gorgonia.org/tensor"
)

type InitWFn func(dt tensor.Dtype, s ...int) interface{}

func Zeroes() InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func Ones() InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func RangedFrom(start int) InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func RangedFromWithStep(start, increment interface{}) InitWFn {
	_ = "STUB: not implemented"
	return *new(InitWFn)
}

func ValuesOf(val interface{}) InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func Gaussian(mean, stdev float64) InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func Uniform(low, high float64) InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func GlorotN(gain float64) InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func GlorotU(gain float64) InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func HeN(gain float64) InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func HeU(gain float64) InitWFn { _ = "STUB: not implemented"; return *new(InitWFn) }

func Gaussian64(mean, stdev float64, s ...int) []float64 { _ = "STUB: not implemented"; return nil }

func Gaussian32(mean, stdev float64, s ...int) []float32 { _ = "STUB: not implemented"; return nil }

func Uniform64(low, high float64, s ...int) []float64 { _ = "STUB: not implemented"; return nil }

func Uniform32(low, high float64, s ...int) []float32 { _ = "STUB: not implemented"; return nil }

func Binomial64(trials, prob float64, s ...int) []float64 { _ = "STUB: not implemented"; return nil }

func Binomial32(trials, prob float64, s ...int) []float32 { _ = "STUB: not implemented"; return nil }

func GlorotEtAlN64(gain float64, s ...int) []float64 { _ = "STUB: not implemented"; return nil }

func GlorotEtAlN32(gain float64, s ...int) []float32 { _ = "STUB: not implemented"; return nil }

func GlorotEtAlU64(gain float64, s ...int) []float64 { _ = "STUB: not implemented"; return nil }

func GlorotEtAlU32(gain float64, s ...int) []float32 { _ = "STUB: not implemented"; return nil }

func HeEtAlN64(gain float64, s ...int) []float64 { _ = "STUB: not implemented"; return nil }

func HeEtAlU64(gain float64, s ...int) []float64 { _ = "STUB: not implemented"; return nil }
