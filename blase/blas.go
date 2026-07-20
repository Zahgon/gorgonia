package blase

/*
#include <stdint.h>
#include <stdio.h>
#include "cblas.h"
#include "work.h"

// useful to help print stuff to see if things are correct
void prrintfnargs(struct fnargs* args){
	printf("HELLO\n");
	printf("fn: %d\n", args->fn);
	printf("o: %d\n", args->order);
	printf("tA: %d\n",args->tA);
	printf("tB: %d\n",args->tB);
	printf("----\n");
	// printf("a0: %f\n", (double*)args->a0);
	// printf("a1: %f\n", (double*)args->a1);
	// printf("a2: %f\n", (double*)args->a2);
	// printf("a3: %f\n", (double*)args->a3);
	printf("----\n");
	printf("i0: %d\n", args->i0);
	printf("i1: %d\n", args->i1);
	printf("i2: %d\n", args->i2);
	printf("i3: %d\n", args->i3);
	printf("i4: %d\n", args->i4);
	printf("i5: %d\n", args->i5);
	printf("----\n");
	printf("d0: %f\n", args->d0);
	printf("d1: %f\n", args->d1);
	printf("d2: %f\n", args->d2);
	printf("d3: %f\n", args->d3);
	printf("=========\n");
}
*/
import "C"

import (
	"gonum.org/v1/gonum/blas"
)

const rowMajor = 101

func (ctx *context) Dgemm(tA blas.Transpose, tB blas.Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (ctx *context) Dgemv(tA blas.Transpose, m int, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	_ = "STUB: not implemented"
	return
}
