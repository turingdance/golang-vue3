package conf

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/turingdance/infra/alikit/osskit"
	"github.com/turingdance/infra/rediskit"
	"turingdance.com/turing/internal/pkg/dysms"
	"turingdance.com/turing/internal/pkg/log"
	"turingdance.com/turing/internal/pkg/storage"
	"turingdance.com/turing/internal/server"
	"turingdance.com/turing/internal/server/auth"
	"turingdance.com/turing/internal/server/middleware"
)

type Engin struct {
	Env     string
	Pidfile string
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
	Engin   Engin
	Http    server.HttpConf
	Log     log.LogConf
	Local   Local
	SysConf DbConf
	Redis   rediskit.RedisConf
	Oss     osskit.OssConf
	Storage []storage.StorageConf
	Dysms   dysms.DysmsConf
	Enpryt  middleware.Enpryt
	Auth    auth.Conf
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
		fmt.Printf("read config failed: %s\n", err.Error())
	}
	AppConf = &c
	return &c
}
