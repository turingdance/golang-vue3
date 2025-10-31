package sys

import (
	"os"

	"github.com/turingdance/infra/dbkit"
	"github.com/turingdance/infra/restkit"
	"turingdance.com/turing/internal/app/sys/logic"
	"turingdance.com/turing/internal/app/sys/model"
	"turingdance.com/turing/internal/app/sys/rest"
	"turingdance.com/turing/internal/conf"
)

func Initialize(appconf *conf.BootConf) (err error) {
	dbconf := appconf.SysConf
	// 初始化数据库
	_, err = logic.InitializeDataBase(dbconf,
		dbkit.WithWriter(os.Stdout),
		dbkit.IgnoreRecordNotFoundError(true),
		dbkit.ParameterizedQueries(true),
		dbkit.SetLogLevel(dbconf.Loglevel),
		dbkit.SingularTable(true),
		dbkit.SetConnMaxIdleTime(dbconf.ConnMaxIdleTime),
		dbkit.SetMaxOpenConns(dbconf.ConnMaxOpen),
		dbkit.SetMaxIdleConns(dbconf.ConnMaxIdle),
		dbkit.SetConnMaxLifetime(dbconf.ConnMaxLifeTime),
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
