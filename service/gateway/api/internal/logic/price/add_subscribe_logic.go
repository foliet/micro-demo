package price

import (
	"context"
	"demo/common/errorx"
	"demo/service/price/rpc/price"
	"encoding/json"
	"regexp"
	"strconv"

	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubscribeLogic {
	return &AddSubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSubscribeLogic) AddSubscribe(req *types.AddSubscribeRequest) (resp *types.CodeResponse, err error) {
	resp = new(types.CodeResponse)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	request := &price.AddSubscribeRequest{
		Subscribe: &price.Subscribe{
			UserId: userId,
		},
	}
	request.Subscribe.ItemId, request.Subscribe.ShopId, err = l.parseUrl(req.Url)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.PriceRpc.AddSubscribe(l.ctx, request)
	if err != nil {
		return nil, err
	}
	return
}

func (l *AddSubscribeLogic) parseUrl(url string) (shopId int64, itemId int64, err error) {
	reg, err := regexp.Compile("https://shopee\\.sg/.*-i\\.(\\d+)\\.(\\d+)")
	if err != nil {
		return
	}
	result := reg.FindAllStringSubmatch(url, -1)
	if len(result) != 1 {
		return 0, 0, errorx.ErrWrongUrlFormat
	}
	shopId, err = strconv.ParseInt(result[0][1], 10, 64)
	itemId, err = strconv.ParseInt(result[0][2], 10, 64)
	return
}
