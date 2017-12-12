package paging

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
