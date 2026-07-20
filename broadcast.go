package gorgonia

import (
	"gorgonia.org/tensor"
)

const (
	bcAllowableAxes = 4
)

type BroadcastPattern byte

func NewBroadcastPattern(leftAxes, rightAxes []byte) BroadcastPattern {
	_ = "STUB: not implemented"
	return *new(BroadcastPattern)
}

func (bcpat BroadcastPattern) bc(left bool, axis byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (bcpat BroadcastPattern) on() (retVal [2][]int) { _ = "STUB: not implemented"; return [2][]int{} }

func Broadcast(a, b *Node, pattern BroadcastPattern) (*Node, *Node, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func autoBroadcastPattern(aShape, bShape tensor.Shape) (leftPattern, rightPattern []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
