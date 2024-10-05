package pagination

type Pagination struct {
	offset   int
	limit    int
	total    int
	pageSize int
	pageNo   int
}

func NewPagination(pageNo, pageSize int) *Pagination {
	return &Pagination{
		pageSize: pageSize,
		pageNo:   pageNo,
		limit:    pageSize,
		offset:   (pageNo - 1) * pageSize,
	}
}

func (p *Pagination) SetTotal(total int) {
	p.total = total
}

func (p *Pagination) GetLimit() int {
	return p.limit
}

func (p *Pagination) GetOffset() int {
	return p.offset
}

type PaginationInfo struct {
	PageSize  int `json:"pageSize"`
	PageNo    int `json:"pageNo"`
	Total     int `json:"total"`
	PageCount int `json:"pageCount"`
	RowCount  int `json:"rowCount"`
}

func (p *Pagination) GetInfo(resultLength int) *PaginationInfo {
	pageCount := p.total/p.pageSize + 1
	if p.total%p.pageSize == 0 {
		pageCount = p.total / p.pageSize
	}

	return &PaginationInfo{
		PageSize:  p.pageSize,
		PageNo:    p.pageNo,
		Total:     p.total,
		PageCount: pageCount,
		RowCount:  resultLength,
	}
}
