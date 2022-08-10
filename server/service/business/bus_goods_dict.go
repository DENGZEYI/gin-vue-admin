package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type BusGoodsDictService struct {
}

// CreateBusGoodsDict 创建BusGoodsDict记录
// Author [piexlmax](https://github.com/piexlmax)
func (BusGoodsDictService *BusGoodsDictService) CreateBusGoodsDict(BusGoodsDict business.BusGoodsDict) (err error) {
	err = global.GVA_DB.Create(&BusGoodsDict).Error
	return err
}

// DeleteBusGoodsDict 删除BusGoodsDict记录
// Author [piexlmax](https://github.com/piexlmax)
func (BusGoodsDictService *BusGoodsDictService) DeleteBusGoodsDict(BusGoodsDict business.BusGoodsDict) (err error) {
	err = global.GVA_DB.Delete(&BusGoodsDict).Error
	return err
}

// DeleteBusGoodsDictByIds 批量删除BusGoodsDict记录
// Author [piexlmax](https://github.com/piexlmax)
func (BusGoodsDictService *BusGoodsDictService) DeleteBusGoodsDictByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.BusGoodsDict{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBusGoodsDict 更新BusGoodsDict记录
// Author [piexlmax](https://github.com/piexlmax)
func (BusGoodsDictService *BusGoodsDictService) UpdateBusGoodsDict(BusGoodsDict business.BusGoodsDict) (err error) {
	err = global.GVA_DB.Save(&BusGoodsDict).Error
	return err
}

// GetBusGoodsDict 根据id获取BusGoodsDict记录
// Author [piexlmax](https://github.com/piexlmax)
func (BusGoodsDictService *BusGoodsDictService) GetBusGoodsDict(id uint) (BusGoodsDict business.BusGoodsDict, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&BusGoodsDict).Error
	return
}

// GetBusGoodsDictInfoList 分页获取BusGoodsDict记录
// Author [piexlmax](https://github.com/piexlmax)
func (BusGoodsDictService *BusGoodsDictService) GetBusGoodsDictInfoList(info businessReq.BusGoodsDictSearch) (list []business.BusGoodsDict, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.BusGoodsDict{})
	var BusGoodsDicts []business.BusGoodsDict
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.GroupID != nil {
		// 避免使用sql关键词GROUP，所以需要使用`进行转义
		db = db.Where("`group_id` = ?", info.GroupID)
	}
	if info.ProviderID != nil {
		db = db.Where("`provider_id` = ?", info.ProviderID)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&BusGoodsDicts).Error
	return BusGoodsDicts, total, err
}

// ApplyGoodsByIds 创建申请记录
func (BusGoodsDictService *BusGoodsDictService) ApplyGoodsByIds(busApplyInfo businessReq.BusApplyInfo, c *gin.Context) (err error) {
	applicantID := utils.GetUserID(c) // 获取申请者ID
	length := len(busApplyInfo.Ids)
	var GoodsSlice []business.BusGoods
	for i := 0; i < length; i++ {
		for j := 0; j < busApplyInfo.Ids[i].Number; j++ {
			good := business.BusGoods{
				GoodsDictID: busApplyInfo.Ids[i].ID,
			}
			GoodsSlice = append(GoodsSlice, good)
		}
	}
	order := business.BusOrder{
		Goods:       GoodsSlice,
		ApplicantID: applicantID,
		State:       global.Processing,
	}
	err = global.GVA_DB.Create(&order).Error
	return err
}
