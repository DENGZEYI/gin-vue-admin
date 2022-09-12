package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusIngressApi struct {
}

var busIngressService = service.ServiceGroupApp.BusinessServiceGroup.BusIngressService

// GetBusIngressList 分页获取BusIngress列表
// @accept application/json
// @Produce application/json
// @Param data query businessReq.BusGroupSearch true "分页获取BusGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGroup/getBusGroupList [get]
func (busIngressApi *BusIngressApi) GetBusIngressList(c *gin.Context) {
	var pageInfo businessReq.BusIngressSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := busIngressService.GetBusIngressInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
