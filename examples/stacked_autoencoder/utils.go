package main

import (
	"image"
)

type sli struct {
	start, end, step int
}

func (s sli) Start() int { _ = "STUB: not implemented"; return 0 }
func (s sli) End() int   { _ = "STUB: not implemented"; return 0 }
func (s sli) Step() int  { _ = "STUB: not implemented"; return 0 }

func s(start int) sli { _ = "STUB: not implemented"; return *new(sli) }

func hasOne(a []float64) bool { _ = "STUB: not implemented"; return false }

func avgF64s(a []float64) (retVal float64) { _ = "STUB: not implemented"; return 0 }

const numLabels = 10
const pixelRange = 255

func pixelWeight(px byte) float64 { _ = "STUB: not implemented"; return 0 }

func reversePixelWeight(px float64) byte { _ = "STUB: not implemented"; return 0 }

func visualizeRow(x []float64) *image.Gray { _ = "STUB: not implemented"; return nil }
