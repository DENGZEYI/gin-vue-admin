package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
	"time"
)

type BusOrderDetails struct {
	global.GVA_MODEL
	BusOrderID  *uint
	GoodsDict   BusGoodsDict `gorm:"foreignKey:GoodsDictID"`
	GoodsDictID *uint        `json:"goods_dict_id" form:"goods_dict_id" `
	Number      uint         `json:"number" form:"number" `
}

// TableName BusOrderDetails 表名
func (BusOrderDetails) TableName() string {
	return "bus_order_details"
}

// BusOrder 结构体
type BusOrder struct {
	global.GVA_MODEL
	ApplicantID     uint
	Applicant       system.SysUser `gorm:"foreignKey:ApplicantID"`
	ApproverID      uint
	Approver        system.SysUser    `gorm:"foreignKey:ApproverID"`
	State           uint              `json:"state" form:"state"`
	BusOrderDetails []BusOrderDetails `gorm:"foreignKey:BusOrderID" json:"bus_order_details" form:"bus_order_details"`
	Note            string            `json:"note" form:"note"` // 备注
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
	Unit          string         `json:"unit" form:"unit" `
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

	BusOrder    BusOrder     `gorm:"foreignKey:BusOrderID"`
	BusOrderID  *uint        `json:"bus_order_id" form:"bus_order_id" `
	GoodsDict   BusGoodsDict `gorm:"foreignKey:GoodsDictID"`
	GoodsDictID *uint        `json:"goods_dict_id" form:"goods_dict_id" `

	SerialNumber   uint      `json:"serial_number" form:"serial_number" `
	Batch          string    `json:"batch" form:"batch" `
	ExpirationDate time.Time `json:"expiration_date" form:"expiration_date" ` // 有效期
	ContractCode   string    `json:"contract_code" form:"contract_code"`      // 合同代码
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
	Name          string `json:"name" form:"name" `                     // 供应商名称
	Corporation   string `json:"corporation" form:"corporation" `       // 法人姓名
	CorporationID string `json:"corporation_id" form:"corporation_id" ` // 法人身份证
	Agent         string `json:"agent" form:"agent" `                   // 供应商经办人姓名
	AgentID       string `json:"agent_id" form:"agent_id" `             // 供应商经办人身份证
	Telephone     int    `json:"telephone" form:"telephone" `
}

// TableName BusProvider 表名
func (BusProvider) TableName() string {
	return "bus_provider"
}
