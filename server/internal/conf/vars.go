package conf

import "fmt"

var AppConf *BootConf

//-X 只能注入 包级、未初始化、可导出（首字母大写） 的变量，
var AppName string
var Version string
var Commit string
var bannerfmt = `
%s@%s,commit=%s
author:winlion
email:27115188@qq.com
_               _                                          
_              (_)               | |                                         
| |_ _   _  ____ _ ____   ____  _ | | ____ ____   ____ ____   ____ ___  ____  
|  _) | | |/ ___) |  _ \ / _  |/ || |/ _  |  _ \ / ___) _  ) / ___) _ \|    \ 
| |_| |_| | |   | | | | ( ( | ( (_| ( ( | | | | ( (__( (/ / ( (__| |_| | | | |
\___)____|_|   |_|_| |_|\_|| |\____|\_||_|_| |_|\____)____|_)____)___/|_|_|_|
					   (_____|                                               
图灵互动,为你链接更多价值
`

func Banner() string {
	return fmt.Sprintf(bannerfmt, AppName, Version, Commit)
}
