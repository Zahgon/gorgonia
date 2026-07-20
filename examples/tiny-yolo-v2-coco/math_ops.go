package main

import (
	"image"
)

func Rectify(x, y, h, w, maxwidth, maxheight int) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func IOUFloat32(r1, r2 image.Rectangle) float32 { _ = "STUB: not implemented"; return 0 }

func MaxInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func MinInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func Softmax(a []float32) []float32 { _ = "STUB: not implemented"; return nil }

func MaxFloat32(cl []float32) (float32, int) { _ = "STUB: not implemented"; return 0, 0 }

func Sigmoid(sum float32) float32 { _ = "STUB: not implemented"; return 0 }
