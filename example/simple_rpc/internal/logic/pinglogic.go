package logic

import (
	"context"
	"fmt"

	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/internal/svc"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/pb/simple_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Ping 测试
func (l *PingLogic) Ping(in *simple_rpc.Request) (*simple_rpc.Response, error) {
	return &simple_rpc.Response{
		Pong: fmt.Sprintf("simple_rpc say: %s", in.Ping),
	}, nil
}
