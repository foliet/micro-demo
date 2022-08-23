package svc

import (
	"demo/service/price/model/sql"
	"demo/service/price/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	ItemInfoModel  sql.ItemInfoModel
	SubscribeModel sql.SubscribeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:         c,
		ItemInfoModel:  sql.NewItemInfoModel(sqlConn, c.CacheRedis),
		SubscribeModel: sql.NewSubscribeModel(sqlConn, c.CacheRedis),
	}
}
