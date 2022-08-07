package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
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

// UpdateBusOrder 更新BusOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (busOrderService *BusOrderService) UpdateBusOrder(busOrder business.BusOrder) (err error) {
	err = global.GVA_DB.Save(&busOrder).Error
	return err
}

// GetBusOrder 根据id获取BusOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (busOrderService *BusOrderService) GetBusOrder(id uint) (busOrder business.BusOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&busOrder).Error
	return
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
	err = db.Limit(limit).Offset(offset).Find(&busOrders).Error
	return busOrders, total, err
}
