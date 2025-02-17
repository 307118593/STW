package handler

import (
	"net/http"
	"wallet/api/server/internal/logic"
	"wallet/api/server/internal/service" // 自定义package
	"wallet/api/server/internal/svc"
)

func BlockChainListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewBlockChainListLogic(r.Context(), svcCtx)
		resp, err := l.BlockChainList()
		service.Response(w, resp, err) //②

	}
}
