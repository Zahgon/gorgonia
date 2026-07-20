//go:build !fastmath
// +build !fastmath

package gorgonia

func SetOptimizationLevel(i int) { _ = "STUB: not implemented"; return }

func _inversef32(x float32) float32 { _ = "STUB: not implemented"; return 0 }
func _inversef64(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func _tanhf32(x float32) float32 { _ = "STUB: not implemented"; return 0 }
func _tanhf64(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func _sigmoidf64(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func _sigmoidf32(x float32) float32 { _ = "STUB: not implemented"; return 0 }

func _inverseSqrtf64(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func _inverseSqrtf32(x float32) float32 { _ = "STUB: not implemented"; return 0 }
