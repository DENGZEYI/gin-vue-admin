package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BusOrderSearch struct {
	business.BusOrder
	request.PageInfo
}

type BusOrderDetailsRst struct {
	GoodsDict   business.BusGoodsDict `json:"GoodsDict" form:"GoodsDict" gorm:"foreignKey:Goods_dict_id"`
	GoodsDictID int                   `json:"goods_dict_id" form:"goods_dict_id"`
	Number      int                   `json:"number" form:"number"`
}

type BusIngressReq struct {
	IngressNumber  uint      `json:"ingress_number" form:"ingress_number"`
	GoodsDictID    *uint     `json:"goods_dict_id" form:"goods_dict_id"`
	BusOrderID     *uint     `json:"bus_order_id" form:"bus_order_id"`
	Batch          string    `json:"batch" form:"batch" `                     // 批号
	ExpirationDate time.Time `json:"expiration_date" form:"expiration_date" ` // 有效期
	ContractCode   string    `json:"contract_code" form:"contract_code"`      // 合同代码
}
