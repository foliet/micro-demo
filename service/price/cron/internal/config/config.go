package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	service.ServiceConf
	Mysql struct {
		Datasource string
	}
	CacheRedis cache.CacheConf
}
