package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusFactoryRouter struct {
}

// InitBusFactoryRouter 初始化 BusFactory 路由信息
func (s *BusFactoryRouter) InitBusFactoryRouter(Router *gin.RouterGroup) {
	busFactoryRouter := Router.Group("busFactory").Use(middleware.OperationRecord())
	busFactoryRouterWithoutRecord := Router.Group("busFactory")
	var busFactoryApi = v1.ApiGroupApp.BusinessApiGroup.BusFactoryApi
	{
		busFactoryRouter.POST("createBusFactory", busFactoryApi.CreateBusFactory)             // 新建BusFactory
		busFactoryRouter.DELETE("deleteBusFactory", busFactoryApi.DeleteBusFactory)           // 删除BusFactory
		busFactoryRouter.DELETE("deleteBusFactoryByIds", busFactoryApi.DeleteBusFactoryByIds) // 批量删除BusFactory
		busFactoryRouter.PUT("updateBusFactory", busFactoryApi.UpdateBusFactory)              // 更新BusFactory
	}
	{
		busFactoryRouterWithoutRecord.GET("findBusFactory", busFactoryApi.FindBusFactory)       // 根据ID获取BusFactory
		busFactoryRouterWithoutRecord.GET("getBusFactoryList", busFactoryApi.GetBusFactoryList) // 获取BusFactory列表
	}
}
