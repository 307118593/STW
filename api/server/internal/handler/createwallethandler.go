package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"wallet/api/server/internal/logic"
	"wallet/api/server/internal/service" // 自定义package
	"wallet/api/server/internal/svc"
	"wallet/api/server/internal/types"
)

func CreateWalletHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateWalletHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateWalletLogic(r.Context(), svcCtx)
		resp, err := l.CreateWallet(&req)
		service.Response(w, resp, err) //②

	}
}
