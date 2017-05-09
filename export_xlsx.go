package golib

import (
	"fmt"
	"reflect"
	//"strconv"
	"github.com/tealeg/xlsx"
	"strings"
	"time"
)

const (
	TAGNAME = "xlsx"
	TAG_SPLITER = ";"
	TAG_KEYSPLITER = ":"
)

type Sheet struct {
	Name string
	Datas interface{}
}

type Tag struct {
	Name string
	TimeFormat string
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
	tagNames := pluckTagName(tags)
	row.WriteSlice(&tagNames, len(tags))

	// 之后循环所有成员 写入 xlsx 文件
	for lineNo := 0; lineNo < value.Len(); lineNo++ {
		xlsxRow := xlsxSheet.AddRow()
		row := value.Index(lineNo)

		for cloumnNo := 0; cloumnNo < row.NumField(); cloumnNo++ {
			cell := xlsxRow.AddCell()
			valueField := row.Field(cloumnNo)

			switch v := valueField.Interface().(type) {
			case time.Time:
				cell.SetString(v.Format(tags[cloumnNo].TimeFormat))
			default:
				cell.SetString(fmt.Sprintf("%v", v))
			}
		}
	}
}

func getXlsxTags(data reflect.Value) []Tag {
	tags := []Tag{}

	for i := 0; i < data.NumField(); i++ {
		tagValue := data.Type().Field(i).Tag.Get(TAGNAME)
		tag := parseTag(tagValue)
		tags = append(tags, tag)
	}

	return tags
}

func parseTag(tagString string) Tag {
	tagInfoArray := strings.Split(tagString, TAG_SPLITER)

	tag := Tag{}
	for _, _tag := range tagInfoArray {
		_tagInfoArray := strings.Split(_tag, TAG_KEYSPLITER)
		if _tagInfoArray[0] == "name" {
			tag.Name = _tagInfoArray[1]
		} else if _tagInfoArray[0] == "format" {
			tag.TimeFormat = _tagInfoArray[1]
		}
	}

	return tag
}

func pluckTagName(tags []Tag) []string {
	tagNames := []string{}
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	return tagNames
}
