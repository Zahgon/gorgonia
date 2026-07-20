package gorgonia

import (
	"hash"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

func AveragePool1D(x *Node, kernel, pad, stride int) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AveragePool2D(x *Node, kernel tensor.Shape, pad, stride []int) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type avgPoolOp struct {
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

func newAvgPoolOp(inputShape, kernel tensor.Shape, pad, stride []int) *avgPoolOp {
	_ = "STUB: not implemented"
	return nil
}

func (op *avgPoolOp) Arity() int { _ = "STUB: not implemented"; return 0 }

func (op *avgPoolOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *avgPoolOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *avgPoolOp) Do(inputs ...Value) (retVal Value, err error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *avgPoolOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op *avgPoolOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op *avgPoolOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op *avgPoolOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *avgPoolOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *avgPoolOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *avgPoolOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *avgPoolOp) DiffWRT(inputs int) []bool { _ = "STUB: not implemented"; return nil }

func (op *avgPoolOp) SymDiff(inputs Nodes, output, grad *Node) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func (op *avgPoolOp) DoDiff(ctx ExecutionContext, inputs Nodes, output *Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (op *avgPoolOp) checkInput(inputs ...Value) (tensor.Tensor, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), nil
}

func (op *avgPoolOp) calcShape(s tensor.Shape) tensor.Shape {
	_ = "STUB: not implemented"
	return *new(tensor.Shape)
}

func (op *avgPoolOp) do(out, in tensor.Tensor) { _ = "STUB: not implemented"; return }

func (op *avgPoolOp) f32s(batches, channels, outH, outW, inH, inW,
	outStride, inStride, maskStride int,
	outData, inData []float32,
	maskData []int) {
	_ = "STUB: not implemented"
	return
}

func (op *avgPoolOp) f64s(batches, channels, outH, outW, inH, inW,
	outStride, inStride, maskStride int,
	outData, inData []float64,
	maskData []int) {
	_ = "STUB: not implemented"
	return
}

type avgPoolDiffOp struct {
	avgPoolOp
}

func (op *avgPoolDiffOp) Arity() int    { _ = "STUB: not implemented"; return 0 }
func (op *avgPoolDiffOp) Type() hm.Type { _ = "STUB: not implemented"; return *new(hm.Type) }

func (op *avgPoolDiffOp) InferShape(inputs ...DimSizer) (tensor.Shape, error) {
	_ = "STUB: not implemented"
	return *new(tensor.Shape), nil
}

func (op *avgPoolDiffOp) Do(inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *avgPoolDiffOp) ReturnsPtr() bool      { _ = "STUB: not implemented"; return false }
func (op *avgPoolDiffOp) CallsExtern() bool     { _ = "STUB: not implemented"; return false }
func (op *avgPoolDiffOp) OverwritesInput() int  { _ = "STUB: not implemented"; return 0 }
func (op *avgPoolDiffOp) WriteHash(h hash.Hash) { _ = "STUB: not implemented"; return }

func (op *avgPoolDiffOp) Hashcode() uint32 { _ = "STUB: not implemented"; return 0 }

func (op *avgPoolDiffOp) String() string { _ = "STUB: not implemented"; return "" }

func (op *avgPoolDiffOp) UsePreallocDo(prealloc Value, inputs ...Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (op *avgPoolDiffOp) checkInput(inputs ...Value) (in, pooled, pooledGrad tensor.Tensor, err error) {
	_ = "STUB: not implemented"
	return *new(tensor.Tensor), *new(tensor.Tensor), *new(tensor.Tensor), nil
}

func (op *avgPoolDiffOp) do(inGrad, in, pooled, pooledGrad tensor.Tensor) {
	_ = "STUB: not implemented"
	return
}

func (op *avgPoolDiffOp) f32s(batches, channels, pooledH, pooledW, inH, inW int,
	inStride, outStride, maskStride int,
	inDiffData, outDiffData []float32,
	maskData []int) {
	_ = "STUB: not implemented"
	return
}

func (op *avgPoolDiffOp) f64s(batches, channels, pooledH, pooledW, inH, inW int,
	inStride, outStride, maskStride int,
	inDiffData, outDiffData []float64,
	maskData []int) {
	_ = "STUB: not implemented"
	return
}
