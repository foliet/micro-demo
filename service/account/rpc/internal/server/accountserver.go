// Code generated by goctl. DO NOT EDIT!
// Source: account.proto

package server

import (
	"context"

	"demo/service/account/rpc/internal/logic"
	"demo/service/account/rpc/internal/svc"
	"demo/service/account/rpc/pb/account"
)

type AccountServer struct {
	svcCtx *svc.ServiceContext
	account.UnimplementedAccountServer
}

func NewAccountServer(svcCtx *svc.ServiceContext) *AccountServer {
	return &AccountServer{
		svcCtx: svcCtx,
	}
}

func (s *AccountServer) Login(ctx context.Context, in *account.LoginRequest) (*account.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}