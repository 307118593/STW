package logic

import (
	"context"
	"fmt"
	"wallet/api/server/internal/service/eth"

	"wallet/api/server/internal/svc"
	"wallet/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWalletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建钱包
func NewCreateWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWalletLogic {
	return &CreateWalletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWalletLogic) CreateWallet(req *types.CreateWalletHandlerRequest) (resp *types.CreateWalletHandlerResponse, err error) {
	blockType := req.Blockchain
	switch blockType {
	case "ETH":
		address, private := eth.CreateWallet()
		fmt.Println("钱包:", address, private)
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
