#!/bin/bash

# 查找是否存在已有的进程
echo "Finding exist Process..."
# shellcheck disable=SC2009
webdict_server_pid=$(ps -ef | grep webdict_server | grep -v "grep" | awk '{print $2}')
for id in $webdict_server_pid
do
    kill -9 "$id"
    echo "Killed $id"
done
echo "Finding exist Process done"

echo "Running..."

# 编译
# go build -o webdict_server ./src/main/main.go

# 赋权限
chmod 777 webdict_server

# 设置生产环境
export GIN_MODE=release

# 运行
./webdict_server
