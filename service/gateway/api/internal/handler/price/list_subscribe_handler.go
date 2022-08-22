package price

import (
	"net/http"

	"demo/service/gateway/api/internal/logic/price"
	"demo/service/gateway/api/internal/svc"
	"demo/service/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListSubscribeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListSubscribeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := price.NewListSubscribeLogic(r.Context(), svcCtx)
		resp, err := l.ListSubscribe(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
