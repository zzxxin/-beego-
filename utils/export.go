package utils

import (
	"github.com/xuri/excelize/v2"
	"net/http"
	_ "time"
)

// ExportExcel 导出 Excel 文件，传入表头和数据
func ExportExcel(sheetName string, headers []string, data [][]interface{}) (*excelize.File, error) {
	// 创建新的 Excel 文件
	f := excelize.NewFile()

	// 创建一个工作表
	index, _ := f.NewSheet(sheetName)

	// 设置表头
	for i, header := range headers {
		col := string(rune('A' + i))
		cell := col + "1"
		f.SetCellValue(sheetName, cell, header)
	}

	// 填充数据
	for rowIndex, rowData := range data {
		for colIndex, value := range rowData {
			col := string(rune('A' + colIndex))
			cell := col + string(rune('2'+rowIndex))
			f.SetCellValue(sheetName, cell, value)
		}
	}

	// 设置工作表为活动工作表
	f.SetActiveSheet(index)

	return f, nil
}

// SaveExcelToFile 保存 Excel 文件到本地
func SaveExcelToFile(f *excelize.File, filePath string) error {
	return f.SaveAs(filePath)
}

// SaveExcelToResponse 将 Excel 文件写入 HTTP 响应
func SaveExcelToResponse(f *excelize.File, filename string, writer http.ResponseWriter) error {
	// 设置文件头信息
	writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	writer.Header().Set("Content-Disposition", "attachment; filename="+filename)
	writer.Header().Set("Content-Transfer-Encoding", "binary")
	writer.Header().Set("Expires", "0")

	// 写入文件内容
	return f.Write(writer)
}
