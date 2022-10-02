package test

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

func Test2(t *testing.T) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	err := f.SetCellValue("Sheet2", "A2", "Hello world.")
	if err != nil {
		return
	}
	err = f.SetCellValue("Sheet1", "B2", 100)
	if err != nil {
		return
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
