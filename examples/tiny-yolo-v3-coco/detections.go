package main

import (
	"image"
)

type DetectionRectangle struct {
	conf  float32
	rect  image.Rectangle
	class string
	score float32
}

func (dr *DetectionRectangle) String() string { _ = "STUB: not implemented"; return "" }

func (dr *DetectionRectangle) GetClass() string { _ = "STUB: not implemented"; return "" }

type Detections []*DetectionRectangle

func (detections Detections) Len() int      { _ = "STUB: not implemented"; return 0 }
func (detections Detections) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (detections Detections) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type DetectionsOrder []*DetectionRectangle

func (detections DetectionsOrder) Len() int      { _ = "STUB: not implemented"; return 0 }
func (detections DetectionsOrder) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (detections DetectionsOrder) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (net *YOLOv3) ProcessOutput(classes []string, scoreTreshold, iouTreshold float32) (Detections, error) {
	_ = "STUB: not implemented"
	return *new(Detections), nil
}

func prepareDetections(data []float32, scoreTreshold float32, netSize int, classes []string) Detections {
	_ = "STUB: not implemented"
	return *new(Detections)
}

func nonMaxSupr(detections Detections, iouTreshold float32) Detections {
	_ = "STUB: not implemented"
	return *new(Detections)
}
