package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/turingdance/infra/alikit/osskit"
	"github.com/turingdance/infra/rediskit"
	"turingdance.com/reliable/internal/pkg/dysms"
	"turingdance.com/reliable/internal/pkg/log"
	"turingdance.com/reliable/internal/pkg/storage"
	"turingdance.com/reliable/internal/server"
	"turingdance.com/reliable/internal/server/auth"
)

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

type Engin struct {
	Env     string
	Pidfile string
}
type Enpryt struct {
	Secret string
	Method string
}
type Attach struct {
	Datadir    string
	Filepatern string
}

type Local struct {
	Datadir string
	Alias   string
	Host    string
	Server  string
}

// oss://[bucket]/[filekey]

type BootConf struct {
	Engin    Engin
	Http     server.HttpConf
	Log      log.LogConf
	Local    Local
	SysConf  DbConf
	ClsConf  DbConf
	EclsConf DbConf
	GenConf  DbConf
	Redis    rediskit.RedisConf
	Oss      osskit.OssConf
	Storage  []storage.StorageConf
	Dysms    dysms.DysmsConf
	Enpryt   Enpryt
	Auth     auth.Conf
}

func ParseConf(appfile string) *BootConf {
	viper.SetConfigFile(appfile)
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config failed: " + err.Error())
	}
	viper.WatchConfig()
	var c BootConf
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&c)
		fmt.Println("refresh config")
	})
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Errorf("read config failed: %s\n", err.Error())
	}
	AppConf = &c
	return &c
}
