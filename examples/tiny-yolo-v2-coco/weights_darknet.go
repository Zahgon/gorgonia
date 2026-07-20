package main

import (
	"gorgonia.org/tensor"
)

func ParseTinyYOLOv2(fname string) []float32 { _ = "STUB: not implemented"; return nil }

func PrepareData(biases, gammas, means, vars, kernels map[string][]float32, data []float32, layerName string, convShape tensor.Shape, lastIdx *int, batchNorm, biased bool) {
	_ = "STUB: not implemented"
	return
}

func DenormalizeWeights(biases, gammas, means, vars, kernels map[string][]float32, layerName string, convShape tensor.Shape, epsilon float32) {
	_ = "STUB: not implemented"
	return
}

func Float32frombytes(bytes []byte) float32 { _ = "STUB: not implemented"; return 0 }

func Float32bytes(float float32) []byte { _ = "STUB: not implemented"; return nil }
