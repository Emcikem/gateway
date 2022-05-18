package service

import (
	"github.com/gin-gonic/gin"
	"singo/dao"
	"singo/serializer"
	"singo/util"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user dao.User

	if err := dao.DB.Where("username = ? and is_deleted = 0", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	tokenString, err := util.GetToken(user.Username)
	if err != nil {
		return serializer.ParamErr("系统错误", nil)
	}
	return serializer.Response{
		Data: map[string]interface{}{
			"token": tokenString,
		},
	}
}
