package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BusIngressSearch struct {
	request.PageInfo
}

type IngressDetail struct {
	GoodsName      string    `json:"goods_name" form:"goods_name"`
	GoodsDictID    *uint     `json:"goods_dict_id"`
	Batch          string    `json:"batch"`
	IngressNumber  uint      `json:"ingress_number" form:"ingress_number"`
	ExpirationDate time.Time `json:"expiration_date" form:"expiration_date" ` // 有效期
}

// BusIngressReq 入库表单
type BusIngressReq struct {
	BusOrderID *uint `json:"order_id" form:"order_id"`
	//InvoiceNumber  string    `json:"invoice_number" form:"invoice_number"`    // 发票号
	DeliveryNumber string          `json:"delivery_number" form:"delivery_number" ` //送货单号
	IngressDate    time.Time       `json:"ingress_date" form:"ingress_date"`        // 入库时间
	IngressDetails []IngressDetail `json:"ingress_details"`
}
