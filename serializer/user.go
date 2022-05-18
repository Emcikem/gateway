package serializer

import "singo/dao"

// User 用户序列化器
type User struct {
	UserName string   `json:"username"`
	Avatar   string   `json:"avatar"`
	Roles    []string `json:"roles"`
}

// BuildUser 序列化用户
func BuildUser(user dao.User) User {
	return User{
		UserName: user.Username,
		Avatar:   user.Avatar,
		Roles:    []string{"admin"},
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user dao.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
