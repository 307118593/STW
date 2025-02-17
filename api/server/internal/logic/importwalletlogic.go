package logic

import (
	"context"
	"wallet/api/server/internal/service/eth"

	"wallet/api/server/internal/svc"
	"wallet/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportWalletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 导入钱包
func NewImportWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportWalletLogic {
	return &ImportWalletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportWalletLogic) ImportWallet(req *types.ImportWalletHandlerRequest) (resp *types.CreateWalletHandlerResponse, err error) {
	// todo: add your logic here and delete this line
	blockType := req.Blockchain
	switch blockType {
	case "ETH":
		address, _ := eth.ImportWallet(req.PrivateKey)
		resp = &types.CreateWalletHandlerResponse{
			Address: address,
		}
		break
	case "SOL":
		break
	case "BSC":
		break

	}
	return
}
