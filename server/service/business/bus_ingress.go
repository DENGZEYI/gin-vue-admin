package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business/reply"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"gorm.io/gorm/clause"
)

type BusIngressService struct {
}

func (busIngressService *BusIngressService) GetBusIngressInfoList(info request.BusIngressSearch) (repList []reply.BusIngressRep, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	var orderList []business.BusOrder
	// 创建db
	db := global.GVA_DB.Model(&business.BusOrder{})
	orderStr := "id desc" //降序排列
	db.Order(orderStr).Preload("BusOrderDetails.GoodsDict.Provider").Preload("BusOrderDetails.GoodsDict.Group").Preload(clause.Associations).Find(&orderList)
	// 填写返回的数据
	for i := 0; i < len(orderList); i++ {
		for j := 0; j < len(orderList[i].BusOrderDetails); j++ {
			//  根据orderID和goodDictID计算已入库数量
			orderID := orderList[i].ID
			goodsDictID := orderList[i].BusOrderDetails[j].GoodsDictID
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
			var rep = reply.BusIngressRep{
				BusOrderID:    &(orderList[i].ID),
				GoodsName:     orderList[i].BusOrderDetails[j].GoodsDict.Name,
				Specification: orderList[i].BusOrderDetails[j].GoodsDict.Specification,
				ApplyNumber:   orderList[i].BusOrderDetails[j].Number,
				ArrivalNumber: uint(arrivalNum),
				GroupName:     orderList[i].BusOrderDetails[j].GoodsDict.Group.Name,
				ProviderName:  orderList[i].BusOrderDetails[j].GoodsDict.Provider.Name,
				ApplyDate:     orderList[i].CreatedAt,
				ApplicantName: orderList[i].Applicant.NickName,
				ApproveDate:   orderList[i].ApproveTime,
				ApproverName:  orderList[i].Approver.NickName,
				PurchaserName: orderList[i].Purchaser.NickName,
				PurchaseDate:  orderList[i].PurchaseTime,
			}
			repList = append(repList, rep)
		}
	}
	total = int64(len(repList))
	// 对返回的列表进行切片
	endIdx := offset + limit
	if endIdx > len(repList) {
		// 避免切片结尾的索引值大于列表长度
		endIdx = len(repList)
	}
	repList = repList[offset:endIdx]
	return repList, total, err
}
