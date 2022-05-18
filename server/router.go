package server

import (
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("admin")
	{
		v1.GET("ping", api.Ping)
		// 用户注册
		v1.POST("/register", api.UserRegister)
		// 用户登录
		v1.POST("/login", api.UserLogin)
		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("/admin_info", api.UserMe)
			auth.DELETE("/logout", api.UserLogout)
		}
	}
	return r
}
