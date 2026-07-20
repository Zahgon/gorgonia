package gorgonia

const (
	bitmapBits = 64
)

type bitmap struct {
	n   []uint64
	max int
}

func newBitmap(size int) *bitmap { _ = "STUB: not implemented"; return nil }

func (bm *bitmap) Set(i int) { _ = "STUB: not implemented"; return }

func (bm *bitmap) IsSet(i int) bool { _ = "STUB: not implemented"; return false }

func (bm *bitmap) Clear(i int) { _ = "STUB: not implemented"; return }

func (bm *bitmap) BlocksWithZero(atleast int) int { _ = "STUB: not implemented"; return 0 }
