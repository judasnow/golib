package golib

import (
	"fmt"
	"reflect"
	"strconv"
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

	// 取出第一个元素 反射一次 取出所有 tag，同时计算出 fields 的数量
	firstRow := value.Index(0)
	tags := getXlsxTags(firstRow)

	// 将刚取出的 tags 作为头部写入 xlsx
	xlsxSheet, err := file.AddSheet(sheet.Name)
	if err != nil {
		return
	}
	row := xlsxSheet.AddRow()
	row.WriteSlice(&tags, len(tags))

	// 之后循环所有成员 写入 xlsx 文件
	for i := 0; i < value.Len(); i++ {
		row := xlsxSheet.AddRow()
		for j := 0; j < value.Index(i).NumField(); j++ {
			cell := row.AddCell()
			valueField := value.Index(i).Field(j)

			switch t := valueField.Interface().(type) {
			case int:
				cell.SetString(strconv.Itoa(t))
			case string:
				cell.SetString(t)
			}
		}
	}
}

func getXlsxTags(data reflect.Value) []string {
	// TODO 检测必须是 struct

	tags := []string{}

	for i := 0; i < data.NumField(); i++ {
		tagValue := data.Type().Field(i).Tag.Get(TAGNAME)
		tags = append(tags, tagValue)
	}

	return tags
}