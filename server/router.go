package server

import (
	"gateway/api"
	"gateway/middleware"

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
		v1.POST("/register", api.UserRegister) // 用户注册
		v1.POST("/login", api.UserLogin)       // 用户登录

		auth := v1.Group("") // 需要登录保护的
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("/admin_info", api.UserMe)  // 用户详情
			auth.POST("/logout", api.UserLogout) // 退出登录
		}
	}
	v2 := r.Group("service")
	{
		v2.GET("/detail", api.ServiceDetailQuery) // 服务详情页
		v2.GET("/list", api.ServiceListQuery)     // 列表页查询
		v2.POST("/save", api.ServiceDetailSave)   // 服务保存
		v2.DELETE("/delete", api.ServiceDelete)   // 服务删除
		v2.POST("/update", api.ServiceUpdate)     // 服务更新
		// 需要登录保护的
		auth := v2.Group("")
		auth.Use(middleware.AuthRequired())
		{

		}
	}
	return r
}
