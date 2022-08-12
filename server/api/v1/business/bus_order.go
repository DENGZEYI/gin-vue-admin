package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusOrderApi struct {
}

var busOrderService = service.ServiceGroupApp.BusinessServiceGroup.BusOrderService

// CreateBusOrder 创建BusOrder
// @Tags BusOrder
// @Summary 创建BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusOrder true "创建BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busOrder/createBusOrder [post]
func (busOrderApi *BusOrderApi) CreateBusOrder(c *gin.Context) {
	var busOrder business.BusOrder
	_ = c.ShouldBindJSON(&busOrder)
	if err := busOrderService.CreateBusOrder(busOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBusOrder 删除BusOrder
// @Tags BusOrder
// @Summary 删除BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusOrder true "删除BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busOrder/deleteBusOrder [delete]
func (busOrderApi *BusOrderApi) DeleteBusOrder(c *gin.Context) {
	var busOrder business.BusOrder
	_ = c.ShouldBindJSON(&busOrder)
	if err := busOrderService.DeleteBusOrder(busOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBusOrderByIds 批量删除BusOrder
// @Tags BusOrder
// @Summary 批量删除BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /busOrder/deleteBusOrderByIds [delete]
func (busOrderApi *BusOrderApi) DeleteBusOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := busOrderService.DeleteBusOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBusOrder 更新BusOrder
// @Tags BusOrder
// @Summary 更新BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusOrder true "更新BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busOrder/updateBusOrder [put]
func (busOrderApi *BusOrderApi) UpdateBusOrder(c *gin.Context) {
	var busOrder business.BusOrder
	_ = c.ShouldBindJSON(&busOrder)
	if err := busOrderService.UpdateBusOrder(busOrder, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBusOrder 用id查询BusOrder
// @Tags BusOrder
// @Summary 用id查询BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BusOrder true "用id查询BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busOrder/findBusOrder [get]
func (busOrderApi *BusOrderApi) FindBusOrder(c *gin.Context) {
	var busOrder business.BusOrder
	_ = c.ShouldBindQuery(&busOrder)
	if rebusOrder, err := busOrderService.GetBusOrder(busOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebusOrder": rebusOrder}, c)
	}
}

// GetOrderDetails 根据ID获取订单详情
func (busOrderApi *BusOrderApi) GetOrderDetails(c *gin.Context) {
	var busOrder business.BusOrder
	_ = c.ShouldBindQuery(&busOrder)
	if rebusOrder, err := busOrderService.GetBusOrderDetails(busOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebusOrder": rebusOrder}, c)
	}
}

// GetBusOrderList 分页获取BusOrder列表
// @Tags BusOrder
// @Summary 分页获取BusOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusOrderSearch true "分页获取BusOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busOrder/getBusOrderList [get]
func (busOrderApi *BusOrderApi) GetBusOrderList(c *gin.Context) {
	var pageInfo businessReq.BusOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := busOrderService.GetBusOrderInfoList(pageInfo); err != nil {
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
