package cmd

const PARAM_KEY_NAME = "turing"
const PARAM_KEY_MEMO = "memo"
const PARAM_KEY_APPNAME = "turing"
const PARAM_APP_NAME = "turing"
const PARAM_APP_TTTLE = "图灵互动框架"

// 环境变量
var env string
var signalstr string
var configfile string = "app-prod.yaml"
var rulefile string = ""
var dirsave string = ""

// -source .  -dstpkg turingdance.com/turing/devfast/api/rest/sys -dstfile routefunc.go
// 需要解析router 的
var source string = "."

// 需要的包路径
var dstpkg string = "turingdance.com/turing"

// 生成的router 文件

var dstfile = "router_file.go"

// show version or not
var version = false
