package mnist

import (
	"io"
)

const numLabels = 10
const pixelRange = 255

const (
	imageMagic = 0x00000803
	labelMagic = 0x00000801

	Width = 28

	Height = 28
)

func readLabelFile(r io.Reader, e error) (labels []Label, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readImageFile(r io.Reader, e error) (imgs []RawImage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
