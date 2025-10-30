package conf

import "time"

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
