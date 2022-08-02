// 自动生成模板Test
package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Test 结构体
type Test struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Address string `json:"address" form:"address" gorm:"column:address;comment:;"`
	Age     *int   `json:"age" form:"age" gorm:"column:age;comment:;"`
}

// TableName Test 表名
func (Test) TableName() string {
	return "test"
}
