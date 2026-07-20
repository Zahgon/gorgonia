package main

import (
	"image"
)

var (
	anchors       = []float32{0.57273, 0.677385, 1.87446, 2.06253, 3.33843, 5.47434, 7.88282, 3.52778, 9.77052, 9.16828}
	classes       = []string{"person", "bicycle", "car", "motorbike", "aeroplane", "bus", "train", "truck", "boat", "traffic light", "fire hydrant", "stop sign", "parking meter", "bench", "bird", "cat", "dog", "horse", "sheep", "cow", "elephant", "bear", "zebra", "giraffe", "backpack", "umbrella", "handbag", "tie", "suitcase", "frisbee", "skis", "snowboard", "sports ball", "kite", "baseball bat", "baseball glove", "skateboard", "surfboard", "tennis racket", "bottle", "wine glass", "cup", "fork", "knife", "spoon", "bowl", "banana", "apple", "sandwich", "orange", "broccoli", "carrot", "hot dog", "pizza", "donut", "cake", "chair", "sofa", "pottedplant", "bed", "diningtable", "toilet", "tvmonitor", "laptop", "mouse", "remote", "keyboard", "cell phone", "microwave", "oven", "toaster", "sink", "refrigerator", "book", "clock", "vase", "scissors", "teddy bear", "hair drier", "toothbrush"}
	scoreTreshold = float32(0.6)
	iouTreshold   = float32(0.2)
)

type Detections []*DetectionRectangle

func (detections Detections) Len() int      { _ = "STUB: not implemented"; return 0 }
func (detections Detections) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (detections Detections) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type DetectionsOrder []*DetectionRectangle

func (detections DetectionsOrder) Len() int      { _ = "STUB: not implemented"; return 0 }
func (detections DetectionsOrder) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (detections DetectionsOrder) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type DetectionRectangle struct {
	conf  float32
	rect  image.Rectangle
	class string
	score float32
}

func (dr DetectionRectangle) GetClass() string { _ = "STUB: not implemented"; return "" }

func (tiny *TinyYOLOv2Net) ProcessOutput() (Detections, error) {
	_ = "STUB: not implemented"
	return *new(Detections), nil
}

func nonMaxSupr(detections Detections) Detections {
	_ = "STUB: not implemented"
	return *new(Detections)
}
