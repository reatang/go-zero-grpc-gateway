package logic

import (
	"context"

	"github.com/reatang/go-zero-grpc-gateway/example/simple_api/internal/svc"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Simple_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSimple_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Simple_apiLogic {
	return &Simple_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Simple_apiLogic) Simple_api(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
