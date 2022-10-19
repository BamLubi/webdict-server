package web

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"strings"
	"webdict-server/pkg/config"
)

// remoteAddr 获取客户端地址, 包括 ip 和 port
func remoteAddr(req *gin.Context) []string {
	var addr string

	addr = req.Request.RemoteAddr

	if addr == "" {
		addr = req.Request.Header.Get("ipv4") + ":0000"
	}

	if addr == "" {
		addr = req.Request.Header.Get("XForwardedFor") + ":0000"
	}

	if addr == "" {
		addr = req.Request.Header.Get("X-Forwarded-For") + ":0000"
	}

	if addr == "" {
		addr = req.Request.Header.Get("X-Real-Ip") + ":0000"
	}

	if addr == "" {
		addr = "127.0.0.1:0000"
	}

	return strings.Split(addr, ":")
}

func isAllowedHost(ip string) bool {
	// Redis
	isAvail, _ := redis.String(config.RD.Do("Get", ip))
	if isAvail == "true" {
		return true
	}
	return false
}

// HttpInterceptor 自定义拦截器
func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		addr := remoteAddr(c)
		// ip过滤
		if !isAllowedHost(addr[0]) {
			config.Logger.Warn("Request Refuse: " + addr[0] + ":" + addr[1] + "; " + c.Request.Method + "; " + c.Request.RequestURI)
			c.AbortWithStatusJSON(500, gin.H{
				"msg": "非法IP访问",
			})
			return
		}
		//通过请求
		config.Logger.Info("Request Allow: " + addr[0] + ":" + addr[1] + "; " + c.Request.Method + "; " + c.Request.RequestURI)
		c.Next()
	}
}
