package gorgonia

import (
	"fmt"
)

type Nodes []*Node

func (ns Nodes) Node() *Node { _ = "STUB: not implemented"; return nil }

func (ns Nodes) Nodes() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (ns Nodes) Err() error { _ = "STUB: not implemented"; return nil }

func (ns Nodes) Len() int           { _ = "STUB: not implemented"; return 0 }
func (ns Nodes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ns Nodes) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (ns Nodes) Set() Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (ns Nodes) Add(n *Node) Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (ns Nodes) Contains(want *Node) bool { _ = "STUB: not implemented"; return false }

func (ns Nodes) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (ns Nodes) Difference(other Nodes) Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (ns Nodes) Intersect(other Nodes) Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (ns Nodes) AllSameGraph() bool { _ = "STUB: not implemented"; return false }

func (ns Nodes) Equals(other Nodes) bool { _ = "STUB: not implemented"; return false }

func (ns Nodes) mapSet() NodeSet { _ = "STUB: not implemented"; return *new(NodeSet) }

func (ns Nodes) index(n *Node) int { _ = "STUB: not implemented"; return 0 }

func (ns Nodes) reverse() { _ = "STUB: not implemented"; return }

func (ns Nodes) replace(what, with *Node) Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

var removers = make(map[string]int)

func (ns Nodes) remove(what *Node) Nodes { _ = "STUB: not implemented"; return *new(Nodes) }

func (ns Nodes) dimSizers() []DimSizer { _ = "STUB: not implemented"; return nil }
