package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BusDictionaryRouter struct {
}

// InitBusDictionaryRouter 初始化 BusDictionary 路由信息
func (s *BusDictionaryRouter) InitBusDictionaryRouter(Router *gin.RouterGroup) {
	busDictionaryRouterWithoutRecord := Router.Group("busDictionary")
	var BusDictionaryApi = v1.ApiGroupApp.BusinessApiGroup.BusDictionaryApi
	{
		busDictionaryRouterWithoutRecord.GET("findBusDictionary", BusDictionaryApi.FindBusDictionary) // 字典
	}
}
