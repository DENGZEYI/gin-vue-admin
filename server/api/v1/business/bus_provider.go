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

type BusProviderApi struct {
}

var busProviderService = service.ServiceGroupApp.BusinessServiceGroup.BusProviderService

// CreateBusProvider 创建BusProvider
// @Tags BusProvider
// @Summary 创建BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusProvider true "创建BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busProvider/createBusProvider [post]
func (busProviderApi *BusProviderApi) CreateBusProvider(c *gin.Context) {
	var busProvider business.BusProvider
	_ = c.ShouldBindJSON(&busProvider)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(busProvider, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := busProviderService.CreateBusProvider(busProvider); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBusProvider 删除BusProvider
// @Tags BusProvider
// @Summary 删除BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusProvider true "删除BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busProvider/deleteBusProvider [delete]
func (busProviderApi *BusProviderApi) DeleteBusProvider(c *gin.Context) {
	var busProvider business.BusProvider
	_ = c.ShouldBindJSON(&busProvider)
	if err := busProviderService.DeleteBusProvider(busProvider); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBusProviderByIds 批量删除BusProvider
// @Tags BusProvider
// @Summary 批量删除BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /busProvider/deleteBusProviderByIds [delete]
func (busProviderApi *BusProviderApi) DeleteBusProviderByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := busProviderService.DeleteBusProviderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBusProvider 更新BusProvider
// @Tags BusProvider
// @Summary 更新BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusProvider true "更新BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busProvider/updateBusProvider [put]
func (busProviderApi *BusProviderApi) UpdateBusProvider(c *gin.Context) {
	var busProvider business.BusProvider
	_ = c.ShouldBindJSON(&busProvider)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(busProvider, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := busProviderService.UpdateBusProvider(busProvider); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBusProvider 用id查询BusProvider
// @Tags BusProvider
// @Summary 用id查询BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BusProvider true "用id查询BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busProvider/findBusProvider [get]
func (busProviderApi *BusProviderApi) FindBusProvider(c *gin.Context) {
	var busProvider business.BusProvider
	_ = c.ShouldBindQuery(&busProvider)
	if rebusProvider, err := busProviderService.GetBusProvider(busProvider.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebusProvider": rebusProvider}, c)
	}
}

// GetBusProviderList 分页获取BusProvider列表
// @Tags BusProvider
// @Summary 分页获取BusProvider列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusProviderSearch true "分页获取BusProvider列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busProvider/getBusProviderList [get]
func (busProviderApi *BusProviderApi) GetBusProviderList(c *gin.Context) {
	var pageInfo businessReq.BusProviderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := busProviderService.GetBusProviderInfoList(pageInfo); err != nil {
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
