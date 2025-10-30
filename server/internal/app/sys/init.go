package sys

import (
	"os"

	"github.com/turingdance/infra/dbkit"
	"github.com/turingdance/infra/restkit"
	"turingdance.com/reliable/internal/app/sys/logic"
	"turingdance.com/reliable/internal/app/sys/model"
	"turingdance.com/reliable/internal/app/sys/rest"
	"turingdance.com/reliable/internal/initial/conf"
)

func Initialize(appconf *conf.BootConf) (err error) {
	// 初始化数据库
	_, err = logic.InitializeDataBase(appconf.SysConf,
		dbkit.WithWriter(os.Stdout),
		dbkit.WithPrefix(appconf.SysConf.Prefix),
		dbkit.IgnoreRecordNotFoundError(true),
		dbkit.ParameterizedQueries(true),
		dbkit.SetLogLevel(appconf.SysConf.Loglevel),
		dbkit.SingularTable(true),
		dbkit.SetConnMaxIdleTime(appconf.ClsConf.ConnMaxIdleTime),
		dbkit.SetMaxOpenConns(appconf.ClsConf.ConnMaxOpen),
		dbkit.SetMaxIdleConns(appconf.ClsConf.ConnMaxIdle),
		dbkit.SetConnMaxLifetime(appconf.ClsConf.ConnMaxLifeTime),
		dbkit.AutoMigrate(model.AllRegistedModel()...),
	)
	// 短信应用
	logic.InitDysmsService(appconf.Dysms)
	rest.InitContext(appconf)
	return
}

func CreateRouter(patern string) *restkit.MuxHandler {
	result := restkit.NewMuxHandler()
	result.Router(patern, rest.Routes...)
	return result
}
