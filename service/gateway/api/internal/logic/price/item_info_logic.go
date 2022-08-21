package price

import (
	"context"
	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"
	"demo/service/price/rpc/price"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type ItemInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewItemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ItemInfoLogic {
	return &ItemInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ItemInfoLogic) ItemInfo() (resp *types.ItemInfos, err error) {
	resp = new(types.ItemInfos)
	resp.ItemInfos = make([]*types.ItemInfo, 0, 4)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	itemInfos, err := l.svcCtx.PriceRpc.ItemInfo(l.ctx, &price.UserId{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	for _, itemInfo := range itemInfos.ItemInfos {
		resp.ItemInfos = append(resp.ItemInfos, &types.ItemInfo{
			ItemId:   itemInfo.ItemId,
			Price:    itemInfo.Price,
			CreateAt: itemInfo.CreateAt,
		})
	}
	return
}
