package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BusIngressSearch struct {
	request.PageInfo
}
type BusIngressReq struct {
	IngressNumber  uint      `json:"ingress_number" form:"ingress_number"`
	GoodsDictID    *uint     `json:"goods_dict_id" form:"goods_dict_id"`
	BusOrderID     *uint     `json:"bus_order_id" form:"bus_order_id"`
	Batch          string    `json:"batch" form:"batch" `                     // 批号
	ExpirationDate time.Time `json:"expiration_date" form:"expiration_date" ` // 有效期
	InvoiceNumber  string    `json:"invoice_number" form:"invoice_number"`    // 发票号
	DeliveryNumber string    `json:"delivery_number" form:"delivery_number" ` //送货单号
}
