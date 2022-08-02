package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusGoodsService struct {
}

// CreateBusGoods 创建BusGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGoodsService *BusGoodsService) CreateBusGoods(busGoods business.BusGoods) (err error) {
	err = global.GVA_DB.Create(&busGoods).Error
	return err
}

// DeleteBusGoods 删除BusGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGoodsService *BusGoodsService) DeleteBusGoods(busGoods business.BusGoods) (err error) {
	err = global.GVA_DB.Delete(&busGoods).Error
	return err
}

// DeleteBusGoodsByIds 批量删除BusGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGoodsService *BusGoodsService) DeleteBusGoodsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.BusGoods{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBusGoods 更新BusGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGoodsService *BusGoodsService) UpdateBusGoods(busGoods business.BusGoods) (err error) {
	err = global.GVA_DB.Save(&busGoods).Error
	return err
}

// GetBusGoods 根据id获取BusGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGoodsService *BusGoodsService) GetBusGoods(id uint) (busGoods business.BusGoods, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&busGoods).Error
	return
}

// GetBusGoodsInfoList 分页获取BusGoods记录
// Author [piexlmax](https://github.com/piexlmax)
func (busGoodsService *BusGoodsService) GetBusGoodsInfoList(info businessReq.BusGoodsSearch) (list []business.BusGoods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.BusGoods{})
	var busGoodss []business.BusGoods
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&busGoodss).Error
	return busGoodss, total, err
}
