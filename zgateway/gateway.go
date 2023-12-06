package zgateway

import (
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zeromicro/go-zero/rest"
)

type RpcGatewayProxy struct {
	f func(mux *runtime.ServeMux)
	m []rest.Middleware

	prefix string
}

func NewRpcGatewayRegister(prefix string) *RpcGatewayProxy {
	return &RpcGatewayProxy{prefix: prefix, m: make([]rest.Middleware, 0)}
}

// Middlewares 注册中间件
func (r *RpcGatewayProxy) Middlewares(m ...rest.Middleware) {
	r.m = m
}

// Register 注册RPC gateway
func (r *RpcGatewayProxy) Register(f func(mux *runtime.ServeMux)) {
	r.f = f
}

// InjectServer 将 rpc支持gateway的接口注册到 rest服务中，只允许使用一次
func (r *RpcGatewayProxy) InjectServer(s *rest.Server) {
	if r.f == nil {
		return
	}

	mux := runtime.NewServeMux()

	r.f(mux)

	s.AddRoutes(
		rest.WithMiddlewares(
			r.m,
			rest.Route{
				Method:  http.MethodPost,
				Path:    fmt.Sprintf("/%s/", r.prefix), // 需要支持前缀路由
				Handler: http.StripPrefix(fmt.Sprintf("/%s", r.prefix), mux).ServeHTTP,
			},
		),
	)
}
