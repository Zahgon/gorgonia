package gorgonia

import (
	"sync"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/gonum"
)

var blasdoor sync.Mutex
var whichblas BLAS

type BLAS interface {
	blas.Float32
	blas.Float64
	blas.Complex64
	blas.Complex128
}

type batchedBLAS interface {
	WorkAvailable() <-chan struct{}
	DoWork()
	Close() error
	BLAS
}

func Use(b BLAS) { _ = "STUB: not implemented"; return }

func WhichBLAS() BLAS { _ = "STUB: not implemented"; return *new(BLAS) }

func init() {
	whichblas = gonum.Implementation{}
}
