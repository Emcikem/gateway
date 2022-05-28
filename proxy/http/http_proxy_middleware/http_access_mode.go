package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
)

func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//service, err := nil, nil
		//if err != nil {
		//	serializer.ResponseError(c, 1001, err)
		//	c.Abort()
		//	return
		//}
		//c.Set("service", service)
		c.Next()
	}
}
