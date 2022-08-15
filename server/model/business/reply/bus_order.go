package reply

import "github.com/flipped-aurora/gin-vue-admin/server/model/business"

type ReBusOrderDetail struct {
	GoodsDict       business.BusGoodsDict `json:"goods_dict" form:"goods_dict" `
	Number          uint                  `json:"number" form:"number" ` // 申请的数量
	NumberAlreadyIn uint                  `json:"number_already_in" form:"number_already_in"`
}
type ReBusOrderDetails struct {
	BusOrderDetails []ReBusOrderDetail `json:"bus_order_details" form:"bus_order_details_f" `
}
