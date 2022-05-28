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

// HTTPBlackListMiddleware 黑名单校验，要求开启验证，并且白名单为空
func HTTPBlackListMiddleware() gin.HandlerFunc {
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
		var blackIpList []string
		if serviceDetail.BlackIpList != "" {
			blackIpList = strings.Split(serviceDetail.BlackIpList, ",")
		}
		if serviceDetail.OpenAuth == 1 && len(whiteIpList) == 0 && len(blackIpList) > 0 {
			if util.InStringSlice(blackIpList, c.ClientIP()) {
				serializer.ResponseError(c, 5001, errors.New(fmt.Sprintf("%s in blackIpList", c.ClientIP())))
				c.Abort()
			}
		}
		c.Next()
	}
}
