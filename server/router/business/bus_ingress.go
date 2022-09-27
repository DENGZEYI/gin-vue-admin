package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusIngressRouter struct {
}

// InitBusIngressRouter 初始化 BusIngress 路由信息
func (s *BusIngressRouter) InitBusIngressRouter(Router *gin.RouterGroup) {
	busIngressRouter := Router.Group("busIngress").Use(middleware.OperationRecord())
	busIngressRouterWithoutRecord := Router.Group("busIngress")
	var busIngressApi = v1.ApiGroupApp.BusinessApiGroup.BusIngressApi
	{
		busIngressRouter.POST("ingress", busIngressApi.Ingress) // 入库
	}
	{
		busIngressRouterWithoutRecord.GET("getBusIngressList", busIngressApi.GetBusIngressList) // 根据ID获取BusOrderDetail
	}
}
