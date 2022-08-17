package svc

import (
	"demo/service/account/rpc/account"
	"demo/service/gateway/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	AccountRpc account.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AccountRpc: account.NewAccount(zrpc.MustNewClient(c.AccountRpc)),
	}
}
