#!/bin/bash

# 编译
go build ./src/main/main.go
# 赋权限
chmod 777 main

# 设置生产环境
export GIN_MODE=release
# 运行
./main
