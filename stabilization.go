package gorgonia

var unaryOpStabilizationFns = make(map[ʘUnaryOperatorType][]func(*Node) (*Node, error))
var binOpStabilizationFns = make(map[ʘBinaryOperatorType][]func(*Node, *Node) (*Node, error))

func init() {
	unaryOpStabilizationFns[lnOpType] = []func(*Node) (*Node, error){
		logSigmoidStabilization,
		logStabilization,
		logSoftmaxStabilization,
	}
	binOpStabilizationFns[subOpType] = []func(*Node, *Node) (*Node, error){
		exp1mStabilization,
		oneMinusSigmoidStabilization,
	}
	unaryOpStabilizationFns[log1pOpType] = []func(*Node) (*Node, error){
		log1pExpStabilization,
		log1pNegSigmoidStabilization,
	}
	unaryOpStabilizationFns[negOpType] = []func(*Node) (*Node, error){negNegOptimization}
}

func logStabilization(a *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func exp1mStabilization(a, b *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func oneMinusSigmoidStabilization(a, b *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func logSigmoidStabilization(a *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func log1pExpStabilization(a *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func log1pNegSigmoidStabilization(a *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func logSoftmaxStabilization(a *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func negNegOptimization(a *Node) (retVal *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
