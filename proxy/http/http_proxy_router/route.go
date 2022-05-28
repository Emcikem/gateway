package http_proxy_router

import (
	http_proxy_middleware2 "gateway/proxy/http/http_proxy_middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	oauth := router.Group("/oauth")
	oauth.Use(
		http_proxy_middleware2.HTTPAccessModeMiddleware(),
		http_proxy_middleware2.HTTPFlowLimitMiddleware(),
		http_proxy_middleware2.HTTPWhiteListMiddleware(),
		http_proxy_middleware2.HTTPBlackListMiddleware(),
		http_proxy_middleware2.HTTPHeaderTransferMiddleware(),
		http_proxy_middleware2.HTTPStripUriMiddleware(),
		http_proxy_middleware2.HTTPURLRewriteMiddleware(),
		http_proxy_middleware2.HTTPReverseProxyMiddleware(),
	)
	return router
}
