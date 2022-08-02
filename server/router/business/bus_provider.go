package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusProviderRouter struct {
}

// InitBusProviderRouter 初始化 BusProvider 路由信息
func (s *BusProviderRouter) InitBusProviderRouter(Router *gin.RouterGroup) {
	busProviderRouter := Router.Group("busProvider").Use(middleware.OperationRecord())
	busProviderRouterWithoutRecord := Router.Group("busProvider")
	var busProviderApi = v1.ApiGroupApp.BusinessApiGroup.BusProviderApi
	{
		busProviderRouter.POST("createBusProvider", busProviderApi.CreateBusProvider)             // 新建BusProvider
		busProviderRouter.DELETE("deleteBusProvider", busProviderApi.DeleteBusProvider)           // 删除BusProvider
		busProviderRouter.DELETE("deleteBusProviderByIds", busProviderApi.DeleteBusProviderByIds) // 批量删除BusProvider
		busProviderRouter.PUT("updateBusProvider", busProviderApi.UpdateBusProvider)              // 更新BusProvider
	}
	{
		busProviderRouterWithoutRecord.GET("findBusProvider", busProviderApi.FindBusProvider)       // 根据ID获取BusProvider
		busProviderRouterWithoutRecord.GET("getBusProviderList", busProviderApi.GetBusProviderList) // 获取BusProvider列表
	}
}
