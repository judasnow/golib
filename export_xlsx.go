package golib

import (
	"fmt"
	"reflect"
	//"strconv"
	"github.com/tealeg/xlsx"
)

const TAGNAME = "xlsx"

type Sheet struct {
	Name string
	Datas interface{}
}

func ExportToXlsx(sheets []Sheet) {
	file := xlsx.NewFile()

	for _, sheet := range sheets {
		exportToSheet(file, sheet)
	}

	if err := file.Save("export_xlsx.xlsx"); err != nil {
		fmt.Printf(err.Error())
		return
	}
}

func exportToSheet(file *xlsx.File, sheet Sheet) {
	value := reflect.ValueOf(sheet.Datas)
	kind := value.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return
	}

	// 取出第一个元素 反射一次 取出所有 tag
	firstRow := value.Index(0)
	tags := getXlsxTags(firstRow)

	// 将刚取出的 tags 作为头部写入 xlsx
	xlsxSheet, addSheetErr := file.AddSheet(sheet.Name);
	if addSheetErr != nil {
		return
	}

	row := xlsxSheet.AddRow()
	row.WriteSlice(&tags, len(tags))

	// 之后循环所有成员 写入 xlsx 文件
	for lineNo := 0; lineNo < value.Len(); lineNo++ {
		xlsxRow := xlsxSheet.AddRow()
		row := value.Index(lineNo)

		for cloumnNo := 0; cloumnNo < row.NumField(); cloumnNo++ {
			cell := xlsxRow.AddCell()
			valueField := row.Field(cloumnNo)
			cell.SetString(fmt.Sprintf("%v", valueField.Interface()))
		}
	}
}

func getXlsxTags(data reflect.Value) []string {
	tags := []string{}

	for i := 0; i < data.NumField(); i++ {
		tagValue := data.Type().Field(i).Tag.Get(TAGNAME)
		tags = append(tags, tagValue)
	}

	return tags
}