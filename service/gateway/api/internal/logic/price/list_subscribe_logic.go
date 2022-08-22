package price

import (
	"context"
	"demo/service/price/rpc/price"
	"encoding/json"

	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSubscribeLogic {
	return &ListSubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSubscribeLogic) ListSubscribe() (resp *types.ListSubscribeResponse, err error) {
	resp = new(types.ListSubscribeResponse)
	resp.Subscribes = make([]*types.Subscribe, 0, 4)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	result, err := l.svcCtx.PriceRpc.ListSubscribe(l.ctx, &price.UserId{Id: userId})
	if err != nil {
		return nil, err
	}
	for _, elm := range result.Subscribes {
		resp.Subscribes = append(resp.Subscribes, &types.Subscribe{
			ItemId: elm.ItemId,
			ShopId: elm.ShopId,
		})
	}
	return
}
