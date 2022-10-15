package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

// RemoteAddr 获取客户端地址, 包括 ip 和 port
func RemoteAddr(req *gin.Context) []string {
	var addr string

	addr = req.Request.RemoteAddr

	if addr == "" {
		addr = req.Request.Header.Get("ipv4")
	}

	if addr == "" {
		addr = req.Request.Header.Get("XForwardedFor")
	}

	if addr == "" {
		addr = req.Request.Header.Get("X-Forwarded-For")
	}

	if addr == "" {
		addr = req.Request.Header.Get("X-Real-Ip")
	}

	if addr == "" {
		addr = "127.0.0.1:0000"
	}

	return strings.Split(addr, ":")
}

// HttpInterceptor 自定义拦截器
func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 example 变量
		c.Set("example", "12345")
		// 请求前
		addr := RemoteAddr(c)
		log.Print("-- 拦截器 -- ", "ip:", addr[0], ";port:", addr[1])
		//定义错误,终止并返回该JSON
		//requestURI := c.Request.RequestURI
		//c.AbortWithStatusJSON(500, "error")
		//通过请求
		c.Next()
	}
}
