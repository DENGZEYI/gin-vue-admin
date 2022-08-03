package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusProviderService struct {
}

// CreateBusProvider 创建BusProvider记录
// Author [piexlmax](https://github.com/piexlmax)
func (busProviderService *BusProviderService) CreateBusProvider(busProvider business.BusProvider) (err error) {
	err = global.GVA_DB.Create(&busProvider).Error
	return err
}

// DeleteBusProvider 删除BusProvider记录
// Author [piexlmax](https://github.com/piexlmax)
func (busProviderService *BusProviderService) DeleteBusProvider(busProvider business.BusProvider) (err error) {
	err = global.GVA_DB.Delete(&busProvider).Error
	return err
}

// DeleteBusProviderByIds 批量删除BusProvider记录
// Author [piexlmax](https://github.com/piexlmax)
func (busProviderService *BusProviderService) DeleteBusProviderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.BusProvider{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBusProvider 更新BusProvider记录
// Author [piexlmax](https://github.com/piexlmax)
func (busProviderService *BusProviderService) UpdateBusProvider(busProvider business.BusProvider) (err error) {
	err = global.GVA_DB.Save(&busProvider).Error
	return err
}

// GetBusProvider 根据id获取BusProvider记录
// Author [piexlmax](https://github.com/piexlmax)
func (busProviderService *BusProviderService) GetBusProvider(id uint) (busProvider business.BusProvider, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&busProvider).Error
	return
}

// GetBusProviderInfoList 分页获取BusProvider记录
// Author [piexlmax](https://github.com/piexlmax)
func (busProviderService *BusProviderService) GetBusProviderInfoList(info businessReq.BusProviderSearch) (list []business.BusProvider, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.BusProvider{})
	var busProviders []business.BusProvider
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&busProviders).Error
	return busProviders, total, err
}

// GetBusProviderDict 获取Provider字典
func (busProviderService *BusProviderService) GetBusProviderDict() (list []business.BusProvider, err error) {
	// 创建db
	db := global.GVA_DB.Model(&business.BusProvider{})
	var busProviders []business.BusProvider
	err = db.Find(&busProviders).Error
	return busProviders, err
}
