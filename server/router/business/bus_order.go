package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusOrderRouter struct {
}

// InitBusOrderRouter 初始化 BusOrder 路由信息
func (s *BusOrderRouter) InitBusOrderRouter(Router *gin.RouterGroup) {
	busOrderRouter := Router.Group("busOrder").Use(middleware.OperationRecord())
	busOrderRouterWithoutRecord := Router.Group("busOrder")
	var busOrderApi = v1.ApiGroupApp.BusinessApiGroup.BusOrderApi
	{
		busOrderRouter.POST("createBusOrder", busOrderApi.CreateBusOrder)             // 新建BusOrder
		busOrderRouter.DELETE("deleteBusOrder", busOrderApi.DeleteBusOrder)           // 删除BusOrder
		busOrderRouter.DELETE("deleteBusOrderByIds", busOrderApi.DeleteBusOrderByIds) // 批量删除BusOrder
		busOrderRouter.PUT("updateBusOrder", busOrderApi.UpdateBusOrder)              // 更新BusOrder
		busOrderRouter.PUT("approveBusOrder", busOrderApi.ApproveBusOrder)            // 审批BusOrder
		busOrderRouter.PUT("ingressBusOrder", busOrderApi.IngressBusOrder)            // 审批BusOrder
	}
	{
		busOrderRouterWithoutRecord.GET("getOrderDetail", busOrderApi.GetOrderDetail)   // 根据ID获取BusOrderDetail
		busOrderRouterWithoutRecord.GET("findBusOrder", busOrderApi.FindBusOrder)       // 根据ID获取BusOrder
		busOrderRouterWithoutRecord.GET("getBusOrderList", busOrderApi.GetBusOrderList) // 获取BusOrder列表
		busOrderRouterWithoutRecord.GET("getOrderDetails", busOrderApi.GetOrderDetails)
	}
}