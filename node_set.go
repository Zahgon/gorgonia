package gorgonia

type NodeSet map[*Node]struct{}

func NewNodeSet(a ...*Node) NodeSet { _ = "STUB: not implemented"; return *new(NodeSet) }

func (set NodeSet) ToSlice() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (set NodeSet) Add(i *Node) bool { _ = "STUB: not implemented"; return false }

func (set NodeSet) Contains(i *Node) bool { _ = "STUB: not implemented"; return false }

func (set NodeSet) ContainsAll(i ...*Node) bool { _ = "STUB: not implemented"; return false }

func (set NodeSet) IsSubset(other NodeSet) bool { _ = "STUB: not implemented"; return false }

func (set NodeSet) IsSuperset(other NodeSet) bool { _ = "STUB: not implemented"; return false }

func (set NodeSet) Union(other NodeSet) NodeSet { _ = "STUB: not implemented"; return *new(NodeSet) }

func (set NodeSet) Intersect(other NodeSet) NodeSet {
	_ = "STUB: not implemented"
	return *new(NodeSet)
}

func (set NodeSet) Difference(other NodeSet) NodeSet {
	_ = "STUB: not implemented"
	return *new(NodeSet)
}

func (set NodeSet) SymmetricDifference(other NodeSet) NodeSet {
	_ = "STUB: not implemented"
	return *new(NodeSet)
}

func (set *NodeSet) Clear() { _ = "STUB: not implemented"; return }

func (set NodeSet) Remove(i *Node) { _ = "STUB: not implemented"; return }

func (set NodeSet) Cardinality() int { _ = "STUB: not implemented"; return 0 }

func (set NodeSet) Iter() <-chan *Node { _ = "STUB: not implemented"; return nil }

func (set NodeSet) Equal(other NodeSet) bool { _ = "STUB: not implemented"; return false }

func (set NodeSet) Clone() NodeSet { _ = "STUB: not implemented"; return *new(NodeSet) }
