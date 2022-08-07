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

type BusGoodsApi struct {
}

var busGoodsService = service.ServiceGroupApp.BusinessServiceGroup.BusGoodsService

// CreateBusGoods 创建BusGoods
// @Tags BusGoods
// @Summary 创建BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusGoods true "创建BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGoods/createBusGoods [post]
func (busGoodsApi *BusGoodsApi) CreateBusGoods(c *gin.Context) {
	var busGoods business.BusGoods
	_ = c.ShouldBindJSON(&busGoods)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(busGoods, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := busGoodsService.CreateBusGoods(busGoods); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBusGoods 删除BusGoods
// @Tags BusGoods
// @Summary 删除BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusGoods true "删除BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busGoods/deleteBusGoods [delete]
func (busGoodsApi *BusGoodsApi) DeleteBusGoods(c *gin.Context) {
	var busGoods business.BusGoods
	_ = c.ShouldBindJSON(&busGoods)
	if err := busGoodsService.DeleteBusGoods(busGoods); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBusGoodsByIds 批量删除BusGoods
// @Tags BusGoods
// @Summary 批量删除BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /busGoods/deleteBusGoodsByIds [delete]
func (busGoodsApi *BusGoodsApi) DeleteBusGoodsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := busGoodsService.DeleteBusGoodsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBusGoods 更新BusGoods
// @Tags BusGoods
// @Summary 更新BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusGoods true "更新BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busGoods/updateBusGoods [put]
func (busGoodsApi *BusGoodsApi) UpdateBusGoods(c *gin.Context) {
	var busGoods business.BusGoods
	_ = c.ShouldBindJSON(&busGoods)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(busGoods, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := busGoodsService.UpdateBusGoods(busGoods); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBusGoods 用id查询BusGoods
// @Tags BusGoods
// @Summary 用id查询BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BusGoods true "用id查询BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busGoods/findBusGoods [get]
func (busGoodsApi *BusGoodsApi) FindBusGoods(c *gin.Context) {
	var busGoods business.BusGoods
	_ = c.ShouldBindQuery(&busGoods)
	if rebusGoods, err := busGoodsService.GetBusGoods(busGoods.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebusGoods": rebusGoods}, c)
	}
}

// GetBusGoodsList 分页获取BusGoods列表
// @Tags BusGoods
// @Summary 分页获取BusGoods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusGoodsSearch true "分页获取BusGoods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGoods/getBusGoodsList [get]
func (busGoodsApi *BusGoodsApi) GetBusGoodsList(c *gin.Context) {
	var pageInfo businessReq.BusGoodsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := busGoodsService.GetBusGoodsInfoList(pageInfo); err != nil {
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

// applyGoodsByIds 申请Goods
// @Tags BusGoods
// @Summary 分页获取BusGoods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusGoodsSearch true "分页获取BusGoods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"申请成功"}"
// @Router /busGoods/applyGoodsByIds [post]
func (busGoodsApi *BusGoodsApi) ApplyGoodsByIds(c *gin.Context) {
	var applyInfo businessReq.BusApplyInfo
	_ = c.ShouldBindJSON(&applyInfo)
	if err := busGoodsService.ApplyGoodsByIds(applyInfo, c); err != nil {
		global.GVA_LOG.Error("申请失败!", zap.Error(err))
		response.FailWithMessage("申请失败", c)
	} else {
		response.OkWithMessage("申请成功", c)
	}
}
