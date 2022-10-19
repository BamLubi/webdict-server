package web

import (
	"github.com/gin-gonic/gin"
	"webdict-server/pkg/config"
	"webdict-server/pkg/controller"
)

func RunServer() {
	r := gin.Default()
	// BUG: 增加tls支持, 支持http重定向, 目前有问题
	//r.Use(TlsHandler())
	//增加拦截器
	r.Use(HttpInterceptor())
	//解决跨域
	r.Use(config.CorsConfig())

	//路由组
	appInfoGroup := r.Group("/literature")
	{
		appInfoGroup.POST("/insert", controller.NewLiteratureController().Insert)
		appInfoGroup.POST("/update", controller.NewLiteratureController().Update)
		appInfoGroup.GET("/getAll", controller.NewLiteratureController().FindAll)
		appInfoGroup.GET("/search/:key", controller.NewLiteratureController().FuzzyFind)
		appInfoGroup.GET("/get/:id", controller.NewLiteratureController().FindById)
	}

	// 如果是生成模式,则使用HTTPS,否则使用HTTP
	if config.MODE == "PRODUCTION" {
		r.RunTLS(config.ServerUrl, config.ServerSSLCert, config.ServerSSLKey)
	} else {
		r.Run(config.ServerUrl)
	}
}
