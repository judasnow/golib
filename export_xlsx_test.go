package golib

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestExportToXlsx(t *testing.T) {
	type S1 struct {
		Foo     string  `xlsx:"name:foo"`
		Baz     string  `xlsx:"name:baz"`
		Bar     int     `xlsx:"name:bar"`
		FooBar  float64 `xlsx:"name:foobar"`
		FooBool bool    `xlsx:"name:foobool;booltext:是的,不是;"`
	}

	type S2 struct {
		Foo  string    `xlsx:"name:foo"`
		Baz  string    `xlsx:"name:baz"`
		Time time.Time `xlsx:"name:time;format:2006-01-02"`
	}

	sheets := []Sheet{
		{
			"SHEET1",
			[]S1{
				{"1", "2", 3, 3.14, false},
				{"1", "2", 4, 14, true},
			},
			[][]interface{}{
				{"2017-05-01", "2017-05-06"},
				{1, 2},
				{3, 4},
			},
		},
		{
			"第二个sheet",
			[]S2{},
			nil,
		},
	}

	xlsxFileContent, _ := ExportToXlsx(sheets)
	ioutil.WriteFile("export_xlsx.xlsx", xlsxFileContent, 0644)
}
