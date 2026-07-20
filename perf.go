package gorgonia

import (
	"sync"

	"github.com/chewxy/hm"
	"gorgonia.org/tensor"
)

var nodePool = &sync.Pool{
	New: func() interface{} { return new(Node) },
}

func borrowNode() *Node { _ = "STUB: not implemented"; return nil }

func returnNode(n *Node) { _ = "STUB: not implemented"; return }

func ReturnNode(n *Node) { _ = "STUB: not implemented"; return }

var dvpool = &sync.Pool{
	New: func() interface{} { return new(dualValue) },
}

func borrowDV() *dualValue { _ = "STUB: not implemented"; return nil }

func returnDV(dv *dualValue) { _ = "STUB: not implemented"; return }

func returnTensor(t tensor.Tensor) { _ = "STUB: not implemented"; return }

func returnValue(v Value) { _ = "STUB: not implemented"; return }

var dimSizerPool = new(sync.Map)

func borrowDimSizers(size int) []DimSizer { _ = "STUB: not implemented"; return nil }

func returnDimSizers(ds []DimSizer) { _ = "STUB: not implemented"; return }

var tensorTypePool = &sync.Pool{
	New: func() interface{} { return new(TensorType) },
}

func borrowTensorType() *TensorType { _ = "STUB: not implemented"; return nil }

func returnTensorType(t *TensorType) { _ = "STUB: not implemented"; return }

func ReturnType(t hm.Type) { _ = "STUB: not implemented"; return }
