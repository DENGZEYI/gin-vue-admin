package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test"
	testReq "github.com/flipped-aurora/gin-vue-admin/server/model/test/request"
)

type TestDicService struct {
}

// CreateTestDic 创建TestDic记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDicService *TestDicService) CreateTestDic(testDic test.TestDic) (err error) {
	err = global.GVA_DB.Create(&testDic).Error
	return err
}

// DeleteTestDic 删除TestDic记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDicService *TestDicService) DeleteTestDic(testDic test.TestDic) (err error) {
	err = global.GVA_DB.Delete(&testDic).Error
	return err
}

// DeleteTestDicByIds 批量删除TestDic记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDicService *TestDicService) DeleteTestDicByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]test.TestDic{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTestDic 更新TestDic记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDicService *TestDicService) UpdateTestDic(testDic test.TestDic) (err error) {
	err = global.GVA_DB.Save(&testDic).Error
	return err
}

// GetTestDic 根据id获取TestDic记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDicService *TestDicService) GetTestDic(id uint) (testDic test.TestDic, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&testDic).Error
	return
}

// GetTestDicInfoList 分页获取TestDic记录
// Author [piexlmax](https://github.com/piexlmax)
func (testDicService *TestDicService) GetTestDicInfoList(info testReq.TestDicSearch) (list []test.TestDic, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&test.TestDic{})
	var testDics []test.TestDic
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&testDics).Error
	return testDics, total, err
}
