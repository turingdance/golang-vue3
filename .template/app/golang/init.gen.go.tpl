// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author {{.App.Author}}

package {{.App.Name|lower}}

import (
	"os"
	"{{.App.Package}}/internal/conf"
	"{{.App.Package}}/internal/app/{{.App.Name|lower}}/logic"
	"{{.App.Package}}/internal/app/{{.App.Name|lower}}/model"
	"{{.App.Package}}/internal/app/{{.App.Name|lower}}/rest"
	"github.com/turingdance/infra/dbkit"
	"github.com/turingdance/infra/restkit"
)

func Initialize(appconf *conf.BootConf) (err error) {

	// 初始化数据库
	dbconf := appconf.Database
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
