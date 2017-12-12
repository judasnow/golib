package golib

type PageInfo struct {
	Count  int
	Limit  int
	Offset int
}

func newPageInfo(count int, limit int, offset int) PageInfo {
	return PageInfo{
		Count:  count,
		Limit:  limit,
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
	count := ((total - 1) / perPage) + 1

	if crtPage <= 0 {
		// 默认第一页
		return newPageInfo(count, perPage, 0)
	}

	return newPageInfo(count, perPage, (crtPage-1)*perPage)
}

func PagingLoop(total int, perPage int, crtPage int, f func(limit int, offset int) bool) {
	for {
		pageInfo := Paging(total, perPage, crtPage)

		if crtPage > pageInfo.Count || crtPage <= 0 {
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

// 同样的 返回 false 会停止循环
func SlicePagingLoop(slice []interface{}, perPage int, crtPage int, f func(slicePerPage []interface{}) bool) {
	var total int = len(slice)

	PagingLoop(total, perPage, crtPage, func(limit int, offset int) bool {
		// 防止最后一页不满 limit 的情况，会导致 index 越界错误
		var _limit int

		if total-offset >= limit {
			_limit = limit
		} else {
			_limit = total - offset
		}

		return f(slice[offset : _limit+offset])
	})
}
