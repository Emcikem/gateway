package dao

import (
	"gateway/dto"
	"gorm.io/gorm"
	"time"
)

type ServiceInfo struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	LoadType    int       `json:"load_type" gorm:"column:load_type" description:"负载类型 0=http 1=tcp 2=grpc"`
	ServiceName string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	ServiceAddr string    `json:"service_addr" gorm:"column:service_addr" description:"服务地址"`
	TotalNode   int       `json:"total_node" gorm:"column:total_node" description:"结点数"`
	UpdatedAt   time.Time `json:"create_at" gorm:"column:create_at" description:"更新时间"`
	CreatedAt   time.Time `json:"update_at" gorm:"column:update_at" description:"添加时间"`
	IsDelete    int8      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
}

type ServiceDetail struct {
	ID           int64     `json:"id" gorm:"primary_key"`
	LoadType     int       `json:"load_type" gorm:"column:load_type" description:"负载类型 0=http 1=tcp 2=grpc"`
	ServiceName  string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc  string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	ServiceAddr  string    `json:"service_addr" gorm:"column:service_addr" description:"服务地址"`
	TotalNode    int       `json:"total_node" gorm:"column:total_node" description:"结点数"`
	RoundType    int8      `json:"round_type" gorm:"column:round_type" description:"轮询方式"`
	IpList       string    `json:"ip_list" gorm:"column:ip_list" description:"ip列表"`
	WeightList   string    `json:"weight_list" gorm:"column:weight_list" description:"权重列表"`
	WhiteIpList  string    `json:"white_ip_list" gorm:"column:white_ip_list" description:"白名单ip列表"`
	BlackIpList  string    `json:"black_ip_list" gorm:"column:black_ip_list" description:"黑名单ip列表"`
	RemoteParams string    `json:"remote_params" gorm:"column:remote_params" description:"业务参数"`
	UpdatedAt    time.Time `json:"create_at" gorm:"column:create_at" description:"更新时间"`
	CreatedAt    time.Time `json:"update_at" gorm:"column:update_at" description:"添加时间"`
	IsDelete     int8      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
}

func (t *ServiceInfo) TableName() string {
	return "gateway_service"
}

func (t *ServiceDetail) TableName() string {
	return "gateway_service"
}

func (t *ServiceInfo) PageList(tx *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	total := int64(0)
	var list []ServiceInfo
	offset := (param.PageNo - 1) * param.PageSize

	query := tx.Table(t.TableName()).
		Where("is_delete = 0")
	if param.Info != "" {
		query = query.Where("(service_name like ? or service_desc like ?)",
			"%"+param.Info+"%", "%"+param.Info+"%")
	}
	if err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(param.PageSize).Offset(offset).Count(&total)
	return list, total, nil
}

func (t *ServiceDetail) QueryById(tx *gorm.DB, id int) (*ServiceDetail, error) {
	detail := ServiceDetail{}

	if err := tx.Table(t.TableName()).Where("id = ? and is_delete = 0", id).Find(&detail).Error; err != nil {
		return nil, err
	}

	return &detail, nil
}
