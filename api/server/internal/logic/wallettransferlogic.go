package logic

import (
	"context"
	"wallet/api/server/internal/service/eth"

	"wallet/api/server/internal/svc"
	"wallet/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletTransferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 转账
func NewWalletTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletTransferLogic {
	return &WalletTransferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletTransferLogic) WalletTransfer(req *types.WalletTransferHandlerRequest) (resp *types.WalletTransferHandlerResponse, err error) {
	//转账
	blockType := req.Blockchain
	switch blockType {
	case "ETH":
		hash := eth.WalletTransfer(req.From, req.PrivateKey, req.To, req.Value, req.Fee)
		resp = &types.WalletTransferHandlerResponse{
			Hash: hash,
		}
		break
	case "SOL":
		break
	case "BSC":
		break

	}

	return
}
