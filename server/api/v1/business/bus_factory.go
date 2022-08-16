package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusFactoryApi struct {
}

var busFactoryService = service.ServiceGroupApp.BusinessServiceGroup.BusFactoryService

// CreateBusFactory 创建BusFactory
// @Tags BusFactory
// @Summary 创建BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusFactory true "创建BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busFactory/createBusFactory [post]
func (busFactoryApi *BusFactoryApi) CreateBusFactory(c *gin.Context) {
	var busFactory business.BusFactory
	_ = c.ShouldBindJSON(&busFactory)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(busFactory, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := busFactoryService.CreateBusFactory(busFactory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBusFactory 删除BusFactory
// @Tags BusFactory
// @Summary 删除BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusFactory true "删除BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busFactory/deleteBusFactory [delete]
func (busFactoryApi *BusFactoryApi) DeleteBusFactory(c *gin.Context) {
	var busFactory business.BusFactory
	_ = c.ShouldBindJSON(&busFactory)
	if err := busFactoryService.DeleteBusFactory(busFactory); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBusFactoryByIds 批量删除BusFactory
// @Tags BusFactory
// @Summary 批量删除BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /busFactory/deleteBusFactoryByIds [delete]
func (busFactoryApi *BusFactoryApi) DeleteBusFactoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := busFactoryService.DeleteBusFactoryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBusFactory 更新BusFactory
// @Tags BusFactory
// @Summary 更新BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusFactory true "更新BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busFactory/updateBusFactory [put]
func (busFactoryApi *BusFactoryApi) UpdateBusFactory(c *gin.Context) {
	var busFactory business.BusFactory
	_ = c.ShouldBindJSON(&busFactory)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(busFactory, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := busFactoryService.UpdateBusFactory(busFactory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBusFactory 用id查询BusFactory
// @Tags BusFactory
// @Summary 用id查询BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BusFactory true "用id查询BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busFactory/findBusFactory [get]
func (busFactoryApi *BusFactoryApi) FindBusFactory(c *gin.Context) {
	var busFactory business.BusFactory
	_ = c.ShouldBindQuery(&busFactory)
	if rebusFactory, err := busFactoryService.GetBusFactory(busFactory.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebusFactory": rebusFactory}, c)
	}
}

// GetBusFactoryList 分页获取BusFactory列表
// @Tags BusFactory
// @Summary 分页获取BusFactory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusFactorySearch true "分页获取BusFactory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busFactory/getBusFactoryList [get]
func (busFactoryApi *BusFactoryApi) GetBusFactoryList(c *gin.Context) {
	var pageInfo businessReq.BusFactorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := busFactoryService.GetBusFactoryInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
