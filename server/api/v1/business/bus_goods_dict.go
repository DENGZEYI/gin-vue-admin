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

type BusGoodsDictApi struct {
}

var BusGoodsDictService = service.ServiceGroupApp.BusinessServiceGroup.BusGoodsDictService
var AuthorityService = service.ServiceGroup{}.SystemServiceGroup.AuthorityService

// CreateBusGoodsDict 创建BusGoodsDict
// @Tags BusGoodsDict
// @Summary 创建BusGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusGoodsDict true "创建BusGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BusGoodsDict/createBusGoodsDict [post]
func (BusGoodsDictApi *BusGoodsDictApi) CreateBusGoodsDict(c *gin.Context) {
	var BusGoodsDict business.BusGoodsDict
	_ = c.ShouldBindJSON(&BusGoodsDict)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(BusGoodsDict, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := BusGoodsDictService.CreateBusGoodsDict(BusGoodsDict); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBusGoodsDict 删除BusGoodsDict
// @Tags BusGoodsDict
// @Summary 删除BusGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusGoodsDict true "删除BusGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /BusGoodsDict/deleteBusGoodsDict [delete]
func (BusGoodsDictApi *BusGoodsDictApi) DeleteBusGoodsDict(c *gin.Context) {
	var BusGoodsDict business.BusGoodsDict
	_ = c.ShouldBindJSON(&BusGoodsDict)
	if err := BusGoodsDictService.DeleteBusGoodsDict(BusGoodsDict); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBusGoodsDictByIds 批量删除BusGoodsDict
// @Tags BusGoodsDict
// @Summary 批量删除BusGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /BusGoodsDict/deleteBusGoodsDictByIds [delete]
func (BusGoodsDictApi *BusGoodsDictApi) DeleteBusGoodsDictByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := BusGoodsDictService.DeleteBusGoodsDictByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBusGoodsDict 更新BusGoodsDict
// @Tags BusGoodsDict
// @Summary 更新BusGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BusGoodsDict true "更新BusGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /BusGoodsDict/updateBusGoodsDict [put]
func (BusGoodsDictApi *BusGoodsDictApi) UpdateBusGoodsDict(c *gin.Context) {
	var BusGoodsDict business.BusGoodsDict
	_ = c.ShouldBindJSON(&BusGoodsDict)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(BusGoodsDict, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := BusGoodsDictService.UpdateBusGoodsDict(BusGoodsDict); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBusGoodsDict 用id查询BusGoodsDict
// @Tags BusGoodsDict
// @Summary 用id查询BusGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BusGoodsDict true "用id查询BusGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /BusGoodsDict/findBusGoodsDict [get]
func (BusGoodsDictApi *BusGoodsDictApi) FindBusGoodsDict(c *gin.Context) {
	var BusGoodsDict business.BusGoodsDict
	_ = c.ShouldBindQuery(&BusGoodsDict)
	if reBusGoodsDict, err := BusGoodsDictService.GetBusGoodsDict(BusGoodsDict.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reBusGoodsDict": reBusGoodsDict}, c)
	}
}

// GetBusGoodsDictList 分页获取BusGoodsDict列表
// @Tags BusGoodsDict
// @Summary 分页获取BusGoodsDict列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusGoodsDictSearch true "分页获取BusGoodsDict列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BusGoodsDict/getBusGoodsDictList [get]
func (BusGoodsDictApi *BusGoodsDictApi) GetBusGoodsDictList(c *gin.Context) {
	var pageInfo businessReq.BusGoodsDictSearch
	// 获取用户角色
	authID := utils.GetUserAuthorityId(c)
	auth, err := AuthorityService.FindAuthorityByID(&authID)
	if err != nil {
		global.GVA_LOG.Error("获取用户角色失败!", zap.Error(err))
		response.FailWithMessage("获取用户角色失败", c)
		return
	}
	// 查询耗材
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := BusGoodsDictService.GetBusGoodsDictInfoList(pageInfo, auth); err != nil {
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

// ApplyBusGoodsByIds 申请Goods
// @Tags BusGoodsDict
// @Summary 分页获取BusGoodsDict列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusGoodsDictSearch true "分页获取BusGoodsDict列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"申请成功"}"
// @Router /BusGoodsDict/applyGoodsByIds [post]
func (BusGoodsDictApi *BusGoodsDictApi) ApplyBusGoodsByIds(c *gin.Context) {
	var applyInfo businessReq.BusApplyInfo
	_ = c.ShouldBindJSON(&applyInfo)
	if err := BusGoodsDictService.ApplyGoodsByIds(applyInfo, c); err != nil {
		global.GVA_LOG.Error("申请失败!", zap.Error(err))
		response.FailWithMessage("申请失败", c)
	} else {
		response.OkWithMessage("申请成功", c)
	}
}
