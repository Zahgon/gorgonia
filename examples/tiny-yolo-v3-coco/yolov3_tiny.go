package main

import (
	"gorgonia.org/gorgonia"
)

type YOLOv3 struct {
	g                                 *gorgonia.ExprGraph
	classesNum, boxesPerCell, netSize int
	out                               []*gorgonia.Node
	layersInfo                        []string
}

func (net *YOLOv3) Print() { _ = "STUB: not implemented"; return }

func (net *YOLOv3) GetOutput() []*gorgonia.Node { _ = "STUB: not implemented"; return nil }

func NewYoloV3Tiny(g *gorgonia.ExprGraph, input *gorgonia.Node, classesNumber, boxesPerCell int, leakyCoef float64, cfgFile, weightsFile string) (*YOLOv3, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
