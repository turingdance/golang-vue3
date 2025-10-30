package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/turingdance/infra/slicekit"
	"github.com/turingdance/infra/stringx"
)

type Route struct {
	Package string
	Module  string
	Func    string
	Path    string
	Method  []string
	Comment string
}

func (p *Route) Println() {
	fmt.Printf("moduel=%s,func=%s,path=%s,method=%s,comment=%s\n", p.Module, p.Func, p.Path, p.Method, p.Comment)
}

type RouteTree struct {
	Route
	Children []*Route
}

// func (ctrl *Account ) Register (
var regfunc *regexp.Regexp = regexp.MustCompile(`\s*func\s+\(\s*\w+\s+\*([A-Z]+\w+)\s*\)\s+([A-Z]+\w+)\s*\(.*`)

// @Summary 注册用户
var regsummary *regexp.Regexp = regexp.MustCompile(`\s*@Summary\s+(.*)`)

// @Router /account/register [POST,GET]
var regrouter *regexp.Regexp = regexp.MustCompile(`\s*@Router\s+(\S+)\s+(\[?.*\]?)`)

// type Account struct{}
var regstruct *regexp.Regexp = regexp.MustCompile(`\s*type\s+(\S+)\s+struct\s*\{.*`)

// package Account
var regpackage *regexp.Regexp = regexp.MustCompile(`\s*package\s+(\S+)`)

// // 简单描述
var regcomment *regexp.Regexp = regexp.MustCompile(`.*\/\/\s*(.*).*`)

// 下划线单词转为小写驼峰单词
func camel(s string) string {
	return stringx.CamelLcFirst(s)
}
func newroute(pkg string) *Route {
	return &Route{
		Package: pkg,
		Method:  []string{"GET", "POST", "OPTIONS"},
	}
}

// 解析所有文件,构建信息结构体
func parseroute(dirsrc string, excludes ...string) (routes []*Route, err error) {
	if dirsrc == "" {
		err = errors.New("请指定接口代码路径")
		return
	}
	routes = make([]*Route, 0)
	err = filepath.WalkDir(dirsrc, func(fpath string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}
		//basepath := strings.ReplaceAll(fpath, dirsrc, "")
		// 过滤掉需要隔离的文件
		if slicekit.Some(excludes, func(arr []string, ele string) bool {
			if strings.Contains(fpath, ele) {
				return true
			} else {
				return false
			}
		}) {
			return nil
		}
		bts, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(bytes.NewReader(bts))
		var proute *Route = nil
		pkg := ""
		_ = pkg
		for scanner.Scan() {
			txt := scanner.Text()
			// 开始
			if regpackage.MatchString(txt) { //package
				arr := regpackage.FindStringSubmatch(txt)
				pkg = arr[1]
				proute = newroute(pkg)
			} else if regstruct.MatchString(txt) { // 结构体
				structs := regstruct.FindStringSubmatch(txt)
				proute.Module = structs[1]
				routes = append(routes, proute)
				// 开启新的结构体
				proute = newroute(pkg)
			} else if regfunc.MatchString(txt) { // 碰到函数
				funcs := regfunc.FindStringSubmatch(txt)
				proute.Module = funcs[1]
				proute.Func = funcs[2]
				// 如果是空的
				if proute.Path == "" {
					proute.Path = path.Join("//", camel(proute.Module), camel(proute.Func))
				}
				routes = append(routes, proute)
				// 开启新的结构体
				proute = newroute(pkg)
			} else if regsummary.MatchString(txt) { // 处理@Summary 函数
				arr := regsummary.FindStringSubmatch(txt)
				proute.Comment = arr[1]
			} else if regrouter.MatchString(txt) { // 处理@Router /a/b/c [POST,DELETE]
				arr := regrouter.FindStringSubmatch(txt)
				proute.Path = arr[1]
				if strings.Contains(arr[2], "[") {
					arr[2] = strings.TrimPrefix(arr[2], "[")
					arr[2] = strings.TrimSuffix(arr[2], "]")
					arr[2] = strings.Trim(arr[2], " ")
				}
				if arr[2] != "" {
					proute.Method = strings.Split(arr[2], ",")
				}
			} else if regcomment.MatchString(txt) { //处理注释
				arr := regcomment.FindStringSubmatch(txt)
				if proute != nil && proute.Comment == "" {
					proute.Comment = arr[1]
				}
			} else {

			}
		}
		return err
	})

	return
}

// 构建树结构
func buildroutetree(routes []*Route) (tree []*RouteTree) {
	routetree := make([]*RouteTree, 0)
	// 这里是处理model
	for _, v := range routes {
		if v.Func == "" && v.Module != "" {
			routetree = append(routetree, &RouteTree{
				Route:    *v,
				Children: make([]*Route, 0),
			})
		}
	}

	for _, v := range routes {
		if v.Func != "" {
			for _, node := range routetree {
				if node.Module == v.Module {
					node.Children = append(node.Children, v)
				}
			}
		}
	}
	for _, v := range routetree {
		slicekit.SortStable(v.Children, func(e1, e2 *Route) bool {
			score1 := score(e1.Path)
			score2 := score(e2.Path)
			return score1 > score2
		})
	}
	return routetree
}

// 为路径打分,mark priay for rule
func score(path string) int {
	var scorechar string = "{:"
	var ret int = 0
	for _, v := range scorechar {
		if strings.Contains(path, string(v)) {
			ret += 1
		}
	}
	return ret
}
