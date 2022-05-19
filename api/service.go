package api

import (
	"gateway/service"
	"gateway/service/model"
	"github.com/gin-gonic/gin"
)

// ServiceDetailQuery 服务详情页查询
func ServiceDetailQuery(c *gin.Context) {
	var detailReq model.ServiceDetailReq
	if err := c.ShouldBind(&detailReq); err == nil {
		res := service.QueryServiceDetail(c, &detailReq)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ServiceListQuery 服务列表页查询
func ServiceListQuery(c *gin.Context) {
	var listInput model.ServiceListReq
	if err := c.ShouldBind(&listInput); err == nil {
		res := service.QueryPageList(c, &listInput)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ServiceDetailSave(c *gin.Context) {
	var detail model.ServiceDetailVO
	if err := c.ShouldBind(&detail); err == nil {
		res := service.SaveServiceDetail(c, &detail)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ServiceDelete(c *gin.Context) {
	var deleteReq model.ServiceDeleteReq
	if err := c.ShouldBind(&deleteReq); err == nil {
		res := service.SaveServiceDelete(c, &deleteReq)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ServiceUpdate(c *gin.Context) {

}
