package price

import (
	"context"
	"demo/service/price/rpc/pb"
	"encoding/json"

	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscribeLogic {
	return &SubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubscribeLogic) Subscribe(req *types.SubscribeRequest) error {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	_, err = l.svcCtx.PriceRpc.Subscribe(l.ctx, &pb.SubscribeRequest{
		UserId: userId,
		ItemId: req.ItemId,
		ShopId: req.ShopId,
	})
	if err != nil {
		return err
	}
	return nil
}
