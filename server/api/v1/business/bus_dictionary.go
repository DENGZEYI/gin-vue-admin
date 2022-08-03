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

// todo 改注释

// FindBusProvider 用id查询BusProvider
// @Tags BusProvider
// @Summary 用id查询BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BusProvider true "用id查询BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busProvider/findBusProvider [get]
func (busProviderApi *BusDictionaryApi) FindBusDictionary(c *gin.Context) {
	var dictInfo businessReq.BusDictionary // 查哪个字典
	_ = c.ShouldBindQuery(&dictInfo)
	if dictInfo.DictName == "provider" {
		if rebusProvidersDict, err := busProviderService.GetBusProviderDict(); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"rebusProvidersDict": rebusProvidersDict}, c)
		}
	} else if dictInfo.DictName == "group" {
		if rebusGroupDict, err := busGroupService.GetBusGroupDict(); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"rebusProvidersDict": rebusGroupDict}, c)
		}
	} else if dictInfo.DictName == "goods" {

	} else {

	}
}
