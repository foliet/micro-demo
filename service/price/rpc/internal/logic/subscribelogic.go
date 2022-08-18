package logic

import (
	"context"

	"demo/service/price/rpc/internal/svc"
	"demo/service/price/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubscribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscribeLogic {
	return &SubscribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubscribeLogic) Subscribe(in *pb.SubscribeRequest) (*pb.Empty, error) {
	// todo: add your logic here and delete this line

	return &pb.Empty{}, nil
}
