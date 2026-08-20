// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion

package testapp

import (
	"os"
	"turingdance.com/reliable/internal/conf"
	"turingdance.com/reliable/internal/app/testapp/logic"
	"turingdance.com/reliable/internal/app/testapp/model"
	"turingdance.com/reliable/internal/app/testapp/rest"
	"github.com/turingdance/infra/dbkit"
	"github.com/turingdance/infra/restkit"
)

func Initialize(appconf *conf.BootConf) (err error) {

	// 初始化数据库
	dbconf := appconf.TestappConf
	if dbconf.Dsn == "" {
		dbconf = appconf.SysConf
	}
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
	// rest.InitContext(appconf)
	return
}

func CreateRouter(patern string) *restkit.MuxHandler {
	result := restkit.NewMuxHandler()
	result.Router(patern, rest.Routes...)
	return result
}
