package dao

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户模型
type User struct {
	ID        int64     `json:"id" gorm:"column:id"`             // 自增id
	Username  string    `json:"username" gorm:"column:username"` // 用户名
	Password  string    `json:"password" gorm:"column:password"` // 密码
	Avatar    string    `json:"avatar" gorm:"column:avatar"`
	Status    string    `json:"status" gorm:"column:status"` // 用户状态
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	IsDeleted int8      `json:"is_deleted" gorm:"column:is_deleted"` // 是否删除
}

func (user *User) TableName() string {
	return "gateway_admin"
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用username获取用户
func GetUser(username string) (User, error) {
	var user User
	result := DB.Where("username = ?", username).First(&user)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
