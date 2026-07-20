//go:build cuda
// +build cuda

package gorgonia

import (
	"log"
	"sync"

	"github.com/pkg/errors"
	"gorgonia.org/cu"
	cudnn "gorgonia.org/cu/dnn"
	"gorgonia.org/gorgonia/cuda"
	"gorgonia.org/tensor"
)

const CUDA = true

var (
	_ External    = &ExternMetadata{}
	_ CUDAMachine = &tapeMachine{}
	_ CUDAMachine = &lispMachine{}
)

const (
	memalign    = 32
	scalarAlign = 8
)

//go:generate cudagen -same-module

var cudaStdLib []cudaLib

type cudaLib struct {
	name  string
	data  string
	funcs []string
}

type CUDAMachine interface {
	External
	Engines() []cuda.Engine
	Contexts() []*cu.BatchedContext
	CUDNNContexts() []*cudnn.Context

	ElemGridSize(n, dev int) (gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ int)
}

type ExternMetadata struct {
	tensor.Engine
	sync.Mutex

	u cu.Device
	b batchedBLAS

	engines       []cuda.Engine
	workAvailable chan bool
	syncChan      chan struct{}
	initialized   bool
}

func (m *ExternMetadata) ElemGridSize(n, dev int) (gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0, 0
}

func (m *ExternMetadata) WorkAvailable() <-chan bool { _ = "STUB: not implemented"; return nil }

func (m *ExternMetadata) Sync() chan struct{} { _ = "STUB: not implemented"; return nil }

func (m *ExternMetadata) DoWork() error { _ = "STUB: not implemented"; return nil }

func (m *ExternMetadata) Engines() []cuda.Engine { _ = "STUB: not implemented"; return nil }

func (m *ExternMetadata) Contexts() []*cu.BatchedContext { _ = "STUB: not implemented"; return nil }

func (m *ExternMetadata) CUDNNContexts() []*cudnn.Context { _ = "STUB: not implemented"; return nil }

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

func (m *ExternMetadata) Signal() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) Reset() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) init(sizes []int64) (err error) {
	m.Lock()
	initialized := m.initialized
	m.Unlock()
	if initialized {
		return nil
	}
	devices, err := cu.NumDevices()
	if err != nil {
		return errors.Wrapf(err, "Failed to get number of devices")
	}

	if devices == 0 {
		return errors.New("No Devices Found")
	}

	cudaLogf("Creating Engines")
	m.Lock()
	defer m.Unlock()
	m.engines = make([]cuda.Engine, len(sizes))
	for i := range m.engines {
		e := &m.engines[i]
		dev, err := cu.GetDevice(i)
		if err != nil {
			return errors.Wrapf(err, "Failed to get device %d", i)
		}

		if err = e.Init(dev, sizes[i]); err != nil {
			return err
		}
		ctx := e.Context()
		go m.collectWork(i, ctx.WorkAvailable())
	}

	m.initialized = true
	cudaLogf("CUDA initialized. Engines: %v", m.engines)
	return nil
}

func (m *ExternMetadata) initFail() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) cleanup() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) collectWork(devID int, workAvailable <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (m *ExternMetadata) collectBLASWork() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) signal() { _ = "STUB: not implemented"; return }

func (m *ExternMetadata) setEngine(e tensor.Engine) { _ = "STUB: not implemented"; return }

func AddToStdLib(name, data string, funcs []string) { _ = "STUB: not implemented"; return }

func init() {
	log.Println("Using CUDA build")
}

func (n *Node) ValueOnDevice(toDev Device, extern External) (retVal Value, allocOnExtern bool, err error) {
	_ = "STUB: not implemented"
	return *new(Value), false, nil
}

func (n *Node) GradOnDevice(toDev Device, extern External) (retVal Value, allocOnExtern bool, err error) {
	_ = "STUB: not implemented"
	return *new(Value), false, nil
}
