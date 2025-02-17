package logic

import (
	"context"

	"wallet/api/server/internal/svc"
	"wallet/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletTransactionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 交易详情
func NewWalletTransactionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletTransactionDetailLogic {
	return &WalletTransactionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletTransactionDetailLogic) WalletTransactionDetail(req *types.CreateWalletHandlerRequest) (resp *types.WalletInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
