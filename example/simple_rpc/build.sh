#!/bin/bash

GEN_DIR=./pb/
PROTO_PATH="-I ${HOME}/protoc/protoc-23.2-osx-aarch_64/include -I ${HOME}/protoc/googleapis -I ./"

# gateway 代码生成
function gen_gateway() {
  CLIENT_FILE=$1

  protoc ${PROTO_PATH} --grpc-gateway_out ${GEN_DIR} \
    --grpc-gateway_opt logtostderr=true \
    $CLIENT_FILE
}

set -x

case $1 in
server)
   # 生成 服务端和客户端
   goctl rpc protoc ./simple_rpc.proto --go_out=${GEN_DIR} --go-grpc_out=${GEN_DIR} --zrpc_out=. ${PROTO_PATH}
   # 生成 grpc gateway
   gen_gateway ./simple_rpc.proto
  ;;
esac