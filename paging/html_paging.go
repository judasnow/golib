package paging

import (
	"strconv"
	"fmt"
)

type SimplePaging struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func GetSimplePagingInfo(crtPage int, perPage int) SimplePaging {
	if perPage < 1 {
		perPage = 20
	} else if perPage > 100 {
		perPage = 100
	}

	if crtPage < 1 {
		crtPage = 1
	}

	offset := (crtPage - 1) * (perPage)
	limit := perPage

	return SimplePaging{
		Limit:  limit,
		Offset: offset,
	}
}

type PageItem struct {
	Text   string
	Link   string
	Enable bool
}

func getPrevPage(crtPage int, baseLink string) PageItem {
	prevPage := crtPage - 1
	if prevPage <= 0 {
		prevPage = 1
	}
	var prevPageEnable bool
	if prevPage == 1 {
		prevPageEnable = false
	} else {
		prevPageEnable = true
	}
	return PageItem{
		Text:   "上一页",
		Link:   fmt.Sprintf("%s?page=%d", baseLink, prevPage),
		Enable: prevPageEnable,
	}
}

func getNextPage(totalPage int, crtPage int, baseLink string) PageItem {
	nextPage := crtPage + 1
	if nextPage > totalPage {
		nextPage = totalPage
	}
	var nextPageEnable bool
	if nextPage == totalPage {
		nextPageEnable = false
	} else {
		nextPageEnable = true
	}
	return PageItem{
		Text:   "上一页",
		Link:   fmt.Sprintf("%s?page=%d", baseLink, nextPage),
		Enable: nextPageEnable,
	}
}

func HtmlPaging(totalPage int, crtPage int, baseLink string) []PageItem {
	var pages []PageItem

	if totalPage <= 8 {
		// 全部显示
		for i := 1; i <= 8; i++ {
			var enable bool
			if crtPage == i {
				enable = false
			} else {
				enable = true
			}
			pages = append(pages, PageItem{
				Text:   strconv.Itoa(i),
				Link:   fmt.Sprintf("%s?page=%s", baseLink, i),
				Enable: enable,
			})
		}
	} else {
		if crtPage < 7 {
			if crtPage <= 5 {
				// 显示前 5 个，以及后两个
			} else {
				// 显示前 8 个，以及最后两个
			}
		} else if crtPage > totalPage-7 {
			if crtPage > totalPage-5 {
				// 显示后 5 个，以及前 2 个
			} else {
				// 显示后 8 个，以及前 2 个
			}
		} else {
			// 显示当前页前后两页，以及开头结尾各两页
		}
	}

	pages = append([]PageItem{
		getPrevPage(crtPage, baseLink),
	}, pages...)
	pages = append(pages, getNextPage(totalPage, crtPage, baseLink))

	return pages
}
