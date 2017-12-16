package paging

import (
	"testing"
	"fmt"
)

func TestGetSimplePagingInfo(t *testing.T) {
	fmt.Print(GetSimplePagingInfo(1, 20))
}

func TestHtmlPaging(t *testing.T) {
	for _, item := range HtmlPaging(100, 1, 10, "http://foobar.com?page=1") {
		fmt.Println(item)
	}
}

func TestGetTotalPage(t *testing.T) {
	fmt.Print(GetTotalPage(10,6))
}
