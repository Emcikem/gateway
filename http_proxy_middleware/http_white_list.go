package http_proxy_middleware

import (
	"errors"
	"fmt"
	"gateway/dao"
	"gateway/serializer"
	"gateway/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func HTTPWhiteListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			serializer.ResponseError(c, 5001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.GatewayService)
		var whiteIpList []string
		if serviceDetail.WhiteIpList != "" {
			whiteIpList = strings.Split(serviceDetail.WhiteIpList, ",")
		}
		if serviceDetail.OpenAuth == 1 && len(whiteIpList) > 0 {
			if util.InStringSlice(whiteIpList, c.ClientIP()) {
				serializer.ResponseError(c, 5001, errors.New(fmt.Sprintf("%s not in whiteIpList", c.ClientIP())))
				c.Abort()
			}
		}
		c.Next()
	}
}
