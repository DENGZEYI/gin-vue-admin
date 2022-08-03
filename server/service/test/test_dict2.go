package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test"
	testReq "github.com/flipped-aurora/gin-vue-admin/server/model/test/request"
)

type TestDict2Service struct {
}

// CreateTestDict2 创建TestDict2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDict2Service *TestDict2Service) CreateTestDict2(testDict2 test.TestDict2) (err error) {
	err = global.GVA_DB.Create(&testDict2).Error
	return err
}

// DeleteTestDict2 删除TestDict2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDict2Service *TestDict2Service) DeleteTestDict2(testDict2 test.TestDict2) (err error) {
	err = global.GVA_DB.Delete(&testDict2).Error
	return err
}

// DeleteTestDict2ByIds 批量删除TestDict2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDict2Service *TestDict2Service) DeleteTestDict2ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]test.TestDict2{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTestDict2 更新TestDict2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDict2Service *TestDict2Service) UpdateTestDict2(testDict2 test.TestDict2) (err error) {
	err = global.GVA_DB.Save(&testDict2).Error
	return err
}

// GetTestDict2 根据id获取TestDict2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDict2Service *TestDict2Service) GetTestDict2(id uint) (testDict2 test.TestDict2, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&testDict2).Error
	return
}

// GetTestDict2InfoList 分页获取TestDict2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDict2Service *TestDict2Service) GetTestDict2InfoList(info testReq.TestDict2Search) (list []test.TestDict2, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&test.TestDict2{})
	var testDict2s []test.TestDict2
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Group != nil {
		// 避免使用sql关键词GROUP，所以需要使用`进行转义
		db = db.Where("`group` = ?", info.Group)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&testDict2s).Error
	return testDict2s, total, err
}
