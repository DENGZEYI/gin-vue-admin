package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test"
)

type TestDict2Search struct {
	test.TestDict2
	request.PageInfo
}
