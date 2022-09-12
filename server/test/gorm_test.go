package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business/reply"
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
	//var results []result
	var list []reply.BusIngressRep
	var orderlist []business.BusOrder
	// 创建db
	db = db.Model(&business.BusOrder{})
	db.Preload("BusOrderDetails.GoodsDict.Provider").Preload("BusOrderDetails.GoodsDict.Group").Preload(clause.Associations).Find(&orderlist)
	// 填写返回的数据
	for i := 0; i < len(orderlist); i++ {
		for j := 0; j < len(orderlist[i].BusOrderDetails); j++ {
			//  根据orderID和goodDictID计算已入库数量
			orderID := orderlist[i].ID
			goodsDictID := orderlist[i].BusOrderDetails[j].GoodsDictID
			arrivalNum := 0
			//  根据orderID和goodDictID计算已入库数量
			var ingresses []business.BusIngress
			global.GVA_DB.Where("bus_order_id = ?", orderID).Preload("BusIngressDetails").Find(&ingresses)
			for j := 0; j < len(ingresses); j++ {
				for k := 0; k < len(ingresses[j].BusIngressDetails); k++ {
					if *ingresses[j].BusIngressDetails[k].GoodsDictID == *goodsDictID {
						arrivalNum += int(ingresses[j].BusIngressDetails[k].IngressNumber)
					}
				}
			}
			// 填写返回的数据
			var rep = reply.BusIngressRep{
				BusOrderID:    &(orderlist[i].ID),
				GoodsName:     orderlist[i].BusOrderDetails[j].GoodsDict.Name,
				Specification: orderlist[i].BusOrderDetails[j].GoodsDict.Specification,
				ApplyNumber:   orderlist[i].BusOrderDetails[j].Number,
				ArrivalNumber: uint(arrivalNum),
				GroupName:     orderlist[i].BusOrderDetails[j].GoodsDict.Group.Name,
				ProviderName:  orderlist[i].BusOrderDetails[j].GoodsDict.Provider.Name,
				ApplyDate:     orderlist[i].CreatedAt,
				ApplicantName: orderlist[i].Applicant.NickName,
				ApproveDate:   orderlist[i].ApproveTime,
				ApproverName:  orderlist[i].Approver.NickName,
				PurchaserName: orderlist[i].Purchaser.NickName,
			}
			list = append(list, rep)
		}
	}
	print("输出")
}
