package http_proxy_middleware

import (
	"encoding/json"
	"errors"
	"gateway/dao"
	"gateway/serializer"
	"gateway/service/model"
	"gateway/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func HTTPStripUriMiddleware() gin.HandlerFunc {
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
		// http:127.0.0.1:8080/test_http_string/abbb
		// http:127.0.0.1:2004/abbb
		if httpRule.RuleType == util.HTTPRuleTypePrefixURL && httpRule.NeedStripUri == 1 {
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, httpRule.Rule, "", 1)
		}
		c.Next()
	}
}
