package middleware

import (
	"gateway/dao"
	"gateway/serializer"
	"gateway/util"

	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Admin-Token")
		if err != nil {
			return
		}
		if username, err := util.GetUsername(tokenString); err != nil {
			return
		} else if username != "" {
			if user, err := dao.GetUser(username); err == nil {
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
