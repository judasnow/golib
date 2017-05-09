package golib

import (
	"testing"
)

func TestExportToXlsx(t *testing.T) {
	type S1 struct {
		Foo string `xlsx:"foo"`
		Baz string `xlsx:"baz"`
		Bar int `xlsx:"bar"`
	}

	type S2 struct {
		Foo string `xlsx:"foo"`
		Baz string `xlsx:"baz"`
	}

	sheets := []Sheet{
		{"SHEET1", []S1{{"1", "2", 3}}},
		{"SHEET2", []S2{{"1", "2"}}},
	}

	ExportToXlsx(sheets)
}
