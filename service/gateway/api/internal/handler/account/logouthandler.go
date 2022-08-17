package account

import (
	"demo/service/gateway/api/internal/logic/account"
	"demo/service/gateway/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
