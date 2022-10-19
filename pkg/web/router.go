package web

import (
	"github.com/gin-gonic/gin"
	"webdict-server/pkg/config"
	"webdict-server/pkg/controller"
)

func RunHttp() {
	r := gin.Default()
	// 增加tls支持
	//r.Use(TlsHandler())
	//增加拦截器
	r.Use(HttpInterceptor())
	//解决跨域
	r.Use(config.CorsConfig())
	//路由组
	appInfoGroup := r.Group("/")
	{
		appInfoGroup.POST("/literature/insert", controller.NewLiteratureController().Insert)
		appInfoGroup.POST("/literature/update", controller.NewLiteratureController().Update)
		appInfoGroup.GET("/literature/getAll", controller.NewLiteratureController().FindAll)
		appInfoGroup.GET("/literature/search/:key", controller.NewLiteratureController().FuzzyFind)
		appInfoGroup.GET("/literature/get/:id", controller.NewLiteratureController().FindById)
	}
	r.RunTLS(":8000", "/ssl/go.rtclab.top/ssl.pem", "/ssl/go.rtclab.top/ssl.key")
	//r.Run(config.HOST + ":" + config.PORT)
}
