//go:build !cuda
// +build !cuda

package gorgonia

func (op elemUnaryOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }
func (op elemBinOp) CallsExtern() bool   { _ = "STUB: not implemented"; return false }
func (op linAlgBinOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func NewAddOp(a, b *Node, ctx ExecutionContext) *ExternalOp { _ = "STUB: not implemented"; return nil }

func NewSubOp(a, b *Node, ctx ExecutionContext) *ExternalOp { _ = "STUB: not implemented"; return nil }

func NewHadamardProdOp(a, b *Node, ctx ExecutionContext) *ExternalOp {
	_ = "STUB: not implemented"
	return nil
}
