package http_proxy_middleware

import (
	"encoding/json"
	"errors"
	"gateway/dao"
	"gateway/serializer"
	"gateway/service/model"
	"github.com/gin-gonic/gin"
	"strings"
)

// HTTPHeaderTransferMiddleware http的header进行修改
func HTTPHeaderTransferMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			serializer.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.GatewayService)
		httpRule := model.GatewayServiceHttpRuleVO{}
		err := json.Unmarshal([]byte(serviceDetail.RemoteParams), &httpRule)
		if err != nil {
			serializer.ResponseError(c, 2002, errors.New("service not found"))
			c.Abort()
		}
		for _, item := range strings.Split(httpRule.HeaderTransfer, ",") {
			items := strings.Split(item, " ")
			if len(items) != 3 {
				continue
			}
			if items[0] == "add" || items[0] == "edit" {
				c.Request.Header.Set(items[1], items[2])
			}
			if items[0] == "del" {
				c.Request.Header.Del(items[1])
			}
		}
		c.Next()
	}
}
