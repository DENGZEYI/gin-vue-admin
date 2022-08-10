package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusGoodsDictSearch struct {
	business.BusGoodsDict
	request.PageInfo
}

type BusApplyGoodsByID struct {
	ID     *uint `json:"ID" form:"ID"`
	Number int   `json:"number" form:"number"`
}

type BusApplyInfo struct {
	Ids []BusApplyGoodsByID `json:"ids" form:"ids"`
}
