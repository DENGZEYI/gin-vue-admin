package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
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
	dsn := "root:123456@tcp(127.0.0.1:3306)/gva?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	var rst []businessReq.BusOrderDetailsRst
	db = db.Model(&business.BusGoods{})
	orderID := 5
	db = db.Preload("GoodsDict").Select("goods_dict_id,count(goods_dict_id) as number").Where("bus_order_id =?", orderID).Group("goods_dict_id").Find(&rst)
	a := 1
	print(a)
}
