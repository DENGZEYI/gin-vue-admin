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
	Name   string `json:"name" form:"name" `
	Number int    `json:"number" form:"number" `
}
