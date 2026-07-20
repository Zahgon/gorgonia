//go:build debug
// +build debug

package cuda

import (
	"log"
	"os"
)

const DEBUG = true

var TABCOUNT uint32
var logger = log.New(os.Stderr, "", 0)
var replacement = "\n"

var (
	allocatorDev = false
)

func tabcount() int { _ = "STUB: not implemented"; return 0 }

func enterLogScope() { _ = "STUB: not implemented"; return }

func leaveLogScope() { _ = "STUB: not implemented"; return }

func logf(format string, others ...interface{}) { _ = "STUB: not implemented"; return }

func allocatorLogf(format string, attrs ...interface{}) { _ = "STUB: not implemented"; return }
