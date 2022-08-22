package logic

import (
	"context"

	"demo/service/price/rpc/internal/svc"
	"demo/service/price/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubscribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubscribeLogic {
	return &QuerySubscribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuerySubscribeLogic) QuerySubscribe(in *pb.Subscribe) (*pb.ItemInfos, error) {
	result, err := l.svcCtx.ItemInfoModel.FindAllByUserIdAndItemId(l.ctx, in.UserId, in.ItemId)
	if err != nil {
		return nil, err
	}
	itemInfos := make([]*pb.ItemInfo, 0, 4)
	for _, elm := range result {
		itemInfos = append(itemInfos, &pb.ItemInfo{
			ItemId:   elm.ItemId,
			Price:    elm.Price,
			CreateAt: elm.CreateAt.UnixMilli(),
		})
	}
	return &pb.ItemInfos{ItemInfos: itemInfos}, nil
}
