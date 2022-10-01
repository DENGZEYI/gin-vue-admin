package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusGoodsDictSearch struct {
	business.BusGoodsDict
	request.PageInfo
}

type ApplyDetail struct {
	GoodsDictID *uint `json:"goods_dict_id"`
	ApplyNumber uint  `json:"apply_number"`
}

type BusApplyReq struct {
	ApplyDetails []ApplyDetail `json:"apply_details"`
}
