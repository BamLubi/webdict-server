package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"webdict-server/pkg/config"
	"webdict-server/pkg/web"
)

func main() {
	var mode string
	// 获取参数
	if len(os.Args) > 1 && os.Args[1] == "-DEV" {
		mode = "DEVELOPMENT"
		viper.SetConfigName("dev")
	} else if len(os.Args) > 1 && os.Args[1] == "-PROD_TEST" {
		mode = "PRODUCTION TEST"
		viper.SetConfigName("app")
	} else {
		mode = "PRODUCTION"
		viper.SetConfigName("app")
	}
	viper.SetConfigType("properties")
	viper.AddConfigPath("./")
	path := viper.GetString("logger.path")
	level := viper.GetString("logger.level")

	// 初始化logger
	config.InitLogger(path, level)
	config.Logger.Info("Server Run Success on Mode(" + mode + ")")
	fmt.Println("Server Run Success on Mode(" + mode + ")")

	// 如果是线上测试版，则直接退出，不需要运行服务器
	if mode == "PRODUCTION TEST" {
		return
	}

	// 启动服务器
	web.RunHttp()
}
