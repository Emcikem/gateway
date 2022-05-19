package service

import (
	"gateway/dao"
	"gateway/serializer"
	"gateway/service/model"
	"github.com/gin-gonic/gin"
)

// QueryPageList 分页查询
func QueryPageList(c *gin.Context, service *model.ServiceListReq) serializer.Response {
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

func QueryServiceDetail(c *gin.Context, service *model.ServiceDetailReq) serializer.Response {
	detail, err := dao.ServiceDetail(service.Id)
	if err != nil {
		return serializer.DBErr("数据库查询失败", err)
	}
	return serializer.Response{
		Data: serializer.BuildServiceVO(&detail),
	}
}

func SaveServiceDetail(c *gin.Context, service *model.ServiceDetailVO) serializer.Response {
	serviceModel := serializer.BuildServiceEntity(service)
	err := serviceModel.SaveService()
	if err != nil {
		return serializer.DBErr("数据库保存失败", err)
	}
	return serializer.Response{
		Data: "",
	}
}

func SaveServiceDelete(c *gin.Context, service *model.ServiceDeleteReq) serializer.Response {
	err := dao.ServiceDelete(service.Id)
	if err != nil {
		return serializer.DBErr("数据库删除", err)
	}
	return serializer.Response{
		Data: "",
	}
}
