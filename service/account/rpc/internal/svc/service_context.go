package svc

import (
	"demo/service/account/model/sql"
	"demo/service/account/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel sql.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:    c,
		UserModel: sql.NewUserModel(sqlConn, c.CacheRedis),
	}
}
