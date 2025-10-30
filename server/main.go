package main

import (
	"turingdance.com/reliable/api/cmd"
)

// @title 环境上报系统
// @version 1.0
// @description golang 编程
//
//go:generate  codectl router  -a winlion -s ./internal/app/$(app)/rest
func main() {
	cmd.Execute()
}
