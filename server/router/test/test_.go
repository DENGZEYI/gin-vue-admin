package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

// InitTestRouter 初始化 Test 路由信息
func (s *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test").Use(middleware.OperationRecord())
	testRouterWithoutRecord := Router.Group("test")
	var testApi = v1.ApiGroupApp.TestApiGroup.TestApi
	{
		testRouter.POST("createTest", testApi.CreateTest)             // 新建Test
		testRouter.DELETE("deleteTest", testApi.DeleteTest)           // 删除Test
		testRouter.DELETE("deleteTestByIds", testApi.DeleteTestByIds) // 批量删除Test
		testRouter.PUT("updateTest", testApi.UpdateTest)              // 更新Test
	}
	{
		testRouterWithoutRecord.GET("findTest", testApi.FindTest)       // 根据ID获取Test
		testRouterWithoutRecord.GET("getTestList", testApi.GetTestList) // 获取Test列表
	}
}
