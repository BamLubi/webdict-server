package web

import (
	"github.com/gin-gonic/gin"
	"webdict-server/pkg/config"
	"webdict-server/pkg/controller"
)

func RunHttp() {
	r := gin.Default()
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
		appInfoGroup.GET("/literature/get/:key", controller.NewLiteratureController().FuzzyFind)
	}
	r.Run(config.HOST + config.PORT)
}
