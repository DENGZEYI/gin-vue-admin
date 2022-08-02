package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestDicRouter struct {
}

// InitTestDicRouter 初始化 TestDic 路由信息
func (s *TestDicRouter) InitTestDicRouter(Router *gin.RouterGroup) {
	testDicRouter := Router.Group("testDic").Use(middleware.OperationRecord())
	testDicRouterWithoutRecord := Router.Group("testDic")
	var testDicApi = v1.ApiGroupApp.TestApiGroup.TestDicApi
	{
		testDicRouter.POST("createTestDic", testDicApi.CreateTestDic)             // 新建TestDic
		testDicRouter.DELETE("deleteTestDic", testDicApi.DeleteTestDic)           // 删除TestDic
		testDicRouter.DELETE("deleteTestDicByIds", testDicApi.DeleteTestDicByIds) // 批量删除TestDic
		testDicRouter.PUT("updateTestDic", testDicApi.UpdateTestDic)              // 更新TestDic
	}
	{
		testDicRouterWithoutRecord.GET("findTestDic", testDicApi.FindTestDic)       // 根据ID获取TestDic
		testDicRouterWithoutRecord.GET("getTestDicList", testDicApi.GetTestDicList) // 获取TestDic列表
	}
}
