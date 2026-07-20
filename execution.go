package gorgonia

import (
	"gorgonia.org/tensor"
)

type Arena interface {
	Get(dev Device, size int64) (tensor.Memory, error)
	GetFromValue(dev Device, v Value) (tensor.Memory, error)
	Put(dev Device, mem tensor.Memory, size int64)
	PutValue(dev Device, v Value)

	Transfer(toDev, fromDev Device, v Value, synchronous bool) (retVal Value, err error)
}

type External interface {
	Arena
	Signal()
	Sync() chan struct{}
}

type ExecutionContext struct {
	External
	Device
}

type ExternalOp struct {
	Op
	ExecutionContext

	Prealloc  Value
	Incr      Value
	UseUnsafe bool
}

func NewExternalOp(op Op, ctx ExecutionContext, prealloc Value) *ExternalOp {
	_ = "STUB: not implemented"
	return nil
}

func (op *ExternalOp) DetermineDevice(inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *ExternalOp) Do(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *ExternalOp) String() string { _ = "STUB: not implemented"; return "" }
