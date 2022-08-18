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
	client, _ := zrpc.NewClient(c.AccountRpc)
	return &ServiceContext{
		Config:     c,
		AccountRpc: account.NewAccount(client),
	}
}
