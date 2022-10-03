package test

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"testing"
)

// CreateXlS data为要写入的数据,fileName 文件名称, headerNameArray 表头数组
func CreateXlS(data [2][2]string, fileName string, headerNameArray [2]string) {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	sheetWords := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z",
	}
	for k, v := range headerNameArray {
		err := f.SetCellValue(sheetName, sheetWords[k]+"1", v)
		if err != nil {
			return
		}
	}

	line := 1
	// 循环写入数据
	for _, v := range data {
		line++
		for kk, _ := range headerNameArray {
			f.SetCellValue(sheetName, sheetWords[kk]+strconv.Itoa(line), v[kk])
		}
	}
	// 保存文件
	if err := f.SaveAs(fileName + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
func Test2(t *testing.T) {
	var stringData = [2][2]string{
		{"1", "2"},
		{"3", "4"},
	}
	var headerName = [2]string{
		"第一列", "第二列",
	}
	CreateXlS(stringData, "test", headerName)
}
