package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusGoodsDictRouter struct {
}

// InitBusGoodsRouter 初始化 BusGoods 路由信息
func (s *BusGoodsDictRouter) InitBusGoodsRouter(Router *gin.RouterGroup) {
	busGoodsRouter := Router.Group("busGoodsDict").Use(middleware.OperationRecord())
	busGoodsRouterWithoutRecord := Router.Group("busGoodsDict")
	var busGoodsApi = v1.ApiGroupApp.BusinessApiGroup.BusGoodsDictApi
	{
		busGoodsRouter.POST("createBusGoodsDict", busGoodsApi.CreateBusGoodsDict)             // 新建BusGoods
		busGoodsRouter.DELETE("deleteBusGoodsDict", busGoodsApi.DeleteBusGoodsDict)           // 删除BusGoods
		busGoodsRouter.DELETE("deleteBusGoodsDictByIds", busGoodsApi.DeleteBusGoodsDictByIds) // 批量删除BusGoods
		busGoodsRouter.PUT("updateBusGoodsDict", busGoodsApi.UpdateBusGoodsDict)              // 更新BusGoods
		busGoodsRouter.POST("applyBusGoodsByIds", busGoodsApi.ApplyBusGoodsByIds)
	}
	{
		busGoodsRouterWithoutRecord.GET("findBusGoodsDict", busGoodsApi.FindBusGoodsDict)       // 根据ID获取BusGoods
		busGoodsRouterWithoutRecord.GET("getBusGoodsDictList", busGoodsApi.GetBusGoodsDictList) // 获取BusGoods列表
	}
}
