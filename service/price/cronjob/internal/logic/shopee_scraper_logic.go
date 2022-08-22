package logic

import (
	"context"
	"demo/service/price/cronjob/internal/svc"
	"demo/service/price/model/sql"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
)

type ShopeeScraperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopeeScarperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopeeScraperLogic {
	return &ShopeeScraperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopeeScraperLogic) Run() {
	result, err := l.svcCtx.SubscribeModel.FindAllWithDistinctItemId(l.ctx)
	if err != nil {
		panic(err)
	}
	scraped := make(map[int64]bool, 16)
	for _, elm := range result {
		if !scraped[elm.ItemId] {
			l.scrape(elm.ShopId, elm.ItemId)
		}
		scraped[elm.ItemId] = true
	}
}

func (l *ShopeeScraperLogic) scrape(shopId, itemId int64) {
	url := fmt.Sprintf("https://shopee.sg/api/v4/item/get?shopid=%d&itemid=%d", shopId, itemId)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	body := &struct {
		Data struct {
			Price float64 `json:"price"`
		} `json:"data"`
		Error    int64  `json:"error,default=0"`
		ErrorMsg string `json:"error_msg"`
	}{}
	err = json.Unmarshal(respBody, body)
	if err != nil {
		panic(err)
	}
	if body.Error != 0 {
		l.Logger.Error("shopee response error: " + body.ErrorMsg)
		return
	}
	_, err = l.svcCtx.ItemInfoModel.Insert(l.ctx, &sql.ItemInfo{
		ItemId: itemId,
		Price:  body.Data.Price / 100000,
	})
	if err != nil {
		panic(err)
	}
}
