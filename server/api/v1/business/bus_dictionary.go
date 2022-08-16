package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusDictionaryApi struct {
}

func (busDictionaryApi *BusDictionaryApi) FindBusDictionary(c *gin.Context) {
	var dictInfo businessReq.BusDictionary // 查哪个字典
	_ = c.ShouldBindQuery(&dictInfo)
	if dictInfo.DictName == "provider" {
		if rebusProvidersDict, err := busProviderService.GetBusProviderDict(); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"rebusDict": rebusProvidersDict}, c)
		}
	} else if dictInfo.DictName == "group" {
		if rebusGroupDict, err := busGroupService.GetBusGroupDict(); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"rebusDict": rebusGroupDict}, c)
		}
	} else if dictInfo.DictName == "factory" {
		if rebusFactoryDict, err := busFactoryService.GetBusFactoryDict(); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"rebusDict": rebusFactoryDict}, c)
		}
	} else if dictInfo.DictName == "applyState" {
		if rebusApplyStateDict, err := busOrderService.GetBusStateDict(); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"rebusDict": rebusApplyStateDict}, c)
		}
	} else {

	}
}
