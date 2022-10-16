package main

import (
	"webdict-server/pkg/config"
	"webdict-server/pkg/web"
)

func main() {
	config.InitLogger("logs/webdict.log", "info")
	config.Logger.Info("main run success!")
	web.RunHttp()
}
