package main

import (
	"go/ast"
)

type nametypePair struct {
	name string
	*ast.FuncType
}

func functions(decls []ast.Decl) (signatures []nametypePair) { _ = "STUB: not implemented"; return nil }

type strRepr struct {
	name     string
	inTypes  []string
	retTypes []string

	printName bool
}

func (s strRepr) String() string { _ = "STUB: not implemented"; return "" }

func processSig(pair nametypePair) strRepr { _ = "STUB: not implemented"; return *new(strRepr) }

func parseTypeExpr(expr ast.Expr) string { _ = "STUB: not implemented"; return "" }

func filterSigs(xs []strRepr, fn func(strRepr) bool) (retVal []strRepr) {
	_ = "STUB: not implemented"
	return nil
}

func functionSignatures() { _ = "STUB: not implemented"; return }
