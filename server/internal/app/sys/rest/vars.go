package rest

import (
	config "turingdance.com/turing/internal/conf"
)

var appCfg *config.BootConf = config.AppConf

func InitContext(conf *config.BootConf) {
	appCfg = conf
}
