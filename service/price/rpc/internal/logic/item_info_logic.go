package logic

import (
	"context"

	"demo/service/price/rpc/internal/svc"
	"demo/service/price/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ItemInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewItemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ItemInfoLogic {
	return &ItemInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ItemInfoLogic) ItemInfo(in *pb.UserId) (*pb.ItemInfos, error) {
	prices, err := l.svcCtx.ItemInfoModel.FindAllByUserId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	itemInfos := make([]*pb.ItemInfo, 0, 4)
	for _, elm := range prices {
		itemInfos = append(itemInfos, &pb.ItemInfo{
			ItemId:   elm.ItemId,
			Price:    elm.Price,
			CreateAt: elm.CreateAt.UnixMilli(),
		})
	}
	return &pb.ItemInfos{ItemInfos: itemInfos}, nil
}
