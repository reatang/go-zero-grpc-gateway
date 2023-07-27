# go-zero-grpc-gateway

go-zero rest框架中融合grpc-gateway

调用链： http client -> ( go-zero-restapi -> grpc-gateway -> grpc-client ) -> grpc-server

## 一、安装

### 1、下载 googleapis 仓库
```shell
git clone https://github.com/googleapis/googleapis ~/protoc/googleapis
```

### 2、下载对应操作系统版本的 protobuf 描述文件
```
https://github.com/protocolbuffers/protobuf/releases
```

解压到 `~/protoc` 文件夹中，如：`~/protoc/protoc-23.2-osx-aarch_64`

### 3、 golang使用的 protobuf 代码生成器
```shell
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## 二、搭建环境

见 `exapmle` 目录中的示例，关键代码文件：

**rpc部分**
- `example/simple_rpc/simple_rpc.proto` grpc-gateway参数的编写
- `example/simple_rpc/build.sh` protobuf生成grpc-gateway代码

**rest部分**
- `example/simple_api/internal/svc/servicecontext.go` GatewayProxy的初始化


## 三、测试
启动服务
```shell
> # 启动rpc
> cd simple_rpc 
> go run .

> # 启动api
> cd simple_api
> go run .
```

发起测试请求
```shell
> curl -XPOST 'http://127.0.0.1:8888/gateway/simple/ping' \
  --header 'Content-Type: application/json' \
  --data-raw '{"ping":"SimpleApi"}'
```

## 四、弊端

- grpc-gateway的protobuf参数必须按规范写：/前缀/服务名/方法名，因为在rest服务注册的路由就固定为这个


## 相关仓库：
- go-zero: https://github.com/zeromicro/go-zero
- grpc-gateway：https://github.com/grpc-ecosystem/grpc-gateway
- grpc-go: https://github.com/grpc/grpc-go
- protobuf: https://github.com/protocolbuffers/protobuf
- googleapis: https://github.com/googleapis/googleapis
