package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"wallet/api/server/internal/logic"
	"wallet/api/server/internal/service" // 自定义package
	"wallet/api/server/internal/svc"
	"wallet/api/server/internal/types"
)

func WalletTransactionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WalletTransactionHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWalletTransactionLogic(r.Context(), svcCtx)
		resp, err := l.WalletTransaction(&req)
		service.Response(w, resp, err) //②

	}
}
