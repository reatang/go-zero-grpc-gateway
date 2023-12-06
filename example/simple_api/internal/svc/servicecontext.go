package svc

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_api/internal/config"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/pb/simple_rpc"
	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/simplerpcclient"
	"github.com/reatang/go-zero-grpc-gateway/zgateway"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	SimpleRpcClient simplerpcclient.SimpleRpc

	// 第一步：添加proxy代理对象
	gateway1 *zgateway.RpcGatewayProxy
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 第二步：初始化客户端，一般配置会在yaml中
	simpleRpcConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{
			"127.0.0.1:8080",
		},
		NonBlock: true,
	})

	// 第三步：初始化proxy
	gateway1 := zgateway.NewRpcGatewayProxy("grpc1")
	//gateway.Middlewares() // 支持注册中间件
	gateway1.Register(func(mux *runtime.ServeMux) {
		// 将grpc客户端注入到gateway mux中
		err := simple_rpc.RegisterSimpleRpcHandler(context.Background(), mux, simpleRpcConn.Conn())
		logx.Must(err)
	})

	return &ServiceContext{
		Config: c,

		// 正常业务可用的rpc客户端
		SimpleRpcClient: simplerpcclient.NewSimpleRpc(simpleRpcConn),

		// gateway代理
		gateway1: gateway1,
	}
}

func (sc *ServiceContext) Setup(server *rest.Server) {
	// 第四步：将gateway 注册到 http_server 中
	sc.gateway1.InjectServer(server)
}
