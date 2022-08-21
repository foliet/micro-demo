package price

import (
	"net/http"

	"demo/service/gateway/api/internal/logic/price"
	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SubscribeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubscribeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := price.NewSubscribeLogic(r.Context(), svcCtx)
		err := l.Subscribe(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
