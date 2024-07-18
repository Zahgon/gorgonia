package gorgonia

import (
	"fmt"
	"sync"

	"gorgonia.org/tensor"
)

var rndCounter int
var rndLock sync.Mutex

func randomName(a Tensor) string {
	rndLock.Lock()
	defer rndLock.Unlock()
	rndCounter++
	return fmt.Sprintf("Random_%d", rndCounter)
}

func Tensors2Basics[DT any](ts ...Tensor) []tensor.Basic[DT] {
	retVal := make([]tensor.Basic[DT], 0, len(ts))
	for _, t := range ts {
		retVal = append(retVal, t.(tensor.Basic[DT]))
	}
	return retVal
}
