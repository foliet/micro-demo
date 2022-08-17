package svc

import (
	"demo/service/account/rpc/accountclient"
	"demo/service/gateway/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	AccountRpc accountclient.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AccountRpc: accountclient.NewAccount(zrpc.MustNewClient(c.AccountRpc)),
	}
}
