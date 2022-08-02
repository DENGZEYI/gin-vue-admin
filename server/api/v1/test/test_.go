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

type TestApi struct {
}

var testService = service.ServiceGroupApp.TestServiceGroup.TestService

// CreateTest 创建Test
// @Tags Test
// @Summary 创建Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Test true "创建Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test/createTest [post]
func (testApi *TestApi) CreateTest(c *gin.Context) {
	var test test.Test
	_ = c.ShouldBindJSON(&test)
	if err := testService.CreateTest(test); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTest 删除Test
// @Tags Test
// @Summary 删除Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Test true "删除Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test/deleteTest [delete]
func (testApi *TestApi) DeleteTest(c *gin.Context) {
	var test test.Test
	_ = c.ShouldBindJSON(&test)
	if err := testService.DeleteTest(test); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestByIds 批量删除Test
// @Tags Test
// @Summary 批量删除Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /test/deleteTestByIds [delete]
func (testApi *TestApi) DeleteTestByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := testService.DeleteTestByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTest 更新Test
// @Tags Test
// @Summary 更新Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Test true "更新Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /test/updateTest [put]
func (testApi *TestApi) UpdateTest(c *gin.Context) {
	var test test.Test
	_ = c.ShouldBindJSON(&test)
	if err := testService.UpdateTest(test); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTest 用id查询Test
// @Tags Test
// @Summary 用id查询Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test.Test true "用id查询Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /test/findTest [get]
func (testApi *TestApi) FindTest(c *gin.Context) {
	var test test.Test
	_ = c.ShouldBindQuery(&test)
	if retest, err := testService.GetTest(test.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retest": retest}, c)
	}
}

// GetTestList 分页获取Test列表
// @Tags Test
// @Summary 分页获取Test列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query testReq.TestSearch true "分页获取Test列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test/getTestList [get]
func (testApi *TestApi) GetTestList(c *gin.Context) {
	var pageInfo testReq.TestSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := testService.GetTestInfoList(pageInfo); err != nil {
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
