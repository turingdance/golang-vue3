package cache

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/turingdance/infra/cachekit"
	"github.com/turingdance/infra/rediskit"
)

var CacheEngin cachekit.Cache

func Initialize(cacheConf rediskit.RedisConf) {
	if cacheConf.Addr != "" {
		CacheEngin = cachekit.NewRedisCache(&cacheConf)
	} else {
		CacheEngin = cachekit.NewMemoryCache(time.Second * 60)
	}
	log.Infof("InitializeCache %s")
}
