package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusFactoryService struct {
}

// CreateBusFactory 创建BusFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (busFactoryService *BusFactoryService) CreateBusFactory(busFactory business.BusFactory) (err error) {
	err = global.GVA_DB.Create(&busFactory).Error
	return err
}

// DeleteBusFactory 删除BusFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (busFactoryService *BusFactoryService) DeleteBusFactory(busFactory business.BusFactory) (err error) {
	err = global.GVA_DB.Delete(&busFactory).Error
	return err
}

// DeleteBusFactoryByIds 批量删除BusFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (busFactoryService *BusFactoryService) DeleteBusFactoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.BusFactory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBusFactory 更新BusFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (busFactoryService *BusFactoryService) UpdateBusFactory(busFactory business.BusFactory) (err error) {
	err = global.GVA_DB.Save(&busFactory).Error
	return err
}

// GetBusFactory 根据id获取BusFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (busFactoryService *BusFactoryService) GetBusFactory(id uint) (busFactory business.BusFactory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&busFactory).Error
	return
}

// GetBusFactoryInfoList 分页获取BusFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (busFactoryService *BusFactoryService) GetBusFactoryInfoList(info businessReq.BusFactorySearch) (list []business.BusFactory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.BusFactory{})
	var busFactorys []business.BusFactory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&busFactorys).Error
	return busFactorys, total, err
}

// GetBusFactoryDict 获取Factory字典
func (busFactoryService *BusFactoryService) GetBusFactoryDict() (list []business.BusFactory, err error) {
	// 创建db
	db := global.GVA_DB.Model(&business.BusFactory{})
	var busFactories []business.BusFactory
	err = db.Find(&busFactories).Error
	return busFactories, err
}
