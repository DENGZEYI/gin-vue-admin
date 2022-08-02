// 自动生成模板TestDic
package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TestDic 结构体
type TestDic struct {
	global.GVA_MODEL
	Group *int   `json:"group" form:"group" gorm:"column:group;comment:;"`
	Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
}

// TableName TestDic 表名
func (TestDic) TableName() string {
	return "test_dic"
}
