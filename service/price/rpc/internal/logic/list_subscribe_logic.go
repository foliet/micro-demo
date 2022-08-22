package logic

import (
	"context"

	"demo/service/price/rpc/internal/svc"
	"demo/service/price/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSubscribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSubscribeLogic {
	return &ListSubscribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListSubscribeLogic) ListSubscribe(in *pb.ListSubscribeRequest) (*pb.ListSubscribeResponse, error) {
	result, err := l.svcCtx.SubscribeModel.FindAllByUserId(l.ctx, in.Page, in.UserId)
	if err != nil {
		return nil, err
	}
	subscribes := make([]*pb.Subscribe, 0, 4)
	for _, elm := range result {
		subscribes = append(subscribes, &pb.Subscribe{
			UserId: elm.UserId,
			ItemId: elm.ItemId,
			ShopId: elm.ShopId,
		})
	}
	return &pb.ListSubscribeResponse{Subscribes: subscribes}, nil
}
