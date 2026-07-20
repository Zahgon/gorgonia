package gorgonia

func forwardDiffAnalysis(outputs, sortedNodes Nodes) (retVal NodeSet, err error) {
	_ = "STUB: not implemented"
	return *new(NodeSet), nil
}

func backwardDiffAnalysis(wrt, sortedNodes Nodes) (retVal NodeSet, err error) {
	_ = "STUB: not implemented"
	return *new(NodeSet), nil
}

func Backpropagate(outputs, gradOutputs, wrt Nodes) (retVal Nodes, err error) {
	_ = "STUB: not implemented"
	return *new(Nodes), nil
}

func SetDerivOf(deriv, of *Node) { _ = "STUB: not implemented"; return }
