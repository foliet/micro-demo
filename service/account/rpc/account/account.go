// Code generated by goctl. DO NOT EDIT!
// Source: account.proto

package account

import (
	"context"

	"demo/service/account/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginRequest    = pb.LoginRequest
	RegisterRequest = pb.RegisterRequest
	UserId          = pb.UserId

	Account interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserId, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserId, error)
	}

	defaultAccount struct {
		cli zrpc.Client
	}
)

func NewAccount(cli zrpc.Client) Account {
	return &defaultAccount{
		cli: cli,
	}
}

func (m *defaultAccount) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserId, error) {
	client := pb.NewAccountClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultAccount) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserId, error) {
	client := pb.NewAccountClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}
