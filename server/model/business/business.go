package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// BusOrder 结构体
type BusOrder struct {
	global.GVA_MODEL
	ApplicantID int
	Applicant   system.SysUser `gorm:"foreignKey:ApplicantID"`
	ApproverID  int
	Approver    system.SysUser `gorm:"foreignKey:ApproverID"`
	State       int            `json:"state" form:"state"`
	Goods       []BusGoods     `gorm:"many2many:bus_order_goods;"`
}

// TableName BusOrder 表名
func (BusOrder) TableName() string {
	return "bus_order"
}

type BusOrderGoods struct {
	BusOrderID int    `gorm:"primaryKey"`
	BusGoodsID int    `gorm:"primaryKey"`
	Number     int    `json:"number" form:"number" `
	Batch      string `json:"batch" form:"batch" `
}

// TableName BusOrderGoods 表名
func (BusOrderGoods) TableName() string {
	return "bus_order_goods"
}

// BusGoods 结构体
type BusGoods struct {
	global.GVA_MODEL
	Name          string `json:"name" form:"name" `
	Price         int    `json:"price" form:"price" `
	Factory       string `json:"factory" form:"factory" `
	Specification string `json:"specification" form:"specification" `
	GroupID       *int
	Group         BusGroup `gorm:"foreignKey:GroupID"`
	ProviderID    *int
	Provider      BusProvider `gorm:"foreignKey:ProviderID"`
}

// TableName BusGoods 表名
func (BusGoods) TableName() string {
	return "bus_goods"
}

// BusGroup 结构体
type BusGroup struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"`
}

// TableName BusGroup 表名
func (BusGroup) TableName() string {
	return "bus_group"
}

// BusProvider 结构体
type BusProvider struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" `
	Telephone int    `json:"telephone" form:"telephone" `
}

// TableName BusProvider 表名
func (BusProvider) TableName() string {
	return "bus_provider"
}
