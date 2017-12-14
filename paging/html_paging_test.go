package paging

import (
	"testing"
	"fmt"
)

func TestGetSimplePagingInfo(t *testing.T) {
	fmt.Print(GetSimplePagingInfo(1, 20))
}

func TestHtmlPaging(t *testing.T) {
	for _, item := range HtmlPaging(15, 10, "http://foobar.com") {
		fmt.Println(item)
	}
}
