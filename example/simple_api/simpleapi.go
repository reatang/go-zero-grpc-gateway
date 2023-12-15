package main

import (
	"flag"
	"fmt"

	"github.com/reatang/go-zero-addons/ahttpx"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_api/internal/config"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_api/internal/handler"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_api/internal/svc"
	"github.com/zeromicro/go-zero/rest/router"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/simpleapi-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// TODO 重要!重要!重要! 替换为前缀优先的路由模块
	server := rest.MustNewServer(c.RestConf, rest.WithRouter(ahttpx.NewPrefixPriorityRouter(router.NewRouter())))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 将http服务注入到ServiceContext中
	ctx.Setup(server)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
