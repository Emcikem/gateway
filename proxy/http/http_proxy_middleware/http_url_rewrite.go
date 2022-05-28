package http_proxy_middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"gateway/dao"
	"gateway/serializer"
	"gateway/service/model"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func HTTPURLRewriteMiddleware() gin.HandlerFunc {
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
		for _, item := range strings.Split(httpRule.UrlRewrite, ",") {
			items := strings.Split(item, " ")
			if len(items) != 2 {
				continue
			}
			regexp, err := regexp.Compile(items[0])
			if err != nil {
				fmt.Println("regexp.Compile err", err)
				continue
			}
			replacePath := regexp.ReplaceAll([]byte(c.Request.URL.Path), []byte(items[1]))
			c.Request.URL.Path = string(replacePath)
		}
		c.Next()
	}
}
