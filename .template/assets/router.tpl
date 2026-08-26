
//don't modify !!!! creeate by ${author} ,at ${datetime}
//go:generate  codectl router -a ${author} -s . -d . -n ${routerfile}
package ${package}

import (
	"github.com/turingdance/infra/restkit"
)
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
var Routes []restkit.Route= []restkit.Route{
	{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
		{{- range $m,$n := $v.Children }}
		{Package:"{{$n.Package}}" ,Module:"{{$n.Module}}",HandlerFunc:{{$module}}Ctrl.{{$n.Func}},Func:"{{$n.Func}}",Path: "{{$n.Path}}",Method:[]string{	{{- range $x,$y := $n.Method }}"{{$y}}",{{end}} },Comment:"{{$n.Comment}}"},
		{{end}}
	{{end}}
}