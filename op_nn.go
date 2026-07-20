package gorgonia

import (
	"hash"
	"sync"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

var (
	_ SDOp = im2colOp{}
	_ Op   = col2imOp{}
	_ Op   = &maxPoolOp{}
	_ Op   = &maxPoolDiffOp{}
	_ Op   = &BatchNormOp{}
	_ Op   = &batchnormDiffOp{}
	_ Op   = &globalAveragePoolOp{}

	_ UsePreallocDoer = im2colOp{}
)

type randomness byte

const (
	uniform randomness = iota
	gaussian
	binomial
)

type randomOp struct {
	which randomness
	shape tensor.Shape
	dt    tensor.Dtype

	a, b float64
}

func makeRandomOp(which randomness, dt tensor.Dtype, a, b float64, shape ...int) randomOp {
	_ = "STUB: not implemented"
	return *new(randomOp)
}

func (op randomOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op randomOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op randomOp) InferShape(...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op randomOp) Do(...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op randomOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op randomOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op randomOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op randomOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op randomOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op randomOp) String() string { _ = "STUB: not implemented"; return "" }

type im2colOp struct {
	h, w                 int
	padH, padW           int
	strideH, strideW     int
	dilationH, dilationW int
}

func makeIm2ColOp(kernelHeight, kernelWidth, padHeight, padWidth, strideHeight, strideWidth, dilationHeight, dilationWidth int) im2colOp {
	_ = "STUB: not implemented"
	return *new(im2colOp)
}

func (op im2colOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op im2colOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op im2colOp) InferShape(shapes ...DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op im2colOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op im2colOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op im2colOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op im2colOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (op im2colOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op im2colOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op im2colOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op im2colOp) String() string { _ = "STUB: not implemented"; return "" }

func (op im2colOp) DiffWRT(i int) []bool { _ = "STUB: not implemented"; return nil }

func (op im2colOp) SymDiff(inputs Nodes, output, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op im2colOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op im2colOp) calcShape(s tensor.Shape) (retVal tensor.Shape) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape)
}

