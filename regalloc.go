package gorgonia

type interval struct {
	start, end int

	result       register
	reads        []register
	ranges       []intervalRange
	usePositions []int
}

func newInterval() *interval { _ = "STUB: not implemented"; return nil }

func (i *interval) String() string { _ = "STUB: not implemented"; return "" }

func (i *interval) setFrom(from int) { _ = "STUB: not implemented"; return }

func (i *interval) fix() { _ = "STUB: not implemented"; return }

func (i *interval) addRange(from, to int) { _ = "STUB: not implemented"; return }

func (i *interval) addUsePositions(up int) { _ = "STUB: not implemented"; return }

func (i *interval) noUsePositions() bool { _ = "STUB: not implemented"; return false }

func (i *interval) liveAt(id int) bool { _ = "STUB: not implemented"; return false }

func (i *interval) lastUse() int { _ = "STUB: not implemented"; return 0 }

func (i *interval) merge(other *interval) { _ = "STUB: not implemented"; return }

type intervalRange struct {
	from, to int
}

type regalloc struct {
	cpucount      int
	gpucount      int
	instructionID int
	df            *dataflow
}

func newRegalloc(df *dataflow) *regalloc { _ = "STUB: not implemented"; return nil }

func (ra *regalloc) newReg(device Device) register {
	_ = "STUB: not implemented"
	return *new(register)
}

func (ra *regalloc) allocArg(nInterv *interval) { _ = "STUB: not implemented"; return }

func (ra *regalloc) allocMutableOp(node *Node, nInterv *interval) {
	_ = "STUB: not implemented"
	return
}

func (ra *regalloc) allocImmutableOp(node *Node, nInterv *interval) {
	_ = "STUB: not implemented"
	return
}

func (ra *regalloc) allocStatement(node *Node, nInterv *interval) {
	_ = "STUB: not implemented"
	return
}

func (ra *regalloc) alloc(sorted Nodes) { _ = "STUB: not implemented"; return }
