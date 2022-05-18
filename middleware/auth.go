package middleware

import (
	"singo/dao"
	"singo/serializer"
	"singo/util"

	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Admin-Token")
		if err != nil {
			return
		}
		username, err := util.GetUsername(tokenString)
		if err != nil {
			return
		}
		if username != "" {
			user, err := dao.GetUser(username)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*dao.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
