package model

type ServiceDetailVO struct {
	ID           int64  `json:"id" gorm:"column:id"`                       // 自增id
	ServiceName  string `json:"service_name" gorm:"column:service_name"`   // 服务名称 6-128 数字字母下划线
	ServiceDesc  string `json:"service_desc" gorm:"column:service_desc"`   // 服务描述
	ServiceAddr  string `json:"service_addr" gorm:"column:service_addr"`   // 服务地址
	LoadType     int8   `json:"load_type" gorm:"column:load_type"`         // 负载类型 0=http 1=tcp 2=grpc
	OpenAuth     int8   `json:"open_auth" gorm:"open_auth"`                // 是否开启权限校验
	ClientLimit  int    `json:"client_limit" gorm:"client_limit"`          // 客户端限流
	ServerLimit  int    `json:"server_limit" gorm:"server_limit"`          // 服务端限流
	RoundType    int8   `json:"round_type" gorm:"column:round_type"`       // 轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash
	IpList       string `json:"ip_list" gorm:"column:ip_list"`             // ip列表
	WeightList   string `json:"weight_list" gorm:"column:weight_list"`     // 权重列表
	WhiteIpList  string `json:"white_ip_list" gorm:"column:white_ip_list"` // 白名单
	BlackIpList  string `json:"black_ip_list" gorm:"column:black_ip_list"` // 黑名单
	RemoteParams string `json:"remote_params" gorm:"column:remote_params"` // 远程调度的参数
}

type GatewayServiceHttpRuleVO struct {
	RuleType               int8   `json:"rule_type" form:"rule_type"`                                      // 匹配类型 0=url前缀url_prefix 1=域名domain
	Rule                   string `json:"rule" form:"rule"`                                                // type=domain表示域名，type=url_prefix时表示url前缀
	NeedHttps              int8   `json:"need_https" form:"need_https"`                                    // 支持https 1=支持
	NeedStripUri           int8   `json:"need_strip_uri" form:"need_strip_uri"`                            // 启用strip_uri 1=启用
	NeedWebsocket          int8   `json:"need_websocket" form:"need_websocket"`                            // 是否支持websocket 1=支持
	UrlRewrite             string `json:"url_rewrite" form:"url_rewrite"`                                  // url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔
	HeaderTransfer         string `json:"header_transfer" form:"header_transfer"`                          // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"column:upstream_connect_timeout"` // 建立连接超时, 单位s
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"column:upstream_header_timeout"`   // 获取header超时, 单位s
	UpstreamIDleTimeout    int    `json:"upstream_idle_timeout" form:"column:upstream_idle_timeout"`       // 链接最大空闲时间, 单位s
	UpstreamMaxIDle        int    `json:"upstream_max_idle" form:"column:upstream_max_idle"`               // 最大空闲链接数
}

type GatewayServiceGrpcRuleVO struct {
	Port           int    `json:"port" form:"port"`                       // 端口
	HeaderTransfer string `json:"header_transfer" form:"header_transfer"` // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

type GatewayServiceTcpRuleVO struct {
	Port int `json:"port" form:"port"` // 端口号
}

// ServiceListReq 列表页查询接口
type ServiceListReq struct {
	KeyWord string `form:"keyword" json:"keyword" binding:"required,min=5,max=30"`
	Page    int    `form:"page" json:"page" binding:"required,min=6,max=40"`
	Size    int    `form:"size" json:"size" binding:"required,min=6,max=40"`
}

// ServiceDetailReq 服务详情查询入参
type ServiceDetailReq struct {
	Id int `form:"id" json:"id"`
}

// ServiceDeleteReq 服务删除入参
type ServiceDeleteReq struct {
	Id int `form:"id" json:"id"`
}
