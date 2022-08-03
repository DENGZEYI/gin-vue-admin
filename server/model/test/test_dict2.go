// 自动生成模板TestDict2
package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TestDict2 结构体
type TestDict2 struct {
	global.GVA_MODEL
	Group *int   `json:"group" form:"group" gorm:"column:group;comment:;"`
	Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
}

// TableName TestDict2 表名
func (TestDict2) TableName() string {
	return "test_dict2"
}
