package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
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
	err = db.Limit(limit).Offset(offset).Find(&busGoodss).Error
	return busGoodss, total, err
}

// ApplyGoodsByIds 创建申请记录
func (busGoodsService *BusGoodsService) ApplyGoodsByIds(busApplyInfo businessReq.BusApplyInfo, c *gin.Context) (err error) {
	applicantID := utils.GetUserID(c) // 获取申请者ID
	length := len(busApplyInfo.Ids)
	GoodsSlice := make([]business.BusGoods, length)
	for i := 0; i < length; i++ {
		good := business.BusGoods{
			ID: uint(busApplyInfo.Ids[i].ID),
		}
		GoodsSlice[i] = good
	}
	order := business.BusOrder{
		Goods:       GoodsSlice,
		ApplicantID: applicantID,
		State:       global.Processing,
	}
	// many2many的自定义连接表也可以完全手动写入，目前是采用GORM写入两个主键并手动写入number和batch
	rst := global.GVA_DB.Create(&order) // 这里只写入了连接表的两个外键，还有number和batch需要手动写入
	if rst.Error != nil {
		return rst.Error
	}
	for i := 0; i < length; i++ {
		busOrderInfo := business.BusOrderGoods{
			BusOrderID: order.ID,
			BusGoodsID: busApplyInfo.Ids[i].ID,
			Number:     busApplyInfo.Ids[i].Number,
		}
		rst = global.GVA_DB.Save(&busOrderInfo)
		if rst.Error != nil {
			return rst.Error
		}
	}

	return err
}
