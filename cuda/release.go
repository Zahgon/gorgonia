//go:build !debug
// +build !debug

package cuda

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", 0)
var replacement = "\n"

func tabcount() int { _ = "STUB: not implemented"; return 0 }

func enterLogScope() { _ = "STUB: not implemented"; return }

func leaveLogScope() { _ = "STUB: not implemented"; return }

func logf(format string, others ...interface{}) { _ = "STUB: not implemented"; return }

func allocatorLogf(format string, attrs ...interface{}) { _ = "STUB: not implemented"; return }
