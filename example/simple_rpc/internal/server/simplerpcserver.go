// Code generated by goctl. DO NOT EDIT.
// Source: simple_rpc.proto

package server

import (
	"context"

	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/internal/logic"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/internal/svc"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/pb/simple_rpc"
)

type SimpleRpcServer struct {
	svcCtx *svc.ServiceContext
	simple_rpc.UnimplementedSimpleRpcServer
}

func NewSimpleRpcServer(svcCtx *svc.ServiceContext) *SimpleRpcServer {
	return &SimpleRpcServer{
		svcCtx: svcCtx,
	}
}

// Ping 测试
func (s *SimpleRpcServer) Ping(ctx context.Context, in *simple_rpc.Request) (*simple_rpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}