package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
	"time"
)

type BusOrderDetail struct {
	global.GVA_MODEL
	BusOrderID  *uint
	GoodsDict   BusGoodsDict `gorm:"foreignKey:GoodsDictID" json:"goods_dict" form:"goods_dict"`
	GoodsDictID *uint        `json:"goods_dict_id" form:"goods_dict_id" `
	Number      uint         `json:"number" form:"number" ` // 申请的数量
}

// TableName BusOrderDetails 表名
func (BusOrderDetail) TableName() string {
	return "bus_order_detail"
}

// BusOrder 结构体
type BusOrder struct {
	global.GVA_MODEL
	ApplicantID     uint
	Applicant       system.SysUser `gorm:"foreignKey:ApplicantID"`
	ApproverID      uint
	Approver        system.SysUser   `gorm:"foreignKey:ApproverID"`
	State           uint             `json:"state" form:"state"`
	BusOrderDetails []BusOrderDetail `gorm:"foreignKey:BusOrderID" json:"bus_order_details" form:"bus_order_details"`
	Note            string           `json:"note" form:"note"` // 备注
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

type BusIngressDetail struct {
	global.GVA_MODEL
	BusIngressID *uint
	GoodsDict    BusGoodsDict `gorm:"foreignKey:GoodsDictID" json:"goods_dict" form:"goods_dict"`
	GoodsDictID  *uint        `json:"goods_dict_id" form:"goods_dict_id" `
	Number       uint         `json:"number" form:"number" ` // 数量
}

// TableName BusIngressDetail 表名
func (BusIngressDetail) TableName() string {
	return "bus_ingress_detail"
}

// BusIngress 结构体
type BusIngress struct {
	global.GVA_MODEL
	// 入库人
	IngressMan        system.SysUser `gorm:"foreignKey:IngressManID"`
	IngressManID      *uint
	BusIngressDetails []BusIngressDetail `gorm:"foreignKey:BusIngressID" json:"bus_ingress_details" form:"bus_ingress_details"`
}

// TableName BusIngress 表名
func (BusIngress) TableName() string {
	return "bus_ingress"
}

type BusEgressDetail struct {
	global.GVA_MODEL
	BusEgressID *uint
	GoodsDict   BusGoodsDict `gorm:"foreignKey:GoodsDictID" json:"goods_dict" form:"goods_dict"`
	GoodsDictID *uint        `json:"goods_dict_id" form:"goods_dict_id" `
	Number      uint         `json:"number" form:"number" ` // 数量
}

// TableName BusEgressDetail 表名
func (BusEgressDetail) TableName() string {
	return "bus_egress_detail"
}

// BusEgress 结构体
type BusEgress struct {
	global.GVA_MODEL
	// 出库人
	EgressMan        system.SysUser `gorm:"foreignKey:EgressManID"`
	EgressManID      *uint
	BusEgressDetails []BusEgressDetail `gorm:"foreignKey:BusEgressID" json:"bus_egress_details" form:"bus_egress_details"`
}

// TableName BusEgress 表名
func (BusEgress) TableName() string {
	return "bus_egress"
}

type BusUseDetail struct {
	global.GVA_MODEL
	BusUseID    *uint
	GoodsDict   BusGoodsDict `gorm:"foreignKey:GoodsDictID" json:"goods_dict" form:"goods_dict"`
	GoodsDictID *uint        `json:"goods_dict_id" form:"goods_dict_id" `
	Number      uint         `json:"number" form:"number" ` // 数量
}

// TableName BusEgressDetail 表名
func (BusUseDetail) TableName() string {
	return "bus_use_detail"
}

// BusUse 结构体
type BusUse struct {
	global.GVA_MODEL
	// 使用耗材的人
	User          system.SysUser `gorm:"foreignKey:UserID"`
	UserID        *uint
	BusUseDetails []BusUseDetail `gorm:"foreignKey:BusUseID" json:"bus_use_details" form:"bus_use_details"`
}

// TableName BusEgressDetail 表名
func (BusUse) TableName() string {
	return "bus_use"
}

// BusGoods 结构体
// 记录了每一个耗材实体
type BusGoods struct {
	ID        *uint          `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间

	// 属于哪个Order
	BusOrder   BusOrder `gorm:"foreignKey:BusOrderID"`
	BusOrderID *uint    `json:"bus_order_id" form:"bus_order_id" `

	// 属于哪个Ingress
	BusIngress   BusIngress `gorm:"foreignKey:BusIngressID"`
	BusIngressID *uint      `json:"bus_ingress_id" form:"bus_ingress_id" `

	// 属于哪个Egress
	BusEgress   BusEgress `gorm:"foreignKey:BusEgressID"`
	BusEgressID *uint     `json:"bus_egress_id" form:"bus_egress_id" `

	GoodsDict   BusGoodsDict `gorm:"foreignKey:GoodsDictID"`
	GoodsDictID *uint        `json:"goods_dict_id" form:"goods_dict_id" `

	SerialNumber   uint      `json:"serial_number" form:"serial_number" `     // 用于打印的序列号
	Batch          string    `json:"batch" form:"batch" `                     // 批号
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
