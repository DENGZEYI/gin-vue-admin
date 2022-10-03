package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessRep "github.com/flipped-aurora/gin-vue-admin/server/model/business/reply"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

type BusOrderService struct {
}

// CreateBusOrder 创建BusOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (busOrderService *BusOrderService) CreateBusOrder(busOrder business.BusOrder) (err error) {
	err = global.GVA_DB.Create(&busOrder).Error
	return err
}

// DeleteBusOrder 删除BusOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (busOrderService *BusOrderService) DeleteBusOrder(busOrder business.BusOrder) (err error) {
	err = global.GVA_DB.Delete(&busOrder).Error
	return err
}

// DeleteBusOrderByIds 批量删除BusOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (busOrderService *BusOrderService) DeleteBusOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.BusOrder{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBusOrder 审批订购单
func (busOrderService *BusOrderService) UpdateBusOrder(busOrder business.BusOrder, c *gin.Context) (err error) {
	userID := utils.GetUserID(c) // 获取审批者ID
	busOrder.ApproverID = &userID
	busOrder.ApproveTime = time.Now()
	err = global.GVA_DB.Updates(&busOrder).Error
	return err
}

// PurchaseBusOrder 采购订购单
func (busOrderService *BusOrderService) PurchaseBusOrder(busOrder business.BusOrder, c *gin.Context) (err error) {
	userID := utils.GetUserID(c) // 获取审批者ID
	busOrder.PurchaserID = &userID
	busOrder.State = global.Purchasing
	busOrder.PurchaseTime = time.Now()
	err = global.GVA_DB.Updates(&busOrder).Error
	return err
}

// PrintBusOrder 打印订购单
func (busOrderService *BusOrderService) PrintBusOrder(busOrder business.BusOrder) (err error, f *excelize.File) {
	db := global.GVA_DB.Model(&business.BusOrder{})
	err = db.Preload("BusOrderDetails.GoodsDict").First(&busOrder).Error
	err, f = generateXls(busOrder)
	if err != nil {
		return err, f
	}
	// 生成订购单
	return err, f
}

func generateXls(busOrder business.BusOrder) (err error, f *excelize.File) {
	// 获取数据
	rowLen := len(busOrder.BusOrderDetails)
	headerNameArray := []string{"医用耗材名称", "规格型号", "单位", "申购数量"}
	columnLen := len(headerNameArray)
	var data [][]string
	for i := 0; i < rowLen; i++ {
		line := make([]string, columnLen)
		line[0] = busOrder.BusOrderDetails[i].GoodsDict.Name
		line[1] = busOrder.BusOrderDetails[i].GoodsDict.Specification
		line[2] = busOrder.BusOrderDetails[i].GoodsDict.Unit
		line[3] = strconv.Itoa(int(busOrder.BusOrderDetails[i].Number))
		data = append(data, line)
	}
	// 写如excel
	f = excelize.NewFile()
	defaultSheetName := "Sheet1"
	columnWords := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z",
	}
	for k, v := range headerNameArray {
		err := f.SetCellValue(defaultSheetName, columnWords[k]+"1", v)
		if err != nil {
			return err, f
		}
	}
	line := 1
	// 循环写入数据
	for _, v := range data {
		line++
		for kk, _ := range headerNameArray {
			err := f.SetCellValue(defaultSheetName, columnWords[kk]+strconv.Itoa(line), v[kk])
			if err != nil {
				return err, f
			}
		}
	}
	return err, f
}

// GetBusOrder 根据id获取BusOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (busOrderService *BusOrderService) GetBusOrder(id uint) (busOrder business.BusOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&busOrder).Error
	return
}

// GetBusOrderDetails 根据id获取GetBusOrderDetails
// Author [piexlmax](https://github.com/piexlmax)
func (busOrderService *BusOrderService) GetBusOrderDetails(orderID uint) (BusOrderDetailsRep businessRep.BusOrderDetailsRep, err error) {
	var busOrder business.BusOrder
	db := global.GVA_DB.Model(&business.BusOrder{})
	err = db.Where("id = ?", orderID).Preload("BusOrderDetails.GoodsDict.Group").First(&busOrder).Error

	for i := 0; i < len(busOrder.BusOrderDetails); i++ {
		reBusOrderDetail := businessRep.BusOrderDetail{
			GoodsDict:       busOrder.BusOrderDetails[i].GoodsDict,
			Number:          busOrder.BusOrderDetails[i].Number,
			NumberAlreadyIn: 0,
		}
		//  根据orderID和goodDictID计算已入库数量
		goodDictID := busOrder.BusOrderDetails[i].GoodsDictID
		var ingresses []business.BusIngress
		global.GVA_DB.Where("bus_order_id = ?", orderID).Preload("BusIngressDetails").Find(&ingresses)
		for j := 0; j < len(ingresses); j++ {
			for k := 0; k < len(ingresses[j].BusIngressDetails); k++ {
				if *ingresses[j].BusIngressDetails[k].GoodsDictID == *goodDictID {
					reBusOrderDetail.NumberAlreadyIn += ingresses[j].BusIngressDetails[k].IngressNumber
				}
			}
		}
		BusOrderDetailsRep.BusOrderDetails = append(BusOrderDetailsRep.BusOrderDetails, reBusOrderDetail)
	}

	return BusOrderDetailsRep, err
}

// GetBusOrderDetail 根据dict_id和bus_order_id获取BusOrderDetail
func (busOrderService *BusOrderService) GetBusOrderDetail(detail *business.BusOrderDetail) (err error) {
	db := global.GVA_DB.Model(&business.BusOrderDetail{})
	err = db.Where("bus_order_id = ? and goods_dict_id = ?", detail.BusOrderID, detail.GoodsDictID).Preload("GoodsDict").First(detail).Error
	return err
}

// GetBusOrderInfoList 分页获取BusOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (busOrderService *BusOrderService) GetBusOrderInfoList(info businessReq.BusOrderSearch) (list []business.BusOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.BusOrder{})
	var busOrders []business.BusOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	orderStr := "id desc"
	err = db.Limit(limit).Offset(offset).Order(orderStr).Preload("Applicant").Preload("Approver").Preload("Purchaser").Find(&busOrders).Error
	return busOrders, total, err
}

// GetBusStateDict 获取申请表中的State字典
func (busOrderService *BusOrderService) GetBusStateDict() (list []global.ApplyState, err error) {
	var applyStateDict []global.ApplyState
	applyStateDict = append(applyStateDict, global.ApplyState{
		ID:   global.Processing,
		Name: "已申请",
	})
	applyStateDict = append(applyStateDict, global.ApplyState{
		ID:   global.Pass,
		Name: "审批通过",
	})
	applyStateDict = append(applyStateDict, global.ApplyState{
		ID:   global.Fail,
		Name: "审批退回",
	})
	applyStateDict = append(applyStateDict, global.ApplyState{
		ID:   global.Purchasing,
		Name: "已采购",
	})
	return applyStateDict, nil
}
