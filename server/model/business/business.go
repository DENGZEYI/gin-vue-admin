package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
	"time"
)

// BusOrder 结构体
type BusOrder struct {
	global.GVA_MODEL
	ApplicantID uint
	Applicant   system.SysUser `gorm:"foreignKey:ApplicantID"`
	ApproverID  uint
	Approver    system.SysUser `gorm:"foreignKey:ApproverID"`
	State       uint           `json:"state" form:"state"`
	Goods       []BusGoods     `gorm:"foreignKey:BusOrderID"`
}

// TableName BusOrder 表名
func (BusOrder) TableName() string {
	return "bus_order"
}

// BusGoodsDict 结构体
// 记录了耗材的基本信息
type BusGoodsDict struct {
	ID            uint           `gorm:"primarykey"` // 主键ID
	CreatedAt     time.Time      // 创建时间
	UpdatedAt     time.Time      // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
	Name          string         `json:"name" form:"name" `
	Price         int            `json:"price" form:"price" `
	Factory       string         `json:"factory" form:"factory" `
	Specification string         `json:"specification" form:"specification" `
	GroupID       *uint
	Group         BusGroup `gorm:"foreignKey:GroupID"`
	ProviderID    *uint
	Provider      BusProvider `gorm:"foreignKey:ProviderID"`
}

// TableName BusGoods 表名
func (BusGoodsDict) TableName() string {
	return "bus_goods_dict"
}

// BusGoods 结构体
type BusGoods struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间

	BusOrderID   uint
	GoodsDict    BusGoodsDict `gorm:"foreignKey:GoodsDictID"`
	GoodsDictID  *uint        `json:"GoodsDictID" form:"GoodsDictID" `
	SerialNumber uint         `json:"serial_number" form:"serial_number" `
	Batch        string       `json:"batch" form:"batch" `
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
