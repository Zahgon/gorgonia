package cuda

const (
	minAllocBits = 8
	minAllocSize = 1 << minAllocBits

	freeAllocTresh = 0.75
)

var nilBlock = memblock{}

type memblock struct {
	address uintptr
	size    int64

	next, prev *memblock
}

func newMemblock(addr uintptr, size int64) *memblock { _ = "STUB: not implemented"; return nil }

func (a memblock) cap() uintptr { _ = "STUB: not implemented"; return 0 }

func (a *memblock) overlaps(b *memblock) bool { _ = "STUB: not implemented"; return false }

func (a *memblock) split(size int64) (b *memblock) { _ = "STUB: not implemented"; return nil }

func (a *memblock) lt(b *memblock) bool { _ = "STUB: not implemented"; return false }

func (a *memblock) String() string { _ = "STUB: not implemented"; return "" }

type freelist struct {
	first, last *memblock
	l           int
}

func (l *freelist) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *freelist) String() string { _ = "STUB: not implemented"; return "" }

func (l *freelist) insert(block *memblock) { _ = "STUB: not implemented"; return }

func (l *freelist) remove(block *memblock) { _ = "STUB: not implemented"; return }

func (l *freelist) splitOrRemove(block *memblock, aligned, size int64) {
	_ = "STUB: not implemented"
	return
}

type bfc struct {
	start        uintptr
	size         int64
	blockSize    int64
	reservedSize int64

	freelist *freelist
	used     map[uintptr]int64

	allocated int64
	allocs    int
	frees     int
}

func newBFC(alignment int64) *bfc { _ = "STUB: not implemented"; return nil }

func makeBFC(alignment int64) bfc { _ = "STUB: not implemented"; return *new(bfc) }

func (b *bfc) reset() { _ = "STUB: not implemented"; return }

func (b *bfc) reserve(start uintptr, size int64) { _ = "STUB: not implemented"; return }

func (b *bfc) release() uintptr { _ = "STUB: not implemented"; return 0 }

func (b *bfc) alloc(size int64) (mem uintptr, err error) { _ = "STUB: not implemented"; return 0, nil }

func (b *bfc) free(address uintptr) { _ = "STUB: not implemented"; return }

func (b *bfc) bestFit(size int64) (best *memblock) { _ = "STUB: not implemented"; return nil }

func (b *bfc) coalesce() { _ = "STUB: not implemented"; return }

func (b *bfc) align(size int64) int64 { _ = "STUB: not implemented"; return 0 }
