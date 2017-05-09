package golib

import (
	"fmt"
	"reflect"
	"strings"
	"time"
	"bytes"
	"bufio"
	"errors"

	"github.com/tealeg/xlsx"
)

// `xlsx:"name:name;format:2006-01-02"`
const (
	TAG_NAME = "xlsx"
	TAG_SPLITER = ";"
	TAG_KEY_VALUE_SPLITER = ":"
)

type Sheet struct {
	Name string
	// 应该是一个 []struct
	Datas interface{}
}

type Tag struct {
	Name string
	// field 类型为 time.Time 时可指定其格式
	TimeFormat string
}

func ExportToXlsx(sheets []Sheet) (bytes.Buffer, error){
	file := xlsx.NewFile()

	for _, sheet := range sheets {
		exportToSheet(file, sheet)
	}

	bufferFile := bytes.Buffer{}
	fileWrite := bufio.NewWriter(&bufferFile)

	if err := file.Write(fileWrite); err != nil {
		return bytes.Buffer{}, err
	} else {
		return bufferFile, nil
	}
}

func exportToSheet(file *xlsx.File, sheet Sheet) error {
	value := reflect.ValueOf(sheet.Datas)
	kind := value.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return errors.New("Sheet.Datas must be Slice or Array")
	}

	if value.Len() <= 0 {
		return errors.New("Sheet.Datas lenght <= 0")
	}

	// 取出第一个元素 取出其所有 tag
	firstRow := value.Index(0)
	tags := getXlsxTags(firstRow)

	// 将刚取出的 tags 作为头部写入 xlsx
	xlsxSheet, addSheetErr := file.AddSheet(sheet.Name);
	if addSheetErr != nil {
		return addSheetErr
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

	return nil
}

func getXlsxTags(data reflect.Value) []Tag {
	tags := []Tag{}

	for i := 0; i < data.NumField(); i++ {
		tagValue := data.Type().Field(i).Tag.Get(TAG_NAME)
		tag := parseTag(tagValue)
		tags = append(tags, tag)
	}

	return tags
}

func parseTag(tagString string) Tag {
	// TODO tag 数据格式的验证
	tagInfoArray := strings.Split(tagString, TAG_SPLITER)

	tag := Tag{}
	for _, _tag := range tagInfoArray {
		_tagInfoArray := strings.Split(_tag, TAG_KEY_VALUE_SPLITER)
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
