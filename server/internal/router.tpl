
//don't modify !!!!
// create at ${datetime}
// creeate by ${author}
//go:generate  codectl router -a ${author} -s . -d . -n ${routerfile}
package ${package}

import (
	"net/http"
	"github.com/turingdance/infra/restkit"
)
type Route struct {
	Package string
	Module  string
	Func    string
	Path    string
	Method  []string
	Comment string
	HandlerFunc http.HandlerFunc
}

var (
{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
	// {{$v.Comment}}
	{{$module}}Ctrl = &{{if ne $v.Package "${package}" }}{{$v.Package}}.{{end}}{{$v.Module}}{}
{{end}}
)
var MapCtrl map[string]any = map[string]any{
	{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
	// {{$v.Comment}}
		"{{$v.Module}}":{{$module}}Ctrl,
	{{end}}
}
var Routes []Route= []Route{
	{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
		{{- range $m,$n := $v.Children }}
		{Package:"{{$n.Package}}" ,Module:"{{$n.Module}}",HandlerFunc:{{$module}}Ctrl.{{$n.Func}},Func:"{{$n.Func}}",Path: "{{$n.Path}}",Method:[]string{	{{- range $x,$y := $n.Method }}"{{$y}}",{{end}} },Comment:"{{$n.Comment}}"},
		{{end}}
	{{end}}
}

var DefaultRouter *restkit.Router = restkit.NewRouter().PathPrefix("/")
// 初始化路由
func InitRouter(router *restkit.Router) {
	{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
	// {{$v.Comment}}
	{{$module}}Ctrl := &{{$v.Module}}{}
	{{$module}}router := router.Subrouter().PathPrefix("{{$v.Path}}")
	{{- range $g,$h := $v.Children }}
	//{{$h.Comment}}
	{{$module}}router.HandleFunc("{{$h.Path}}", {{$module}}Ctrl.{{$h.Func}}).Methods({{range $h.Method}}"{{.}}",{{end}})
	{{end}}
	{{end}}
}
func init() {
	InitRouter(DefaultRouter)
}
