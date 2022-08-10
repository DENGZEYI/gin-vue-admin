package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusOrderSearch struct {
	business.BusOrder
	request.PageInfo
}

type BusOrderDetailsRst struct {
	GoodsDict     business.BusGoodsDict `json:"GoodsDict" form:"GoodsDict" gorm:"foreignKey:Goods_dict_id"`
	Goods_dict_id int                   `json:"goods_dict_id" form:"goods_dict_id"`
	Number        int                   `json:"number" form:"number"`
}
