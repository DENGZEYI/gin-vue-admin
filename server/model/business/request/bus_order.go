package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusOrderSearch struct {
	business.BusOrder
	request.PageInfo
}
