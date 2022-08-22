package price

import (
	"context"
	"demo/service/price/rpc/price"
	"encoding/json"

	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubscribeLogic {
	return &QuerySubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySubscribeLogic) QuerySubscribe(req *types.QuerySubscribeRequest) (resp *types.QuerySubscribeResponse, err error) {
	resp = new(types.QuerySubscribeResponse)
	resp.ItemInfos = make([]*types.ItemInfo, 0, 4)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	result, err := l.svcCtx.PriceRpc.QuerySubscribe(l.ctx, &price.QuerySubscribeRequest{
		Subscribe: &price.Subscribe{
			UserId: userId,
			ItemId: req.ItemId,
			ShopId: req.ShopId,
		}})
	if err != nil {
		return nil, err
	}
	for _, itemInfo := range result.ItemInfos {
		resp.ItemInfos = append(resp.ItemInfos, &types.ItemInfo{
			ItemId:   itemInfo.ItemId,
			Price:    itemInfo.Price,
			CreateAt: itemInfo.CreateAt,
		})
	}
	return
}
