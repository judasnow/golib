package golib

import (
	"testing"
	"time"
)

func TestExportToXlsx(t *testing.T) {
	type S1 struct {
		Foo string `xlsx:"name:foo"`
		Baz string `xlsx:"name:baz"`
		Bar int `xlsx:"name:bar"`
	}

	type S2 struct {
		Foo string `xlsx:"name:foo"`
		Baz string `xlsx:"name:baz"`
		Time time.Time `xlsx:"name:time;format:2006-01-02"`
	}

	sheets := []Sheet{
		{"SHEET1", []S1{{"1", "2", 3}}},
		{"SHEET2", []S2{{"1", "2", time.Now()}}},
	}

	ExportToXlsx(sheets)
}
