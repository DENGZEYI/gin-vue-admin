package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	. "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

type result struct {
	ID   *uint
	Name string `gorm:"column:nick_name"`
}

func Test1(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gva?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	entities := []SysBaseMenu{
		// 字典
		{GVA_MODEL: global.GVA_MODEL{ID: 100}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "busDict", Name: "busDict", Component: "", Sort: 0, Meta: Meta{Title: "字典", Icon: "message"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 101}, MenuLevel: 0, Hidden: false, ParentId: "100", Path: "busGoodsDict", Name: "busGoodsDict", Component: "view/busGoodsDict/busGoodsDict.vue", Sort: 0, Meta: Meta{Title: "耗材字典", Icon: "message"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 102}, MenuLevel: 0, Hidden: false, ParentId: "100", Path: "busFactory", Name: "busFactory", Component: "view/busFactory/busFactory.vue", Sort: 0, Meta: Meta{Title: "生产厂商字典", Icon: "message"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 103}, MenuLevel: 0, Hidden: false, ParentId: "100", Path: "busProvider", Name: "busProvider", Component: "view/busProvider/busProvider.vue", Sort: 0, Meta: Meta{Title: "供应商字典", Icon: "message"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 104}, MenuLevel: 0, Hidden: false, ParentId: "100", Path: "busGroup", Name: "busGroup", Component: "view/busGroup/busGroup.vue", Sort: 0, Meta: Meta{Title: "组别字典", Icon: "message"}},
		// 申请

		// 申请单的审批、采购
		{GVA_MODEL: global.GVA_MODEL{ID: 201}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "busOrder", Name: "busOrder", Component: "view/busOrder/busOrder.vue", Sort: 0, Meta: Meta{Title: "申请单", Icon: "message"}},
		// 入库
		{GVA_MODEL: global.GVA_MODEL{ID: 301}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "busIngress", Name: "busIngress", Component: "view/busIngress/busIngress.vue", Sort: 0, Meta: Meta{Title: "入库", Icon: "message"}},
	}
	if err = db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entities).Error; err != nil {
		print("失败")
		return
	}
	print("成功")
}
