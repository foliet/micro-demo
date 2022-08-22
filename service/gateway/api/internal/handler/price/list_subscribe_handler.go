package price

import (
	"net/http"

	"demo/service/gateway/api/internal/logic/price"
	"demo/service/gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListSubscribeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := price.NewListSubscribeLogic(r.Context(), svcCtx)
		resp, err := l.ListSubscribe()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
