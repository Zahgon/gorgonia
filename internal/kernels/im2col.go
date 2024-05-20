package kernels

import (
	"context"
	"sync"

	"gorgonia.org/gorgonia/internal"
	"gorgonia.org/tensor"
)

type ImColParams struct {
	H, W                 int // kernel height and  width
	PadH, PadW           int
	StrideH, StrideW     int
	DilationH, DilationW int

	Chans  int
	Height int
	Width  int

	ChanStride int

	RetH, RetW int
}

// Im2Col implements an optimized version of the kernel for the Im2Col Op for CPU execution.
func Im2Col[DT tensor.Num](ctx context.Context, op ImColParams, im, col []DT, wg *sync.WaitGroup, workers chan struct{}) {
	if err := internal.HandleCtx(ctx); err != nil {
		return
	}
	workers <- struct{}{}

	chans, height, width := op.Chans, op.Height, op.Width
	chanStride := op.ChanStride
	retHeight := op.RetH
	retWidth := op.RetW
	var colIdx, inputRow, inputCol int
	var z DT
	for outputRow := 0; outputRow < retHeight; outputRow++ {
		for outputCol := 0; outputCol < retWidth; outputCol++ {
			for ch := 0; ch < chans; ch++ {
				for kernelRow := 0; kernelRow < op.H; kernelRow++ {
					inputRow = -op.PadH + kernelRow*op.DilationH + outputRow*op.StrideH
					for kernelCol := 0; kernelCol < op.W; kernelCol++ {
						if inputRow < 0 || inputRow >= height {
							col[colIdx] = z
							colIdx++
							continue
						}
						inputCol = -op.PadW + kernelCol*op.DilationW + outputCol*op.StrideW
						if inputCol < 0 || inputCol >= width {
							col[colIdx] = z
						} else {
							imIdx := chanStride*ch + inputRow*width + inputCol
							col[colIdx] = im[imIdx]
						}
						colIdx++
					}
				}
			}
		}
	}
	<-workers
	wg.Done()
}

// Col2Im is the differentiation operation for Im2Col.
func Col2Im[DT tensor.Num](ctx context.Context, op ImColParams, col, im []DT, wg *sync.WaitGroup, workers chan struct{}) {
	if err := internal.HandleCtx(ctx); err != nil {
		return
	}
	workers <- struct{}{}

	chans, height, width := op.Chans, op.Height, op.Width
	chanStride := op.ChanStride
	retHeight := op.RetH
	retWidth := op.RetW

	// memset im to 0
	var z DT
	for i := 0; i < len(im); i++ {
		im[i] = z
	}
	colIdx := 0
	var inputRow int
	var inputCol int
	for outputRow := 0; outputRow < retHeight; outputRow++ {
		for outputCol := 0; outputCol < retWidth; outputCol++ {
			for ch := 0; ch < chans; ch++ {
				for kernelRow := 0; kernelRow < op.H; kernelRow++ {
					inputRow = -op.PadH + kernelRow*op.DilationH + outputRow*op.StrideH
					for kernelCol := 0; kernelCol < op.W; kernelCol++ {
						if inputRow < 0 || inputRow >= height {
							colIdx++
							continue
						}
						inputCol = -op.PadW + kernelCol*op.DilationW + outputCol*op.StrideW
						if inputCol >= 0 && inputCol < width {
							imIdx := chanStride*ch + inputRow*width + inputCol
							im[imIdx] += col[colIdx]
						}
						colIdx++
					}
				}
			}
		}
	}
	<-workers
	wg.Done()
}
