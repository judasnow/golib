package paging

import (
	"testing"
	"fmt"
)

func TestGetSimplePagingInfo(t *testing.T) {
	fmt.Print(GetSimplePagingInfo(1, 20))
}

func TestHtmlPaging(t *testing.T) {
	fmt.Print(HtmlPaging(8, 1, "http://foobar.com"))
}