package controller

import (
	"gateway/dao"
	"gateway/dto"
	"gateway/middleware"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {
}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
	group.GET("/detail", service.ServiceDetail)
}

// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(c *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//从db中分页读取基本信息
	serviceInfo := &dao.ServiceInfo{}
	list, total, err := serviceInfo.PageList(tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//格式化输出信息
	outList := []dto.ServiceListItemOutput{}
	for _, listItem := range list {
		outItem := dto.ServiceListItemOutput{
			ID:          listItem.ID,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			LoadType:    listItem.LoadType,
			ServiceAddr: listItem.ServiceAddr,
			TotalNode:   listItem.TotalNode,
		}
		outList = append(outList, outItem)
	}

	out := &dto.ServiceListOutput{
		Total: total,
		List:  outList,
	}
	middleware.ResponseSuccess(c, out)
}

// ServiceDetail godoc
// @Summary 服务详情
// @Description 服务详情
// @Tags 服务管理
// @ID /service/detail
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} middleware.Response{data=dao.ServiceDetail} "success"
// @Router /service/detail [get]
func (service *ServiceController) ServiceDetail(c *gin.Context) {
	params := &dto.ServiceDetailInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	serviceDetail := &dao.ServiceDetail{}
	detail, err := serviceDetail.QueryById(tx, params.Id)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, detail)
}
