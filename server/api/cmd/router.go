package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

var dirsrc string = ""
var dirdst string = ""
var author string = ""
var routerfile string = "router.go"
var tplfile string = "./router.tpl"

// 排除在外的
var excludes []string = []string{}

var tplrouter string = `
//don't modify !!!!
// create at ${datetime}
// creeate by ${author}
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
		"{{$v.Module}}":{{$module}}Ctrl, 	// {{$v.Comment}}
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
`

// replace keyword to ruled str
func replace(input string, rule map[string]string) string {
	for k, v := range rule {
		input = strings.ReplaceAll(input, k, v)
	}
	return input
}

type ModuleItem struct {
	Module  string
	Comment string
	Path    string
}

// 生成代码
func _gencode(dirdst string, routes []*Route) (err error) {
	bts, err := os.ReadFile(tplfile)
	if err != nil {
		// 如果是文件不存在问题 则直接
		if os.IsNotExist(err) {
			os.WriteFile(tplfile, []byte(tplrouter), 0755)
		} else {
			return err
		}
	} else {
		tplrouter = string(bts)
	}
	//NodeRoute
	routetree := buildroutetree(routes)
	dirdst, _ = filepath.Abs(dirdst)
	pkg := filepath.Base(dirdst)
	datatime := time.Now().Format("2006-01-02 15:04:05")
	tpl, err := template.New("root").Funcs(template.FuncMap{"join": func(str []string) string {
		return strings.Join(str, ",")
	},
		"json": func(input any) string {
			bts, _ := json.Marshal(input)
			return string(bts)
		},
		"camel": camel,
	}).Parse(replace(tplrouter, map[string]string{
		"${package}":    pkg,
		"${datetime}":   datatime,
		"${author}":     author,
		"${dirdst}":     dirdst,
		"${dirsrc}":     dirsrc,
		"${routerfile}": routerfile,
	}))
	//
	if err != nil {
		return err
	}
	dstfilename := path.Join(dirdst, routerfile)
	_, err = os.Stat(dstfilename)
	// 如果文件不存在,则创建
	if err == nil || os.IsNotExist(err) {
		err = os.Rename(dstfilename, dstfilename+".bak")
	}
	dstfile, err := os.Create(dstfilename)
	if err != nil {
		return err
	}
	dstfile.Truncate(0)
	defer dstfile.Close()
	err = tpl.Execute(dstfile, routetree)
	return err

}

func gencode(dirsrc string, dirdst string, excludes ...string) error {
	excludes = append(excludes, dstfile)
	routes, err := parseroute(dirsrc, excludes...)

	if err != nil {
		return err
	}
	return _gencode(dirdst, routes)
}

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var routerCmd = &cobra.Command{
	Use:   "router", // Use这里定义的就是命令的名称
	Short: "generate route by annotation",
	Long:  `generate route by annotation`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		if dirsrc != "" && dirdst == "" {
			dirdst = dirsrc
		}
		//扫描目录下的每一个文件
		if err := gencode(dirsrc, dirdst, excludes...); err != nil {
			fmt.Println("gen route ", filepath.Join(dirsrc), "->", filepath.Join(dirdst, routerfile), err.Error())
		} else {
			fmt.Println("gen route ", filepath.Join(dirsrc), "->", filepath.Join(dirdst, routerfile), " ok")
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		//这个在命令执行前执行
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		//这个在命令执行后执行
	},
	// 还有其他钩子函数
}

func init() {
	routerCmd.Flags().StringVarP(&dirsrc, "src", "o", "", "dir of source")
	routerCmd.Flags().StringVarP(&dirdst, "dst", "d", "", "dir for save")
	routerCmd.Flags().StringVarP(&tplfile, "tpl", "t", "./router.tpl", "tpl for router")
	routerCmd.Flags().StringVarP(&author, "author", "a", "winlion", "author of code")
	routerCmd.Flags().StringVarP(&routerfile, "name", "n", "router.go", "name of router file")
	routerCmd.Flags().StringArrayVarP(&excludes, "exclude", "x", []string{}, "file will be exclude for scan...")
	rootCmd.AddCommand(routerCmd)
}
