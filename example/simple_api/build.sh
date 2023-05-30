#!/bin/bash

case $1 in
server)
   # 更新api端代码
   goctl api go --api=./simple_api.api --dir=./
  ;;
esac