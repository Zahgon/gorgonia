package main

import (
	"io"
	"text/template"
)

type UnaryOpInterfaceData struct {
	OpTypes []string
	Dtype   string
}

const unaryOpInterfaceRaw = `func (f *s{{.Dtype}}UnaryOperator) unaryOpType() ʘUnaryOperatorType {
	{{$dt := .Dtype -}}
	switch f {
		{{range $i, $op := .OpTypes -}}
		case &{{$op}}{{$dt}}:
			return {{$op}}OpType
		{{end -}}
	}
	return maxʘUnaryOperator
}

func (f *s{{.Dtype}}UnaryOperator) String() string { return f.unaryOpType().String() }

`

var unaryOpInterface *template.Template

func init() {
	unaryOpInterface = template.Must(template.New("UnOpInterface").Funcs(funcmap).Parse(unaryOpInterfaceRaw))
}

func generateUnaryInterface(outFile io.Writer) { _ = "STUB: not implemented"; return }
