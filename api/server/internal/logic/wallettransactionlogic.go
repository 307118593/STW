package logic

import (
	"context"

	"wallet/api/server/internal/svc"
	"wallet/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletTransactionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询钱包交易记录
func NewWalletTransactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletTransactionLogic {
	return &WalletTransactionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletTransactionLogic) WalletTransaction(req *types.WalletTransactionHandlerRequest) (resp *types.WalletTransactionHandlerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
