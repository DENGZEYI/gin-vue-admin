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

type TestDict2Api struct {
}

var testDict2Service = service.ServiceGroupApp.TestServiceGroup.TestDict2Service

// CreateTestDict2 创建TestDict2
// @Tags TestDict2
// @Summary 创建TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.TestDict2 true "创建TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testDict2/createTestDict2 [post]
func (testDict2Api *TestDict2Api) CreateTestDict2(c *gin.Context) {
	var testDict2 test.TestDict2
	_ = c.ShouldBindJSON(&testDict2)
	if err := testDict2Service.CreateTestDict2(testDict2); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestDict2 删除TestDict2
// @Tags TestDict2
// @Summary 删除TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.TestDict2 true "删除TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testDict2/deleteTestDict2 [delete]
func (testDict2Api *TestDict2Api) DeleteTestDict2(c *gin.Context) {
	var testDict2 test.TestDict2
	_ = c.ShouldBindJSON(&testDict2)
	if err := testDict2Service.DeleteTestDict2(testDict2); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestDict2ByIds 批量删除TestDict2
// @Tags TestDict2
// @Summary 批量删除TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testDict2/deleteTestDict2ByIds [delete]
func (testDict2Api *TestDict2Api) DeleteTestDict2ByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := testDict2Service.DeleteTestDict2ByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestDict2 更新TestDict2
// @Tags TestDict2
// @Summary 更新TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.TestDict2 true "更新TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testDict2/updateTestDict2 [put]
func (testDict2Api *TestDict2Api) UpdateTestDict2(c *gin.Context) {
	var testDict2 test.TestDict2
	_ = c.ShouldBindJSON(&testDict2)
	if err := testDict2Service.UpdateTestDict2(testDict2); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestDict2 用id查询TestDict2
// @Tags TestDict2
// @Summary 用id查询TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test.TestDict2 true "用id查询TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testDict2/findTestDict2 [get]
func (testDict2Api *TestDict2Api) FindTestDict2(c *gin.Context) {
	var testDict2 test.TestDict2
	_ = c.ShouldBindQuery(&testDict2)
	if retestDict2, err := testDict2Service.GetTestDict2(testDict2.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestDict2": retestDict2}, c)
	}
}

// GetTestDict2List 分页获取TestDict2列表
// @Tags TestDict2
// @Summary 分页获取TestDict2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query testReq.TestDict2Search true "分页获取TestDict2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testDict2/getTestDict2List [get]
func (testDict2Api *TestDict2Api) GetTestDict2List(c *gin.Context) {
	var pageInfo testReq.TestDict2Search
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := testDict2Service.GetTestDict2InfoList(pageInfo); err != nil {
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
