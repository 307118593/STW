package logic

import (
	"context"
	"wallet/api/server/internal/service/eth"
	"wallet/api/server/internal/svc"
	"wallet/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询钱包信息
func NewWalletInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletInfoLogic {
	return &WalletInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletInfoLogic) WalletInfo(req *types.WalletInfoRequest) (resp *types.WalletInfoResponse, err error) {
	blockType := req.Blockchain
	switch blockType {
	case "ETH":
		address, balance, defi := eth.WalletInfo(req.Address)
		resp = &types.WalletInfoResponse{
			Address: address,
			Balance: balance,
			Defi:    defi,
		}
		break
	case "SOL":
		break
	case "BSC":
		break
	}
	return
}
