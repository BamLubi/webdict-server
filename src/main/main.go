package main

import (
	"fmt"
	"webdict-server/pkg/config"
	"webdict-server/pkg/web"
)

func main() {
	// 初始化logger
	config.InitLogger(config.LogPath, config.LogLevel)
	config.Logger.Info("Server Run Success on Mode(" + config.MODE + ")")
	fmt.Println("Server Run Success on Mode(" + config.MODE + ")")

	// 初始化Redis
	config.InitRedis()
	config.InitDB()

	// 如果是线上测试版, 则直接退出, 不需要运行服务器
	if config.MODE == "PRODUCTION TEST" {
		return
	}

	// 启动服务器
	web.RunHttp()
}
