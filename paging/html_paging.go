package paging

import (
	"strconv"
	"fmt"
	"strings"
	"net/url"
	"github.com/Sirupsen/logrus"
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
		prevPage = 0
	}
	var prevPageEnable bool
	if prevPage == 0 {
		prevPageEnable = false
	} else {
		prevPageEnable = true
	}

	link := createLink(baseLink, prevPage)
	return PageItem{
		Text:   "上一页",
		Link:   link,
		Enable: prevPageEnable,
	}

}

func getNextPage(totalPage int, crtPage int, baseLink string) PageItem {
	nextPage := crtPage + 1
	if nextPage > totalPage + 1 {
		nextPage = totalPage
	}
	var nextPageEnable bool
	if nextPage == totalPage + 1{
		nextPageEnable = false
	} else {
		nextPageEnable = true
	}

	link := createLink(baseLink, nextPage)
	return PageItem{
		Text:   "下一页",
		Link:   link,
		Enable: nextPageEnable,
	}
}

func createLink(baseLink string, page int) string {
	// 删除可能已经存在的 page 参数
	url, pareUrlErr := url.Parse(baseLink)
	if pareUrlErr != nil {
		logrus.WithError(pareUrlErr).Error("解析 url 失败")
		return ""
	}
	v := url.Query()
	v.Del("page")
	url.RawQuery = v.Encode()

	var link string
	if strings.Contains(baseLink, "?") {
		link = fmt.Sprintf("%s&page=%d", url.String(), page)
	} else {
		link = fmt.Sprintf("%s?page=%d", url.String(), page)
	}
	return link
}

func createPages(begin int, end int, crtPage int, baseLink string) []PageItem {
	var pages []PageItem
	for i := begin; i <= end; i++ {
		var enable bool
		if crtPage == i {
			enable = false
		} else {
			enable = true
		}

		link := createLink(baseLink, i)
		pages = append(pages, PageItem{
			Text:   strconv.Itoa(i),
			Link:   link,
			Enable: enable,
		})
	}
	return pages
}

func GetTotalPage(count int, perPage int) int {
	return ((count - 1) / perPage) + 1
}

func HtmlPaging(count int, perPage int, crtPage int, baseLink string) []PageItem {
	if crtPage <= 0 {
		crtPage = 1
	}

	totalPage := GetTotalPage(count, perPage)
	var pages []PageItem

	if totalPage <= 8 {
		// 全部显示
		pages = createPages(1, totalPage, crtPage, baseLink)
	} else {
		if crtPage < 7 {
			if crtPage <= 5 {
				// 显示前 5 个，以及后两个
				pages = createPages(1, 5, crtPage, baseLink)
				pages = append(pages, PageItem{
					Text:   "...",
					Enable: false,
				})
				pages = append(pages, PageItem{
					Text: fmt.Sprintf("%d", totalPage-1),
					Link: createLink(baseLink, totalPage-1),
				})
				pages = append(pages, PageItem{
					Text: fmt.Sprintf("%d", totalPage),
					Link: createLink(baseLink, totalPage),
				})
			} else {
				// 显示前 8 个，以及最后两个
				pages = createPages(1, 8, crtPage, baseLink)
				pages = append(pages, PageItem{
					Text:   "...",
					Enable: false,
				})
				pages = append(pages, PageItem{
					Text: fmt.Sprintf("%d", totalPage-1),
					Link: createLink(baseLink, totalPage-1),
				})
				pages = append(pages, PageItem{
					Text: fmt.Sprintf("%d", totalPage),
					Link: createLink(baseLink, totalPage),
				})
			}
		} else if crtPage > totalPage-7 {
			if crtPage > totalPage-5 {
				// 显示后 5 个，以及前 2 个
				pages = createPages(1, 2, crtPage, baseLink)
				pages = append(pages, PageItem{
					Text:   "...",
					Enable: false,
				})
				pagesTail := createPages(totalPage-4, totalPage, crtPage, baseLink)
				pages = append(pages, pagesTail...)
			} else {
				// 显示后 8 个，以及前 2 个
				pages = createPages(1, 2, crtPage, baseLink)
				pages = append(pages, PageItem{
					Text:   "...",
					Enable: false,
				})
				pagesTail := createPages(totalPage-7, totalPage, crtPage, baseLink)
				pages = append(pages, pagesTail...)
			}
		} else {
			// 显示当前页前后两页，以及开头结尾各两页
			pages = createPages(1, 2, crtPage, baseLink)
			pages = append(pages, PageItem{
				Text:   "...",
				Enable: false,
			})
			pagesMiddle := createPages(crtPage-2, crtPage+2, crtPage, baseLink)
			pages = append(pages, pagesMiddle...)
			pages = append(pages, PageItem{
				Text:   "...",
				Enable: false,
			})
			pagesTail := createPages(totalPage-1, totalPage, crtPage, baseLink)
			pages = append(pages, pagesTail...)
		}
	}

	pages = append([]PageItem{
		getPrevPage(crtPage, baseLink),
	}, pages...)
	pages = append(pages, getNextPage(totalPage, crtPage, baseLink))

	return pages
}
