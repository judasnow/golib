package golib

import (
	"fmt"
	"reflect"
	"strings"
	"time"
	"errors"
	"bytes"

	"github.com/tealeg/xlsx"
)

// `xlsx:"name:name;format:2006-01-02"`
const (
	tag_name              = "xlsx"
	tag_spliter           = ";"
	tag_key_value_spliter = ":"
	booltext_spliter = ","
)

// xlsx 文件的一个 sheet
type Sheet struct {
	Name string
	// 类型应该是一个 []struct
	Datas interface{}
	// 额外的数据 [][]interface{}
	ExtraDatas [][]interface{}
}

// column meta datas
type tag struct {
	Name string
	// field 类型为 time.Time 时可指定其格式
	TimeFormat string
	// field 类型为 bool 时，最终表单中的文字描述
	TrueText string
	FalseText string
}

func ExportToXlsx(sheets []Sheet) ([]byte, error){
	file := xlsx.NewFile()

	for _, sheet := range sheets {
		if err := exportToSheet(file, sheet); err != nil {
			return []byte{}, err
		}
	}

	bufferFile := bytes.Buffer{}
	if err := file.Write(&bufferFile); err != nil {
		return []byte{}, err
	} else {
		return bufferFile.Bytes(), nil
	}
}

func exportToSheet(file *xlsx.File, sheet Sheet) error {
	// sheet.Datas kind 必须是一个 array 或者 slice 类型
	datasValue := reflect.ValueOf(sheet.Datas)
	datasKind := datasValue.Kind()
	if datasKind != reflect.Slice && datasKind != reflect.Array {
		return errors.New("Sheet.Datas 类型（Kind）需要是 Slice 或 Array")
	}

	datasElemType := reflect.TypeOf(sheet.Datas).Elem()
	tags := getXlsxTags(datasElemType)

	// 将刚取出的 tags 作为头部写入 xlsx
	xlsxSheet, addSheetErr := file.AddSheet(sheet.Name)
	if addSheetErr != nil {
		return addSheetErr
	}

	// 写入标题
	// 标题由量部分组成 tags 以及 extraDatas 中的第一排
	tagNames := pluckTagName(tags)
	titles := tagNames
	if sheet.ExtraDatas != nil {
		for _, title := range sheet.ExtraDatas[0] {
			titleStr, ok := title.(string)
			if !ok {
				return errors.New("Sheet.ExtraDatas 中的第一行，作为 title，必须可以断言为 string")
			} else {
				titles = append(titles, titleStr)
			}
		}
	}
	row := xlsxSheet.AddRow()
	row.WriteSlice(&titles, len(titles))

	// 之后循环所有成员 写入 xlsx 文件
	for lineNo := 0; lineNo < datasValue.Len(); lineNo++ {
		xlsxRow := xlsxSheet.AddRow()
		row := datasValue.Index(lineNo)

		for cloumnNo := 0; cloumnNo < row.NumField(); cloumnNo++ {
			cell := xlsxRow.AddCell()
			valueField := row.Field(cloumnNo)

			switch v := valueField.Interface().(type) {
			case bool:
				if v {
					cell.SetString(tags[cloumnNo].TrueText)
				} else {
					cell.SetString(tags[cloumnNo].FalseText)
				}
			case time.Time:
				cell.SetString(v.Format(tags[cloumnNo].TimeFormat))
			default:
				cell.SetString(fmt.Sprintf("%v", v))
			}
		}

		// 写入相应的额外数据 额外数据类型只能是 string
		if sheet.ExtraDatas != nil {
			for _, dataRow := range sheet.ExtraDatas[lineNo+1] {
				cell := xlsxRow.AddCell()
				switch v := dataRow.(type) {
				default:
					cell.SetString(fmt.Sprintf("%v", v))
				}
			}
		}

	}

	return nil
}

func getXlsxTags(data reflect.Type) []tag {
	tags := []tag{}

	for i := 0; i < data.NumField(); i++ {
		tagValue := data.Field(i).Tag.Get(tag_name)
		tag := parseTag(tagValue)
		tags = append(tags, tag)
	}

	return tags
}

func parseTag(tagString string) tag {
	tagItems := strings.Split(tagString, tag_spliter)

	tag := tag{}
	for _, tagItem := range tagItems {
		tagItemPair := strings.Split(tagItem, tag_key_value_spliter)

		if tagItemPair[0] == "name" {
			tag.Name = tagItemPair[1]
		} else if tagItemPair[0] == "format" {
			tag.TimeFormat = tagItemPair[1]
		} else if tagItemPair[0] == "booltext" {
			value := tagItemPair[1]
			tag.TrueText, tag.FalseText, _ = parseBooltextTag(value)
		}
	}

	return tag
}

// 解析 tag 中的 booltext
func parseBooltextTag(value string) (string, string, error) {
	booltextArgs := strings.Split(value, booltext_spliter)
	return booltextArgs[0], booltextArgs[1], nil
}

func pluckTagName(tags []tag) []string {
	tagNames := []string{}

	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	return tagNames
}
