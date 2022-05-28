package http_proxy_middleware

import (
	"errors"
	"fmt"
	"gateway/dao"
	"gateway/serializer"
	"gateway/util"
	"github.com/gin-gonic/gin"
)

// HTTPFlowLimitMiddleware 服务端限流，就是对于整个服务限流，客户端限流，就是对于某个ip进行限流
func HTTPFlowLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			serializer.ResponseError(c, 5001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.GatewayService)

		// 服务端限流
		serverQps := float64(serviceDetail.ServerLimit)
		if serverQps != 0 {
			serviceLimiter, err := util.FlowLimiterHandler.GetLimiter(
				fmt.Sprintf("%s_%s", util.FlowServicePrefix, serviceDetail.ServiceName),
				serverQps,
			)
			if err != nil {
				serializer.ResponseError(c, 5001, err)
				c.Abort()
				return
			}
			if !serviceLimiter.Allow() {
				serializer.ResponseError(c, 5001, errors.New(fmt.Sprintf("service flow limit %v", serverQps)))
				c.Abort()
				return
			}
		}

		// 客户端限流
		clientQps := float64(serviceDetail.ServerLimit)
		if clientQps != 0 {
			clientLimiter, err := util.FlowLimiterHandler.GetLimiter(
				fmt.Sprintf("%s_%s_%s", util.FlowServicePrefix, serviceDetail.ServiceName, c.ClientIP()),
				clientQps,
			)
			if err != nil {
				serializer.ResponseError(c, 5001, err)
				c.Abort()
				return
			}
			if !clientLimiter.Allow() {
				serializer.ResponseError(c, 5001, errors.New(fmt.Sprintf("service flow limit %v", clientQps)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
