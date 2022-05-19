package dao

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type GatewayService struct {
	ID           int64     `json:"id" gorm:"column:id"`                       // 自增id
	LoadType     int8      `json:"load_type" gorm:"column:load_type"`         // 负载类型 0=http 1=tcp 2=grpc
	ServiceName  string    `json:"service_name" gorm:"column:service_name"`   // 服务名称 6-128 数字字母下划线
	ServiceDesc  string    `json:"service_desc" gorm:"column:service_desc"`   // 服务描述
	ServiceAddr  string    `json:"service_addr" gorm:"column:service_addr"`   // 服务地址
	TotalNode    int       `json:"total_node" gorm:"column:total_node"`       // 结点数
	OpenAuth     int8      `json:"open_auth" gorm:"open_auth"`                // 是否开启权限校验
	ClientLimit  int       `json:"client_limit" gorm:"client_limit"`          // 客户端限流
	ServerLimit  int       `json:"server_limit" gorm:"server_limit"`          // 服务端限流
	RoundType    int8      `json:"round_type" gorm:"column:round_type"`       // 轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash
	IpList       string    `json:"ip_list" gorm:"column:ip_list"`             // ip列表
	WeightList   string    `json:"weight_list" gorm:"column:weight_list"`     // 权重列表
	WhiteIpList  string    `json:"white_ip_list" gorm:"column:white_ip_list"` // 白名单
	BlackIpList  string    `json:"black_ip_list" gorm:"column:black_ip_list"` // 黑名单
	RemoteParams string    `json:"remote_params" gorm:"column:remote_params"` // 远程调度的参数
	UpdateAt     time.Time `json:"update_at" gorm:"column:update_at"`         // 更新时间
	CreateAt     time.Time `json:"create_at" gorm:"column:create_at"`         // 新增时间
	IsDelete     int8      `json:"is_delete" gorm:"column:is_delete"`         // 是否删除 1=删除
}

func (m *GatewayService) TableName() string {
	return "gateway_service"
}

func ServicePageList(keyword string, size, page int) ([]GatewayService, int64, error) {
	var list []GatewayService
	total := int64(0)
	offset := (page - 1) * size

	query := DB.Select([]string{"id", "load_type", "service_name", "service_desc", "service_addr", "total_node"}).
		Where("is_delete = 0")
	if keyword != "" {
		query = query.Where("(service_name like ? or service_desc like ?)",
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

func ServiceDetail(id int) (GatewayService, error) {
	var detail GatewayService
	result := DB.Where("id = ? and is_delete = 0", id).First(&detail)
	return detail, result.Error
}

func ServiceDelete(id int) error {
	err := DB.Where("id = ? and is_delete = 0", id).Update("is_delete", 1).Error
	return err
}

func (m *GatewayService) SaveService() error {
	return DB.Save(m).Error
}

func (m *GatewayService) GetIpListByModel() []string {
	return strings.Split(m.IpList, ",")
}

func (m *GatewayService) GetWeightListByModel() []string {
	return strings.Split(m.WeightList, ",")
}
