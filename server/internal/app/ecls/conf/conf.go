
package conf

type ENVDEF string

const PROD ENVDEF = "prod"
const DEV ENVDEF = "dev"

type AppConf struct {
	DbType  string
	Dsn     string
	Env     ENVDEF
	LogFile string
	LogLevel string
}
var DefaultAppConf *AppConf = &AppConf{
	DbType: "sqlite",
	Dsn:     "ecls.db",
	Env:     PROD,
	LogFile: "ecls.log",
	LogLevel: "INFO",
}
