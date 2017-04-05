package golib


type PageInfo struct {
	Count int
	Limit  int
	Offset int
}
func newPageInfo(count int, limit int, offset int) PageInfo {
	return PageInfo{
		Count: count,
		Limit: limit,
		Offset: offset,
	}
}
func Paging(total int, perPage int, crtPage int) PageInfo {
	if total <= 0 {
		return newPageInfo(0, perPage, 0)
	}
	if perPage <= 0 {
		return newPageInfo(0, perPage, 0)
	}
	// total > 0, perPage > 0
	var count int = ((total - 1) / perPage) + 1

	if crtPage <= 0 {
		// 默认第一页
		return newPageInfo(count, perPage, 0)
	}

	return newPageInfo(count, perPage, (crtPage - 1) * perPage)
}

func PagingLoop(total int, perPage int, crtPage int, f func(limit int, offset int) bool) {
	for {
		pageInfo := Paging(total, perPage, crtPage)

		if (crtPage > pageInfo.Count || crtPage <= 0) {
			return
		} else {
			// loop content
			if !f(pageInfo.Limit, pageInfo.Offset) {
				return
			}
			crtPage = crtPage + 1
		}

	}
}
