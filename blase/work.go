package blase

/*
#cgo CFLAGS: -g -O3 -std=gnu99

#include <stdio.h>
#include <stdint.h>
#include "work.h"
#include "cblas.h"

uintptr_t process(struct fnargs* fa, int count) {
	uintptr_t ret;

	// printf("How much work: %d\n", count);

	ret = processFn(&fa[0]);
	if (count > 1) {
		ret = processFn(&fa[1]);
	}
	if (count > 2) {
		ret = processFn(&fa[2]);
	}

	return ret;
}

*/
import "C"
import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/netlib/blas/netlib"
)

var impl = newContext()

var (
	_ blas.Float32    = impl
	_ blas.Float64    = impl
	_ blas.Complex64  = impl
	_ blas.Complex128 = impl
)

func Implementation() *context { _ = "STUB: not implemented"; return nil }

const workbufLen int = 3

type Worker interface {
	WorkAvailable() <-chan struct{}
	DoWork()
}

type call struct {
	args *fnargs

	blocking bool
}

type context struct {
	netlib.Implementation

	workAvailable chan struct{}
	work          chan call

	fns   []C.struct_fnargs
	queue []call
}

func newContext() *context { _ = "STUB: not implemented"; return nil }

func (ctx *context) enqueue(c call) { _ = "STUB: not implemented"; return }

func (ctx *context) DoWork() { _ = "STUB: not implemented"; return }

func (ctx *context) WorkAvailable() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (*context) String() string { _ = "STUB: not implemented"; return "" }