func (op im2colOp) retHW(h, w int) (retHeight, retWidth int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (op im2colOp) do(prealloc, input Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op im2colOp) f64s(chans, height, width, chanStride, inRowStride, retHeight, retWidth int, im, col []float64, wg *sync.WaitGroup, workers chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (op im2colOp) f32s(chans, height, width, chanStride, inRowStride, retHeight, retWidth int, im, col []float32, wg *sync.WaitGroup, workers chan struct{}) {
	_ = "STUB: not implemented"
	return
}

type col2imOp struct {
	unpaddedB int
	unpaddedC int
	unpaddedH int
	unpaddedW int

	im2colOp
}

func (op col2imOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op col2imOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op col2imOp) InferShape(shapes ...DimSizer) (retVal tensor.Shape, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op col2imOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op col2imOp) ReturnsPtr() bool     { _ = "STUB: not implemented"; return false }
func (op col2imOp) CallsExtern() bool    { _ = "STUB: not implemented"; return false }
func (op col2imOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op col2imOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op col2imOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op col2imOp) String() string { _ = "STUB: not implemented"; return "" }

func (op col2imOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op col2imOp) do(prealloc, input Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op col2imOp) f64s(chans, height, width, chanStride, retHeight, retWidth int, col, im []float64, wg *sync.WaitGroup, workers chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (op col2imOp) f32s(chans, height, width, chanStride, retHeight, retWidth int, col, im []float32, wg *sync.WaitGroup, workers chan struct{}) {
	_ = "STUB: not implemented"
	return
}

type maxPoolOp struct {
	unpaddedB int
	unpaddedC int
	unpaddedH int
	unpaddedW int

	h, w              int
	padNorth, padWest int
	padSouth, padEast int
	explicitPadding   bool
	strideH, strideW  int

	mask tensor.Tensor
}

func newMaxPoolOp(inputShape, kernel tensor.Shape, pad, stride []int) *maxPoolOp {
	_ = "STUB: not implemented"
	return nil
}

func (op *maxPoolOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *maxPoolOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *maxPoolOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *maxPoolOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *maxPoolOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op *maxPoolOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op *maxPoolOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op *maxPoolOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *maxPoolOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *maxPoolOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *maxPoolOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *maxPoolOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op *maxPoolOp) SymDiff(inputs Nodes, output, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *maxPoolOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *maxPoolOp) checkInput(inputs ...Value) (tensor.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func (op *maxPoolOp) calcShape(s tensor.Shape) tensor.Shape {
	_ = "STUB: not implemented"
	return *new(tensor.Shape)
}

func (op *maxPoolOp) strideValue(strides []int) int { _ = "STUB: not implemented"; return 0 }

func (op *maxPoolOp) do(out, in tensor.Tensor) { _ = "STUB: not implemented"; return }

func (op *maxPoolOp) f32s(batches, channels, outH, outW, inH, inW,
	outStride, inStride, maskStride int,
	outData, inData []float32,
	maskData []int) {
	_ = "STUB: not implemented"
	return
}

func (op *maxPoolOp) f64s(batches, channels, outH, outW, inH, inW,
	outStride, inStride, maskStride int,
	outData, inData []float64,
	maskData []int) {
	_ = "STUB: not implemented"
	return
}

type maxPoolDiffOp struct {
	maxPoolOp
}

func (op *maxPoolDiffOp) Arity() int    { _ = "STUB: not implemented"; return 0 }
func (op *maxPoolDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *maxPoolDiffOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *maxPoolDiffOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *maxPoolDiffOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op *maxPoolDiffOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op *maxPoolDiffOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op *maxPoolDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *maxPoolDiffOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *maxPoolDiffOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *maxPoolDiffOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *maxPoolDiffOp) checkInput(inputs ...Value) (in, pooled, pooledGrad tensor.Tensor, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), *new(tensor.Tensor), *new(tensor.Tensor), nil
}

func (op *maxPoolDiffOp) do(inGrad, in, pooled, pooledGrad tensor.Tensor) {
	_ = "STUB: not implemented"
	return
}

func (op *maxPoolDiffOp) f32s(batches, channels, pooledH, pooledW int,
	inStride, outStride, maskStride int,
	inDiffData, outDiffData []float32,
	maskData []int) {
	_ = "STUB: not implemented"
	return
}

func (op *maxPoolDiffOp) f64s(batches, channels, pooledH, pooledW int,
	inStride, outStride, maskStride int,
	inDiffData, outDiffData []float64,
	maskData []int) {
	_ = "STUB: not implemented"
	return
}

type clampOp struct {
	min, max Scalar
}

func (op *clampOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *clampOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *clampOp) InferShape(shps ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *clampOp) Do(vals ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *clampOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *clampOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *clampOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *clampOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *clampOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }
func (op *clampOp) String() string   { _ = "STUB: not implemented"; return "" }

type BatchNormOp struct {
	momentum float64
	epsilon  float64
	dims     int

	runningMean, runningVariance *tensor.Dense
	saveMean, saveVariance       *tensor.Dense

	alpha, beta *tensor.Dense

	training bool
}

func (op *BatchNormOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *BatchNormOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *BatchNormOp) InferShape(ns ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *BatchNormOp) Do(values ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *BatchNormOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (op *BatchNormOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (op *BatchNormOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (op *BatchNormOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *BatchNormOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *BatchNormOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *BatchNormOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *BatchNormOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op *BatchNormOp) SymDiff(inputs Nodes, output *Node, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *BatchNormOp) UsePreallocDo(prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *BatchNormOp) Stats() (runningMean tensor.Tensor, runningVariance tensor.Tensor) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), *new(tensor.Tensor)
}

func (op *BatchNormOp) SetStats(runningMean tensor.Tensor, runningVariance tensor.Tensor) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *BatchNormOp) SetTraining(isTraining bool) error { _ = "STUB: not implemented"; return nil }

func (op *BatchNormOp) Reset() error { _ = "STUB: not implemented"; return nil }

func (op *BatchNormOp) updateStatsF64(batchSize, channels, spatialDim int, inputT *tensor.Dense) (saveMean []float64, saveVar []float64) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *BatchNormOp) calculateAlphaAndBetaF64(batchSize, channels, spatialDim int, saveMean, saveVar []float64) (alpha []float64, beta []float64) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *BatchNormOp) f64s(input, output *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *BatchNormOp) updateStatsF32(batchSize, channels, spatialDim int, inputT *tensor.Dense) (saveMean []float32, saveVar []float32) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *BatchNormOp) calculateAlphaAndBetaF32(batchSize, channels, spatialDim int, saveMean, saveVar []float32) (alpha []float32, beta []float32) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (op *BatchNormOp) f32s(input, output *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type batchnormDiffOp struct{ *BatchNormOp }

func (op *batchnormDiffOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *batchnormDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *batchnormDiffOp) InferShape(ns ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *batchnormDiffOp) Do(values ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *batchnormDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *batchnormDiffOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *batchnormDiffOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *batchnormDiffOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op *batchnormDiffOp) SymDiff(inputs Nodes, output *Node, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *batchnormDiffOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *batchnormDiffOp) UsePreallocDo(prealloc Value, inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *batchnormDiffOp) f64s(input, prealloc, outGrad *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *batchnormDiffOp) f32s(input, prealloc, outGrad *tensor.Dense) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type globalAveragePoolOp struct{}

func (g *globalAveragePoolOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (g *globalAveragePoolOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (g *globalAveragePoolOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (g *globalAveragePoolOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (g *globalAveragePoolOp) ReturnsPtr() bool { _ = "STUB: not implemented"; return false }

func (g *globalAveragePoolOp) CallsExtern() bool { _ = "STUB: not implemented"; return false }

func (g *globalAveragePoolOp) OverwritesInput() int { _ = "STUB: not implemented"; return 0 }

func (g *globalAveragePoolOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (g *globalAveragePoolOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (g *globalAveragePoolOp) String() string { _ = "STUB: not implemented"; return "" }
