package main

import (
	"image"
	"image/color"
)

func Float32frombytes(bytes []byte) float32 { _ = "STUB: not implemented"; return 0 }

func GetFloat32Image(fname string, resizeWidth, resizeHeight int) ([]float32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Image2Float32(img image.Image) ([]float32, error) { _ = "STUB: not implemented"; return nil, nil }

func resizeImage(img image.Image, width int, height int) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func getAverageColor(img image.Image, minX int, maxX int, minY int, maxY int) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func IOUFloat32(r1, r2 image.Rectangle) float32 { _ = "STUB: not implemented"; return 0 }

func Softmax(a []float32) []float32 { _ = "STUB: not implemented"; return nil }

func MaxFloat32(cl []float32) (float32, int) { _ = "STUB: not implemented"; return 0, 0 }

func Rectify(x, y, w, h, maxwidth, maxheight int) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func MaxInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func MinInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func SigmoidF32(sum float32) float32 { _ = "STUB: not implemented"; return 0 }
