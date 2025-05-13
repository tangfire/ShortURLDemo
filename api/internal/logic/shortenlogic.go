package logic

import (
	"ShortURLDemo/rpc/transform/transformer"
	"context"

	"ShortURLDemo/api/internal/svc"
	"ShortURLDemo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	// 手动代码开始
	rpcResp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return nil, err
	}

	return &types.ShortenResp{
		Shorten: rpcResp.Shorten,
	}, nil
	// 手动代码结束
}
