package dao

import (
	"gorm.io/gorm"
	"time"
)

type Tenant struct {
	ID       int64     `json:"id" gorm:"column:id"`               // 自增id
	AppID    string    `json:"app_id" gorm:"column:app_id"`       // 租户id
	Name     string    `json:"name" gorm:"column:name"`           // 租户名称
	Secret   string    `json:"secret" gorm:"column:secret"`       // 密钥
	WhiteIps string    `json:"white_ips" gorm:"column:white_ips"` // ip白名单，支持前缀匹配
	Qpd      int64     `json:"qpd" gorm:"column:qpd"`             // 日请求量限制
	Qps      int64     `json:"qps" gorm:"column:qps"`             // 每秒请求量限制
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"` // 新增时间
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"` // 更新时间
	IsDelete int8      `json:"is_delete" gorm:"column:is_delete"` // 是否删除 1=删除
}

func (m *Tenant) TableName() string {
	return "gateway_app"
}

func TenantPageList(keyword string, size, page int) ([]Tenant, int64, error) {
	var list []Tenant
	total := int64(0)
	offset := (page - 1) * size

	query := DB.Select([]string{"id", "app_id", "name", "secret", "qpd", "qps"}).
		Where("is_delete = 0")
	if keyword != "" {
		query = query.Where("(app_id like ? or name like ?)",
			"%"+keyword+"%", "%"+keyword+"%")
	}
	if err := query.Limit(size).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	// TODO:能不能合起来？
	if err := query.Limit(size).Offset(offset).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func TenantDetail(id int) (Tenant, error) {
	var detail Tenant
	result := DB.Where("id = ? and is_delete = 0", id).First(&detail)
	return detail, result.Error
}
