package golib

import (
	"reflect"
)

// 遍历递归结结构的 json
// 暂时只有两种情况
// value 类型为 slice 需要
// 类型为 map 需要遍历每一个 solt
func TraversalJson(data interface{}, f func(data interface{})) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		f(data)
		TraversalMap(data.(map[string]interface{}), f)
	case reflect.Slice:
		f(data)
		TraversalSlice(data.([]interface{}), f)
	case reflect.String:
		f(data.(string))
	}
}

func TraversalMap(m map[string]interface{}, f func(data interface{})) {
	for _, v := range m {
		TraversalJson(v, f)
	}
}

func TraversalSlice(s []interface{}, f func(data interface{})) {
	for _, item := range s {
		TraversalJson(item, f)
	}
}
