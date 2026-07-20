//go:build !debug
// +build !debug

package gorgonia

const DEBUG = false

const (
	compileDev        = false
	shapeInferenceDev = false
	typeSystemDev     = false
	symdiffDev        = false
	autodiffDev       = false
	machineDev        = false
	stabilizationDev  = false
	cudaDev           = false
	allocatorDev      = false
)

func tabcount() int { _ = "STUB: not implemented"; return 0 }

func enterLogScope()                                   { _ = "STUB: not implemented"; return }
func leaveLogScope()                                   { _ = "STUB: not implemented"; return }
func logf(format string, others ...interface{})        { _ = "STUB: not implemented"; return }
func compileLogf(format string, attrs ...interface{})  { _ = "STUB: not implemented"; return }
func shapeLogf(format string, attrs ...interface{})    { _ = "STUB: not implemented"; return }
func typeSysLogf(format string, attrs ...interface{})  { _ = "STUB: not implemented"; return }
func symdiffLogf(format string, attrs ...interface{})  { _ = "STUB: not implemented"; return }
func autodiffLogf(format string, attrs ...interface{}) { _ = "STUB: not implemented"; return }
func machineLogf(format string, attrs ...interface{})  { _ = "STUB: not implemented"; return }
func stabLogf(format string, attrs ...interface{})     { _ = "STUB: not implemented"; return }
func solverLogf(format string, attrs ...interface{})   { _ = "STUB: not implemented"; return }
func cudaLogf(format string, attrs ...interface{})     { _ = "STUB: not implemented"; return }
func allocatorLogf(format string, attr ...interface{}) { _ = "STUB: not implemented"; return }
func recoverFrom(format string, attrs ...interface{})  { _ = "STUB: not implemented"; return }

func GraphCollisionStats() (int, int, int) { _ = "STUB: not implemented"; return 0, 0, 0 }

func incrCC() { _ = "STUB: not implemented"; return }
func incrEC() { _ = "STUB: not implemented"; return }
func incrNN() { _ = "STUB: not implemented"; return }

func logCompileState(name string, g *ExprGraph, df *dataflow) { _ = "STUB: not implemented"; return }

func (df *dataflow) debugIntervals(sorted Nodes) { _ = "STUB: not implemented"; return }
