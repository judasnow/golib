package golib

import (
	"reflect"
	"fmt"
)


func StructConv(sourceStruct interface{}, targetStruct reflect.Type) (interface{}, error) {
	// 将 sourceStruct 转换为 targetStruct 类型，只设置名称相同的 field
	t := reflect.TypeOf(sourceStruct)
	fmt.Printf("%+v", t)

	return nil, nil
}
