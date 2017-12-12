package golib

import (
	"reflect"
	"testing"
)

func TestStructConv(t *testing.T) {
	type Foo struct {
		Name string
	}
	foo := Foo{Name: "foo"}

	type Bar struct {
		Name string
	}

	StructConv(foo, reflect.TypeOf(Bar{}))
}
