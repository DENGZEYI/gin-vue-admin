package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestDict2Router struct {
}

// InitTestDict2Router 初始化 TestDict2 路由信息
func (s *TestDict2Router) InitTestDict2Router(Router *gin.RouterGroup) {
	testDict2Router := Router.Group("testDict2").Use(middleware.OperationRecord())
	testDict2RouterWithoutRecord := Router.Group("testDict2")
	var testDict2Api = v1.ApiGroupApp.TestApiGroup.TestDict2Api
	{
		testDict2Router.POST("createTestDict2", testDict2Api.CreateTestDict2)             // 新建TestDict2
		testDict2Router.DELETE("deleteTestDict2", testDict2Api.DeleteTestDict2)           // 删除TestDict2
		testDict2Router.DELETE("deleteTestDict2ByIds", testDict2Api.DeleteTestDict2ByIds) // 批量删除TestDict2
		testDict2Router.PUT("updateTestDict2", testDict2Api.UpdateTestDict2)              // 更新TestDict2
	}
	{
		testDict2RouterWithoutRecord.GET("findTestDict2", testDict2Api.FindTestDict2)       // 根据ID获取TestDict2
		testDict2RouterWithoutRecord.GET("getTestDict2List", testDict2Api.GetTestDict2List) // 获取TestDict2列表
	}
}
