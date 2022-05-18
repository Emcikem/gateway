package service

import (
	"github.com/gin-gonic/gin"
	"singo/dao"
	"singo/serializer"
)

// RemoteListService 列表页查询接口
type RemoteListService struct {
	KeyWord string `form:"keyword" json:"keyword" binding:"required,min=5,max=30"`
	Page    int    `form:"page" json:"page" binding:"required,min=6,max=40"`
	Size    int    `form:"size" json:"size" binding:"required,min=6,max=40"`
}

// RemoteDetailService 服务详情
type RemoteDetailService struct {
	Id int `form:"id" json:"id" binding:"required,min=6,max=40"`
}

type GatewayServiceHttpRule struct {
	RuleType       int8   `json:"rule_type" form:"rule_type"`             // 匹配类型 0=url前缀url_prefix 1=域名domain
	Rule           string `json:"rule" form:"rule"`                       // type=domain表示域名，type=url_prefix时表示url前缀
	NeedHttps      int8   `json:"need_https" form:"need_https"`           // 支持https 1=支持
	NeedStripUri   int8   `json:"need_strip_uri" form:"need_strip_uri"`   // 启用strip_uri 1=启用
	NeedWebsocket  int8   `json:"need_websocket" form:"need_websocket"`   // 是否支持websocket 1=支持
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite"`         // url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔
	HeaderTransfer string `json:"header_transfer" form:"header_transfer"` // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

type GatewayServiceGrpcRule struct {
	Port           int    `json:"port" form:"port"`                       // 端口
	HeaderTransfer string `json:"header_transfer" form:"header_transfer"` // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

type GatewayServiceTcpRule struct {
	Port int `json:"port" form:"port"` // 端口号
}

// PageList 分页查询
func (service *RemoteListService) PageList(c *gin.Context) serializer.Response {

	pageList, total, err := dao.ServicePageList(service.KeyWord, service.Size, service.Page)
	if err != nil {
		return serializer.Response{}
	}

	// 转化为前端对象，同时从redis中读取访问量

	return serializer.Response{
		Data: map[string]interface{}{
			"pageList": pageList,
			"total":    total,
		},
	}
}
