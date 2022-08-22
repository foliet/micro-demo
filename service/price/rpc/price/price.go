// Code generated by goctl. DO NOT EDIT!
// Source: price.proto

package price

import (
	"context"

	"demo/service/price/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Empty      = pb.Empty
	ItemInfo   = pb.ItemInfo
	ItemInfos  = pb.ItemInfos
	Subscribe  = pb.Subscribe
	Subscribes = pb.Subscribes
	UserId     = pb.UserId

	Price interface {
		AddSubscribe(ctx context.Context, in *Subscribe, opts ...grpc.CallOption) (*Empty, error)
		ListSubscribe(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Subscribes, error)
		QuerySubscribe(ctx context.Context, in *Subscribe, opts ...grpc.CallOption) (*ItemInfos, error)
	}

	defaultPrice struct {
		cli zrpc.Client
	}
)

func NewPrice(cli zrpc.Client) Price {
	return &defaultPrice{
		cli: cli,
	}
}

func (m *defaultPrice) AddSubscribe(ctx context.Context, in *Subscribe, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewPriceClient(m.cli.Conn())
	return client.AddSubscribe(ctx, in, opts...)
}

func (m *defaultPrice) ListSubscribe(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Subscribes, error) {
	client := pb.NewPriceClient(m.cli.Conn())
	return client.ListSubscribe(ctx, in, opts...)
}

func (m *defaultPrice) QuerySubscribe(ctx context.Context, in *Subscribe, opts ...grpc.CallOption) (*ItemInfos, error) {
	client := pb.NewPriceClient(m.cli.Conn())
	return client.QuerySubscribe(ctx, in, opts...)
}
