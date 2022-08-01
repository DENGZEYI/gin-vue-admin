package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Test struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Address string `json:"address" form:"address" gorm:"column:address;comment:;"`
	Age     int    `json:"age" form:"age" gorm:"column:age;comment:;"`
}

func Test1(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(Test{})
	if err != nil {
		return
	}
}
