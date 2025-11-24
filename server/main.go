package main

import (
	"flag"

	"turingdance.com/turing/api"
)

// @title 环境上报系统
// @version 1.0
// @description golang 编程
//
//go:generate  codectl router  -s ./internal/app/$(app)/rest
var (
	version   bool
	signalstr string
)

func main() {
	flag.StringVar(&api.AppName, "n", "turing", "应用名称,安装|卸载服务时需要")
	flag.StringVar(&api.AppMemo, "m", "turing", "应用描述,安装服务时需要")
	flag.StringVar(&api.AppConfig, "c", "app-prod.yaml", "配置文件路径")
	flag.BoolVar(&version, "v", false, "应用程序版本")
	flag.StringVar(&signalstr, "s", "start", "系统运行指令,取值如下:\nstart:启动应用 stop:结束应用 install:安装服务 uninstall:卸载服务\n")
	flag.Parse()
	// 如果为0
	app := api.NewApp(api.AppName)
	if api.AppMemo != "" {
		app.WithOption(api.Description(api.AppMemo))
	}
	if api.AppConfig != "" {
		app.WithOption(api.Config(api.AppConfig))
	}
	if version {
		signalstr = "version"
	}
	app.DispatchCmd(signalstr)

}
