package gorgonia

import (
	"text/template"
)

const exprNodeTemplText = `<
<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0" PORT="anchor" {{if isLeaf .}} COLOR="#00FF00;"{{else if isRoot . }} COLOR="#FF0000;" {{else if isMarked .}} COLOR="#0000FF;" {{end}}{{if isInput .}} BGCOLOR="lightyellow"{{else if isStmt .}} BGCOLOR="lightblue"{{end}}>

<TR><TD>{{printf "%x" .ID}}</TD><TD>{{printf "%v" .Name | html | dotEscape}} :: {{nodeType . | html | dotEscape }}</TD></TR>
{{if printOp . }}<TR><TD>Op</TD><TD>{{ opStr . | html | dotEscape }} :: {{ opType . | html | dotEscape }}</TD></TR>{{end}}
{{if hasShape .}}<TR><TD>Shape</TD><TD>{{ getShape .}}</TD></TR>{{end}}
<TR><TD>Overwrites Input {{overwritesInput . }}</TD><TD>Data On: {{.Device}}</TD></TR>
{{if hasGrad .}}<TR><TD>Value</TD><TD>Grad</TD></TR>
<TR><TD>{{printf "%+3.3s" .Value | dotEscape}}</TD><TD>{{getGrad . | dotEscape }} </TD></TR>
<TR><TD>Ptr: {{getValPtr . | dotEscape}} </TD><TD>Ptr: {{getGradPtr . | dotEscape}} </TD></TR>
{{else}}
<TR><TD>Value</TD><TD>{{printf "%+3.3s" .Value | dotEscape}}</TD></TR>
{{end}}

</TABLE>
>`

func dotEscape(s string) string { _ = "STUB: not implemented"; return "" }

func printOp(n *Node) bool  { _ = "STUB: not implemented"; return false }
func isLeaf(n *Node) bool   { _ = "STUB: not implemented"; return false }
func isInput(n *Node) bool  { _ = "STUB: not implemented"; return false }
func isMarked(n *Node) bool { _ = "STUB: not implemented"; return false }
func isRoot(n *Node) bool   { _ = "STUB: not implemented"; return false }
func isStmt(n *Node) bool   { _ = "STUB: not implemented"; return false }
func hasShape(n *Node) bool { _ = "STUB: not implemented"; return false }
func hasGrad(n *Node) bool  { _ = "STUB: not implemented"; return false }
func opStr(n *Node) string  { _ = "STUB: not implemented"; return "" }
func opType(n *Node) string { _ = "STUB: not implemented"; return "" }

func nodeType(n *Node) string { _ = "STUB: not implemented"; return "" }

func overwritesInput(n *Node) int { _ = "STUB: not implemented"; return 0 }

func getShape(n *Node) string { _ = "STUB: not implemented"; return "" }

func getGrad(n *Node) string { _ = "STUB: not implemented"; return "" }

func getGradPtr(n *Node) string { _ = "STUB: not implemented"; return "" }

func getValPtr(n *Node) string { _ = "STUB: not implemented"; return "" }

var funcMap = template.FuncMap{
	"dotEscape":       dotEscape,
	"printOp":         printOp,
	"isRoot":          isRoot,
	"isLeaf":          isLeaf,
	"isInput":         isInput,
	"isStmt":          isStmt,
	"isMarked":        isMarked,
	"hasShape":        hasShape,
	"hasGrad":         hasGrad,
	"getShape":        getShape,
	"getValPtr":       getValPtr,
	"getGrad":         getGrad,
	"getGradPtr":      getGradPtr,
	"overwritesInput": overwritesInput,
	"opStr":           opStr,
	"opType":          opType,
	"nodeType":        nodeType,
}

var (
	exprNodeTempl     *template.Template
	exprNodeJSONTempl *template.Template
)

func init() {
	exprNodeTempl = template.Must(template.New("node").Funcs(funcMap).Parse(exprNodeTemplText))
}
