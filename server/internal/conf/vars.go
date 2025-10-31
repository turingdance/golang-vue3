package conf

import "fmt"

var AppConf *BootConf

var Version = "0.0.0.1"
var Commit = "0.0.0.1"
var bannerfmt = `
rpcapi @%s-%s
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
var Banner = fmt.Sprintf("%s %s", Version, Commit)
