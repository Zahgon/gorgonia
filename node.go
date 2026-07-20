package gorgonia

import (
	"hash"

	"github.com/awalterschulze/gographviz"
	"github.com/chewxy/hm"
	"gorgonia.org/gorgonia/internal/encoding"
	"gorgonia.org/tensor"
)

type NodeID int64

type Node struct {
	t     hm.Type
	shape tensor.Shape

	op       Op
	children Nodes

	name string

	group string

	groups encoding.Groups

	g *ExprGraph

	boundTo Value
	dataOn  Device

	derivOf Nodes
	deriv   *Node

	id   int64
	hash uint32

	hashed        bool
	inferredShape bool
	unchanged     bool
	isStmt        bool
	ofInterest    bool
}

type NodeConsOpt func(*Node)

func WithType(t hm.Type) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func WithChildren(children Nodes) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func WithOp(op Op) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func In(g *ExprGraph) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func WithName(name string) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func WithValue(any interface{}) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func WithGrad(any interface{}) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func WithInit(fn InitWFn) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func WithShape(shp ...int) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func WithGroupName(name string) NodeConsOpt { _ = "STUB: not implemented"; return *new(NodeConsOpt) }

func withGroup(group encoding.Group) NodeConsOpt {
	_ = "STUB: not implemented"
	return *new(NodeConsOpt)
}

func (n *Node) Groups() encoding.Groups { _ = "STUB: not implemented"; return *new(encoding.Groups) }

func newNode(opts ...NodeConsOpt) *Node { _ = "STUB: not implemented"; return nil }

func NewUniqueNode(opts ...NodeConsOpt) *Node { _ = "STUB: not implemented"; return nil }

func (n *Node) ID() int64 { _ = "STUB: not implemented"; return 0 }

func (n *Node) Node() *Node { _ = "STUB: not implemented"; return nil }

func (n *Node) Nodes() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (n *Node) Err() error { _ = "STUB: not implemented"; return nil }

func (n *Node) DataSize() int { _ = "STUB: not implemented"; return 0 }

func (n *Node) DerivOf() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (n *Node) Deriv() *Node { _ = "STUB: not implemented"; return nil }

func (n *Node) isArg() bool      { _ = "STUB: not implemented"; return false }
func (n *Node) isInput() bool    { _ = "STUB: not implemented"; return false }
func (n *Node) isMutable() bool  { _ = "STUB: not implemented"; return false }
func (n *Node) isConstant() bool { _ = "STUB: not implemented"; return false }
func (n *Node) isRandom() bool   { _ = "STUB: not implemented"; return false }

func (n *Node) isRoot() bool { _ = "STUB: not implemented"; return false }

func (n *Node) IsVar() bool { _ = "STUB: not implemented"; return false }

func (n *Node) IsScalar() bool { _ = "STUB: not implemented"; return false }

func (n *Node) IsVector() bool { _ = "STUB: not implemented"; return false }

func (n *Node) IsColVec() bool { _ = "STUB: not implemented"; return false }

func (n *Node) IsRowVec() bool { _ = "STUB: not implemented"; return false }

func (n *Node) IsMatrix() bool { _ = "STUB: not implemented"; return false }

func (n *Node) Graph() *ExprGraph { _ = "STUB: not implemented"; return nil }

func (n *Node) CloneTo(g *ExprGraph) *Node { _ = "STUB: not implemented"; return nil }

func (n *Node) Clone() (retVal interface{}) { _ = "STUB: not implemented"; return nil }

func (n *Node) Value() Value { _ = "STUB: not implemented"; return *new(Value) }

func (n *Node) Grad() (Value, error) { _ = "STUB: not implemented"; return *new(Value), nil }

func (n *Node) Dims() int { _ = "STUB: not implemented"; return 0 }

func (n *Node) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (n *Node) Dtype() tensor.Dtype { _ = "STUB: not implemented"; return *new(tensor.Dtype) }

func (n *Node) Shape() tensor.Shape { _ = "STUB: not implemented"; return *new(tensor.Shape) }

func (n *Node) Strides() []int { _ = "STUB: not implemented"; return nil }

func (n *Node) Device() Device { _ = "STUB: not implemented"; return *new(Device) }

func (n *Node) Op() Op { _ = "STUB: not implemented"; return *new(Op) }

func (n *Node) IsVec() bool { _ = "STUB: not implemented"; return false }

func (n *Node) Name() string { _ = "STUB: not implemented"; return "" }

func (n *Node) WriteHash(h hash.Hash32) { _ = "STUB: not implemented"; return }

func (n *Node) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (n *Node) ToDot() string { _ = "STUB: not implemented"; return "" }

func (n *Node) RestrictedToDot(up, down int) string { _ = "STUB: not implemented"; return "" }

func (n *Node) String() string { _ = "STUB: not implemented"; return "" }

func (n *Node) bind(v Value) error { _ = "STUB: not implemented"; return nil }

func (n *Node) bindCopy(v Value) (err error) { _ = "STUB: not implemented"; return nil }

func (n *Node) unbind() { _ = "STUB: not implemented"; return }

func (n *Node) dotCluster() string { _ = "STUB: not implemented"; return "" }

func (n *Node) dot(g *gographviz.Escape, graphName string, seen map[*Node]string) string {
	_ = "STUB: not implemented"
	return ""
}

func (n *Node) fix() { _ = "STUB: not implemented"; return }

func (n *Node) fixChildren() { _ = "STUB: not implemented"; return }

func (n *Node) fixEdges() { _ = "STUB: not implemented"; return }

func (n *Node) setShape(s tensor.Shape, inferred bool) { _ = "STUB: not implemented"; return }

func (n *Node) setGroup(grp string) { _ = "STUB: not implemented"; return }

func (n *Node) clone(opts ...NodeConsOpt) *Node { _ = "STUB: not implemented"; return nil }

func (n *Node) diffWRT() []bool { _ = "STUB: not implemented"; return nil }

func (n *Node) seqWalk() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (n *Node) dotString(g *gographviz.Escape, graphName string) string {
	_ = "STUB: not implemented"
	return ""
}
