package gorgonia

import (
	"log"

	"gorgonia.org/tensor"
)

type VM interface {
	RunAll() error
	Reset()

	Close() error
}

const (
	fwdOnly byte = iota
	bwdOnly
	watchNaN
	watchInf
	watchPointer
	allocVals
	spare2
	spare3
	watchAll
)

type VMOpt func(m VM)

func EvalMode() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func WithLogger(logger *log.Logger) VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func WithValueFmt(format string) VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func WithWatchlist(list ...interface{}) VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func WithNaNWatch() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func WithInfWatch() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func WithPointerWatch() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func ExecuteFwdOnly() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func ExecuteBwdOnly() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func LogFwd() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func LogBwd() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func LogBothDir() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func TraceExec() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func BindDualValues(nodes ...*Node) VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func WithPrecompiled(prog *program, locMap map[*Node]register) VMOpt {
	_ = "STUB: not implemented"
	return *new(VMOpt)
}

func WithManualGradient() VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }

func WithEngine(e tensor.Engine) VMOpt { _ = "STUB: not implemented"; return *new(VMOpt) }
