package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusGoodsRouter struct {
}

// InitBusGoodsRouter 初始化 BusGoods 路由信息
func (s *BusGoodsRouter) InitBusGoodsRouter(Router *gin.RouterGroup) {
	busGoodsRouter := Router.Group("busGoods").Use(middleware.OperationRecord())
	busGoodsRouterWithoutRecord := Router.Group("busGoods")
	var busGoodsApi = v1.ApiGroupApp.BusinessApiGroup.BusGoodsApi
	{
		busGoodsRouter.POST("createBusGoods", busGoodsApi.CreateBusGoods)             // 新建BusGoods
		busGoodsRouter.DELETE("deleteBusGoods", busGoodsApi.DeleteBusGoods)           // 删除BusGoods
		busGoodsRouter.DELETE("deleteBusGoodsByIds", busGoodsApi.DeleteBusGoodsByIds) // 批量删除BusGoods
		busGoodsRouter.PUT("updateBusGoods", busGoodsApi.UpdateBusGoods)              // 更新BusGoods
		busGoodsRouter.POST("applyGoodsByIds", busGoodsApi.ApplyGoodsByIds)
	}
	{
		busGoodsRouterWithoutRecord.GET("findBusGoods", busGoodsApi.FindBusGoods)       // 根据ID获取BusGoods
		busGoodsRouterWithoutRecord.GET("getBusGoodsList", busGoodsApi.GetBusGoodsList) // 获取BusGoods列表
	}
}
