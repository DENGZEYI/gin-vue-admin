package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusGroupService struct {
}

// CreateBusGroup 创建BusGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGroupService *BusGroupService) CreateBusGroup(busGroup business.BusGroup) (err error) {
	err = global.GVA_DB.Create(&busGroup).Error
	return err
}

// DeleteBusGroup 删除BusGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGroupService *BusGroupService) DeleteBusGroup(busGroup business.BusGroup) (err error) {
	err = global.GVA_DB.Delete(&busGroup).Error
	return err
}

// DeleteBusGroupByIds 批量删除BusGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGroupService *BusGroupService) DeleteBusGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.BusGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBusGroup 更新BusGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGroupService *BusGroupService) UpdateBusGroup(busGroup business.BusGroup) (err error) {
	err = global.GVA_DB.Save(&busGroup).Error
	return err
}

// GetBusGroup 根据id获取BusGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGroupService *BusGroupService) GetBusGroup(id uint) (busGroup business.BusGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&busGroup).Error
	return
}

// GetBusGroupInfoList 分页获取BusGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGroupService *BusGroupService) GetBusGroupInfoList(info businessReq.BusGroupSearch) (list []business.BusGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.BusGroup{})
	var busGroups []business.BusGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&busGroups).Error
	return busGroups, total, err
}

// GetBusGroupDict 获取Group字典
func (busGroupService *BusGroupService) GetBusGroupDict() (list []business.BusGroup, err error) {
	// 创建db
	db := global.GVA_DB.Model(&business.BusGroup{})
	var busGroups []business.BusGroup
	err = db.Find(&busGroups).Error
	return busGroups, err
}
