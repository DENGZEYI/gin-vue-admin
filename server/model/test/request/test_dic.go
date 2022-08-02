package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test"
)

type TestDicSearch struct {
	test.TestDic
	request.PageInfo
}
