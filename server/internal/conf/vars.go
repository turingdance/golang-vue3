package conf

import "fmt"

var AppConf *BootConf
var Version = "1.0.0"
var Commit = "0.0.0.1"
var bannerfmt = `
_               _                                          
_              (_)               | |                                         
| |_ _   _  ____ _ ____   ____  _ | | ____ ____   ____ ____   ____ ___  ____  
|  _) | | |/ ___) |  _ \ / _  |/ || |/ _  |  _ \ / ___) _  ) / ___) _ \|    \ 
| |_| |_| | |   | | | | ( ( | ( (_| ( ( | | | | ( (__( (/ / ( (__| |_| | | | |
\___)____|_|   |_|_| |_|\_|| |\____|\_||_|_| |_|\____)____|_)____)___/|_|_|_|
					   (_____|                                               
图灵互动开发平台,为你链接更多价值
suport @ www.turingdance.com
version= 1.0.0.%s
author = winlion
`
var Banner = fmt.Sprintf(bannerfmt, Version, Commit)
