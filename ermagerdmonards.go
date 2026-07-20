package gorgonia

var (
	_ Result = (*Node)(nil)
	_ Result = (Nodes)(nil)
	_ Result = gErr{}
)

type Result interface {
	Input
	Errer
}

type Input interface {
	Node() *Node
	Nodes() Nodes
}

type Errer interface {
	Err() error
}

type Mker interface {
	Mk(...Input) Input
}

func Lift1(fn func(a *Node) (*Node, error)) func(a Input) Result {
	_ = "STUB: not implemented"
	return nil
}

func Lift1Axial(fn func(a *Node, axes ...int) (*Node, error)) func(a Input, axes ...int) Result {
	_ = "STUB: not implemented"
	return nil
}

func Lift2(fn func(a, b *Node) (*Node, error)) func(a, b Input) Result {
	_ = "STUB: not implemented"
	return nil
}

func Lift2Broadcast(fn func(a, b *Node, pat1, pat2 []byte) (*Node, error)) func(a, b Input, pat1, pat2 []byte) Result {
	_ = "STUB: not implemented"
	return nil
}

type gErr struct{ error }

func Err(e error) gErr { _ = "STUB: not implemented"; return *new(gErr) }

func (err gErr) Node() *Node  { _ = "STUB: not implemented"; return nil }
func (err gErr) Nodes() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }
func (err gErr) Err() error   { _ = "STUB: not implemented"; return nil }

type resultM struct{ Input }

func (r resultM) Err() error { _ = "STUB: not implemented"; return nil }

func LiftResult(a Input, err error) Result { _ = "STUB: not implemented"; return *new(Result) }

func TransformResult(ins ...Input) func(a Input, err error) Result {
	_ = "STUB: not implemented"
	return nil
}

func CheckOne(in Input) error { _ = "STUB: not implemented"; return nil }

func NodesFromInputs(xs ...Input) (Nodes, error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}
