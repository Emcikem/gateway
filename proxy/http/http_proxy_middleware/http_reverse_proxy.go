package http_proxy_middleware

import (
	"errors"
	"gateway/dao"
	"gateway/proxy/reverse_proxy"
	"gateway/serializer"
	"github.com/gin-gonic/gin"
)

func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			serializer.ResponseError(c, 5001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.GatewayService)

		lb, err := LoadBalancerHandler.GetLoadBalance(serviceDetail)
		if err != nil {
			serializer.ResponseError(c, 5001, errors.New("service not found"))
			c.Abort()
			return
		}

		trans, err := TransportHandler.GetTrans(serviceDetail)
		if err != nil {
			serializer.ResponseError(c, 5001, err)
			c.Abort()
			return
		}

		proxy := reverse_proxy.NewLoadBalanceReverseProxy(c, lb, trans)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
