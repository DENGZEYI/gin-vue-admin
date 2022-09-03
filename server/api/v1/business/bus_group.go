package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusGroupApi struct {
}

var busGroupService = service.ServiceGroupApp.BusinessServiceGroup.BusGroupService

// CreateBusGroup 创建BusGroup
// @Tags BusGroup
// @Summary 创建BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusGroup true "创建BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGroup/createBusGroup [post]
func (busGroupApi *BusGroupApi) CreateBusGroup(c *gin.Context) {
	var busGroup system.BusGroup
	_ = c.ShouldBindJSON(&busGroup)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(busGroup, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := busGroupService.CreateBusGroup(busGroup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBusGroup 删除BusGroup
// @Tags BusGroup
// @Summary 删除BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusGroup true "删除BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busGroup/deleteBusGroup [delete]
func (busGroupApi *BusGroupApi) DeleteBusGroup(c *gin.Context) {
	var busGroup system.BusGroup
	_ = c.ShouldBindJSON(&busGroup)
	if err := busGroupService.DeleteBusGroup(busGroup); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBusGroupByIds 批量删除BusGroup
// @Tags BusGroup
// @Summary 批量删除BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /busGroup/deleteBusGroupByIds [delete]
func (busGroupApi *BusGroupApi) DeleteBusGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := busGroupService.DeleteBusGroupByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBusGroup 更新BusGroup
// @Tags BusGroup
// @Summary 更新BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusGroup true "更新BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busGroup/updateBusGroup [put]
func (busGroupApi *BusGroupApi) UpdateBusGroup(c *gin.Context) {
	var busGroup system.BusGroup
	_ = c.ShouldBindJSON(&busGroup)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(busGroup, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := busGroupService.UpdateBusGroup(busGroup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBusGroup 用id查询BusGroup
// @Tags BusGroup
// @Summary 用id查询BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BusGroup true "用id查询BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busGroup/findBusGroup [get]
func (busGroupApi *BusGroupApi) FindBusGroup(c *gin.Context) {
	var busGroup system.BusGroup
	_ = c.ShouldBindQuery(&busGroup)
	if rebusGroup, err := busGroupService.GetBusGroup(busGroup.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebusGroup": rebusGroup}, c)
	}
}

// GetBusGroupList 分页获取BusGroup列表
// @Tags BusGroup
// @Summary 分页获取BusGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusGroupSearch true "分页获取BusGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGroup/getBusGroupList [get]
func (busGroupApi *BusGroupApi) GetBusGroupList(c *gin.Context) {
	var pageInfo businessReq.BusGroupSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := busGroupService.GetBusGroupInfoList(pageInfo); err != nil {
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
