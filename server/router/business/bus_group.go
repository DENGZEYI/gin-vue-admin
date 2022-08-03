package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusGroupRouter struct {
}

// InitBusGroupRouter 初始化 BusGroup 路由信息
func (s *BusGroupRouter) InitBusGroupRouter(Router *gin.RouterGroup) {
	busGroupRouter := Router.Group("busGroup").Use(middleware.OperationRecord())
	busGroupRouterWithoutRecord := Router.Group("busGroup")
	var busGroupApi = v1.ApiGroupApp.BusinessApiGroup.BusGroupApi
	{
		busGroupRouter.POST("createBusGroup", busGroupApi.CreateBusGroup)             // 新建BusGroup
		busGroupRouter.DELETE("deleteBusGroup", busGroupApi.DeleteBusGroup)           // 删除BusGroup
		busGroupRouter.DELETE("deleteBusGroupByIds", busGroupApi.DeleteBusGroupByIds) // 批量删除BusGroup
		busGroupRouter.PUT("updateBusGroup", busGroupApi.UpdateBusGroup)              // 更新BusGroup
	}
	{
		busGroupRouterWithoutRecord.GET("findBusGroup", busGroupApi.FindBusGroup)       // 根据ID获取BusGroup
		busGroupRouterWithoutRecord.GET("getBusGroupList", busGroupApi.GetBusGroupList) // 获取BusGroup列表
	}
}
