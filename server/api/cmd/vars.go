package cmd

const PARAM_KEY_NAME = "envdocs"
const PARAM_KEY_MEMO = "memo"
const PARAM_KEY_APPNAME = "envdocs"

const PARAM_APP_NAME = "envdocs"
const PARAM_APP_TTTLE = "环境档案"

// 环境变量
var env string
var signalstr string
var configfile string = "app-prod.yml"
var rulefile string = ""
var dirsave string = ""

// -source .  -dstpkg turingdance.com/turing/devfast/api/rest/sys -dstfile routefunc.go
// 需要解析router 的
var source string = "."

// 需要的包路径
var dstpkg string = "turingdance.com/reliable"

// 生成的router 文件

var dstfile = "router_file.go"
