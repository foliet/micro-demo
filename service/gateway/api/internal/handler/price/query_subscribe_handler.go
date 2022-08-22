package price

import (
	"net/http"

	"demo/service/gateway/api/internal/logic/price"
	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QuerySubscribeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuerySubscribeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := price.NewQuerySubscribeLogic(r.Context(), svcCtx)
		resp, err := l.QuerySubscribe(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
