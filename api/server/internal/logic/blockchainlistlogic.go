package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"wallet/api/server/internal/svc"
)

type BlockChainListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 链列表
func NewBlockChainListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockChainListLogic {
	return &BlockChainListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlockChainListLogic) BlockChainList() (resp []string, err error) {
	// todo: add your logic here and delete this line
	resp = []string{"ETH", "SOL", "BSC"}
	return
}
