//go:build !cuda
// +build !cuda

package gorgonia

import "gorgonia.org/tensor"

const CUDA = false

var _ tensor.Engine = ExternMetadata{}

type ExternMetadata struct {
	tensor.Engine
	b             batchedBLAS
	workAvailable chan bool
	syncChan      chan struct{}
}

func (m *ExternMetadata) init() error {
	m.syncChan = make(chan struct{})
	if m.b != nil {
		m.workAvailable = make(chan bool)
		go m.collectBLASWork()
	}
	return nil
}

func (m *ExternMetadata) initFail() { _ = "STUB: not implemented"; return }

func (m ExternMetadata) HasFunc(name string) bool { _ = "STUB: not implemented"; return false }

func (m *ExternMetadata) WorkAvailable() <-chan bool { _ = "STUB: not implemented"; return nil }

func (m *ExternMetadata) Sync() chan struct{} { _ = "STUB: not implemented"; return nil }

func (m *ExternMetadata) DoWork() error { _ = "STUB: not implemented"; return nil }

func (m *ExternMetadata) Get(dev Device, size int64) (tensor.Memory, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Memory), nil
}

func (m *ExternMetadata) GetFromValue(dev Device, v Value) (tensor.Memory, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Memory), nil
}

func (m *ExternMetadata) Put(dev Device, mem tensor.Memory, size int64) {
	_ = "STUB: not implemented"
	return
}

func (m *ExternMetadata) PutValue(dev Device, v Value) { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) Transfer(toDev, fromDev Device, v Value, synchronous bool) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (m *ExternMetadata) Reset() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) Cleanup() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) Signal() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) collectBLASWork() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) signal() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) setEngine(e tensor.Engine) { _ = "STUB: not implemented"; return }

func (n *Node) ValueOnDevice(dev Device, extern External) (retVal Value, allocOnExtern bool, err error) {
	_ = "STUB: not implemented"
	return *new(Value), false, nil
}

func (n *Node) GradOnDevice(dev Device, extern External) (retVal Value, allocOnExtern bool, err error) {
	_ = "STUB: not implemented"
	return *new(Value), false, nil
}
