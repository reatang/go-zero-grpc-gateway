package main

import (
	"flag"
	"fmt"

	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/internal/config"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/internal/server"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/internal/svc"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/pb/simple_rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/simplerpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		simple_rpc.RegisterSimpleRpcServer(grpcServer, server.NewSimpleRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
