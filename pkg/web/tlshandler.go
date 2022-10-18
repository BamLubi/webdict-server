package web

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":8000",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		if err != nil {
			return
		}

		c.Next()
	}
}
