package dataviz

import (
	"io"
	"time"

	"gorgonia.org/gorgonia"
	xvm "gorgonia.org/gorgonia/x/vm"
)

func DumpTrace(traces []xvm.Trace, g *gorgonia.ExprGraph, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

type group struct {
	Group string      `json:"group"`
	Data  []dataGroup `json:"data"`
}

type dataGroup struct {
	Label string      `json:"label"`
	Data  []dataLabel `json:"data"`
}

type dataLabel struct {
	TimeRange []time.Time `json:"timeRange"`
	Val       interface{} `json:"val"`
}
