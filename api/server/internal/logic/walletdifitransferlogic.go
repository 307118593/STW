package logic

import (
	"context"
	"wallet/api/server/internal/service/eth"

	"wallet/api/server/internal/svc"
	"wallet/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletDefiTransferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 代币转账
func NewWalletDefiTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletDefiTransferLogic {
	return &WalletDefiTransferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletDefiTransferLogic) WalletDefiTransfer(req *types.WalletDefiTransferHandlerRequest) (resp *types.WalletDefiTransferHandlerResponse, err error) {
	// todo: add your logic here and delete this line
	blockType := req.Blockchain
	switch blockType {
	case "ETH":
		hash := eth.DefiTransfer(req.From, req.PrivateKey, req.To, req.Value, req.Fee, req.Contract)
		resp = &types.WalletDefiTransferHandlerResponse{
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
