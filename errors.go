package gorgonia

type NoOpError interface {
	NoOp() bool
}

type noopError struct{}

func (e noopError) NoOp() bool    { _ = "STUB: not implemented"; return false }
func (e noopError) Error() string { _ = "STUB: not implemented"; return "" }

type errNoStabilization interface {
	error
	noStabilization() bool
}

type noStabilizationErr struct{}

func (noStabilizationErr) Error() string         { _ = "STUB: not implemented"; return "" }
func (noStabilizationErr) noStabilization() bool { _ = "STUB: not implemented"; return false }

type noIncrErr struct {
	v Value
}

func (noIncrErr) Error() string  { _ = "STUB: not implemented"; return "" }
func (e noIncrErr) Value() Value { _ = "STUB: not implemented"; return *new(Value) }

type oomError struct {
	res       int64
	allocated int64
}

func (e oomError) Reserved() int64  { _ = "STUB: not implemented"; return 0 }
func (e oomError) Allocated() int64 { _ = "STUB: not implemented"; return 0 }
func (e oomError) Error() string    { _ = "STUB: not implemented"; return "" }

type AutoDiffError struct{}

func (err AutoDiffError) Error() string { _ = "STUB: not implemented"; return "" }

type vmContextualError struct {
	error
	node  *Node
	instr int
}

func (err vmContextualError) Node() *Node        { _ = "STUB: not implemented"; return nil }
func (err vmContextualError) Value() Value       { _ = "STUB: not implemented"; return *new(Value) }
func (err vmContextualError) InstructionID() int { _ = "STUB: not implemented"; return 0 }
func (err vmContextualError) Err() error         { _ = "STUB: not implemented"; return nil }

func nyi(what string, implFor interface{}) error { _ = "STUB: not implemented"; return nil }

func nondiffErr(op Op) error { _ = "STUB: not implemented"; return nil }

func checkErrSetDeriv(err error, dv *dualValue) error { _ = "STUB: not implemented"; return nil }

type SymDiffError struct {
	nodes   Nodes
	single  *Node
	grad    *Node
	gradMap map[*Node]Nodes
	err     error
}

func (err SymDiffError) Error() string { _ = "STUB: not implemented"; return "" }

func (err SymDiffError) Nodes() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (err SymDiffError) Node() *Node { _ = "STUB: not implemented"; return nil }

func (err SymDiffError) Grads() map[*Node]Nodes { _ = "STUB: not implemented"; return nil }

func (err SymDiffError) Grad() *Node { _ = "STUB: not implemented"; return nil }
