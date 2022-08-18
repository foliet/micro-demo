package svc

import (
	"demo/service/account/rpc/account"
	"demo/service/gateway/api/internal/config"
	"demo/service/price/rpc/price"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	AccountRpc account.Account
	PriceRpc   price.Price
}

func NewServiceContext(c config.Config) *ServiceContext {
	accountClient, _ := zrpc.NewClient(c.AccountRpc)
	priceClient, _ := zrpc.NewClient(c.PriceRpc)
	return &ServiceContext{
		Config:     c,
		AccountRpc: account.NewAccount(accountClient),
		PriceRpc:   price.NewPrice(priceClient),
	}
}
