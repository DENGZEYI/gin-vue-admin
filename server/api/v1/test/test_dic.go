package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test"
	testReq "github.com/flipped-aurora/gin-vue-admin/server/model/test/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestDicApi struct {
}

var testDicService = service.ServiceGroupApp.TestServiceGroup.TestDicService

// CreateTestDic 创建TestDic
// @Tags TestDic
// @Summary 创建TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.TestDic true "创建TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testDic/createTestDic [post]
func (testDicApi *TestDicApi) CreateTestDic(c *gin.Context) {
	var testDic test.TestDic
	_ = c.ShouldBindJSON(&testDic)
	if err := testDicService.CreateTestDic(testDic); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestDic 删除TestDic
// @Tags TestDic
// @Summary 删除TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.TestDic true "删除TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testDic/deleteTestDic [delete]
func (testDicApi *TestDicApi) DeleteTestDic(c *gin.Context) {
	var testDic test.TestDic
	_ = c.ShouldBindJSON(&testDic)
	if err := testDicService.DeleteTestDic(testDic); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestDicByIds 批量删除TestDic
// @Tags TestDic
// @Summary 批量删除TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testDic/deleteTestDicByIds [delete]
func (testDicApi *TestDicApi) DeleteTestDicByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := testDicService.DeleteTestDicByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestDic 更新TestDic
// @Tags TestDic
// @Summary 更新TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.TestDic true "更新TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testDic/updateTestDic [put]
func (testDicApi *TestDicApi) UpdateTestDic(c *gin.Context) {
	var testDic test.TestDic
	_ = c.ShouldBindJSON(&testDic)
	if err := testDicService.UpdateTestDic(testDic); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestDic 用id查询TestDic
// @Tags TestDic
// @Summary 用id查询TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test.TestDic true "用id查询TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testDic/findTestDic [get]
func (testDicApi *TestDicApi) FindTestDic(c *gin.Context) {
	var testDic test.TestDic
	_ = c.ShouldBindQuery(&testDic)
	if retestDic, err := testDicService.GetTestDic(testDic.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestDic": retestDic}, c)
	}
}

// GetTestDicList 分页获取TestDic列表
// @Tags TestDic
// @Summary 分页获取TestDic列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query testReq.TestDicSearch true "分页获取TestDic列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testDic/getTestDicList [get]
func (testDicApi *TestDicApi) GetTestDicList(c *gin.Context) {
	var pageInfo testReq.TestDicSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := testDicService.GetTestDicInfoList(pageInfo); err != nil {
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
