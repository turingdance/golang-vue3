package rest

import (
	config "turingdance.com/reliable/internal/initial/conf"
)

var appCfg *config.BootConf = config.AppConf

func InitContext(conf *config.BootConf) {
	appCfg = conf
}
