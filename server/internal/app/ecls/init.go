
// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion@turingdance.com

package ecls

import (
	"os"
	"turingdance.com/reliable/internal/initial/conf"
	"turingdance.com/reliable/internal/app/ecls/logic"
	"turingdance.com/reliable/internal/app/ecls/model"
	"turingdance.com/reliable/internal/app/ecls/rest"
	"github.com/turingdance/infra/dbkit"
	"github.com/turingdance/infra/restkit"
)

func Initialize(appconf *conf.BootConf) (err error) {

	// 初始化数据库
	dbconf := appconf.EclsConf
	_, err = logic.InitializeDataBase(dbconf,
		dbkit.WithWriter(os.Stdout),
		dbkit.WithPrefix(dbconf.Prefix),
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
