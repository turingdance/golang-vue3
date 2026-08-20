// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author {{.App.Author}}

package conf

import (
	"fmt"
	"time"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/turingdance/infra/alikit/osskit"
	"github.com/turingdance/infra/rediskit"
	"{{.App.Package}}/internal/pkg/dysms"
	"{{.App.Package}}/internal/pkg/log"
	"{{.App.Package}}/internal/pkg/storage"
	"{{.App.Package}}/internal/server"
	"{{.App.Package}}/internal/server/auth"
	"{{.App.Package}}/internal/server/middleware"

)

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

type DbConf struct {
	Dsn             string
	Prefix          string
	Engin           string
	Charset         string
	Loglevel        string
	ConnMaxIdleTime time.Duration
	ConnMaxLifeTime time.Duration
	ConnMaxOpen     int
	ConnMaxIdle     int
}


// oss://[bucket]/[filekey]
type BootConf struct {
	Engin        Engin
	Http         server.HttpConf
	Log          log.LogConf
	Local        Local
	SysConf      DbConf
	{{.App.Name|ucfirst}}Conf DbConf
	Redis        rediskit.RedisConf
	Oss          osskit.OssConf
	Storage      []storage.StorageConf
	Dysms        dysms.DysmsConf
	Enpryt       middleware.Enpryt
	Auth         auth.Conf
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
		fmt.Println("read config failed: ", err.Error())
	}
	AppConf = &c
	return &c
}

func Reload(){
	 viper.Unmarshal(AppConf)
}
