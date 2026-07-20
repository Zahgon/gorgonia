package gorgonia

import (
	"fmt"
	"reflect"
)

type mapFmt struct {
	m reflect.Value
}

func FmtNodeMap(m interface{}) mapFmt { _ = "STUB: not implemented"; return *new(mapFmt) }

func (mf mapFmt) defaultFmt(s fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (mf mapFmt) format(s fmt.State, c rune) string { _ = "STUB: not implemented"; return "" }

func (mf mapFmt) Format(s fmt.State, c rune) { _ = "STUB: not implemented"; return }
