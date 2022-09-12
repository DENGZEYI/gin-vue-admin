package reply

import "time"

type BusIngressRep struct {
	BusOrderID    *uint     `json:"ID" gorm:"column:id"`
	GoodsName     string    `json:"goods_name"`
	Specification string    `json:"specification"`
	ApplyNumber   uint      `json:"apply_number"`
	ArrivalNumber uint      `json:"arrival_number"`
	GroupName     string    `json:"group_name"`
	ProviderName  string    `json:"provider_name"`
	ApplyDate     time.Time `json:"apply_date"`
	ApplicantName string    `json:"applicant_name"`
	ApproveDate   time.Time `json:"approve_date"`
	ApproverName  string    `json:"approver_name"`
	PurchaserName string    `json:"purchaser_name"`
	PurchaseDate  time.Time `json:"purchase_date"`
}
