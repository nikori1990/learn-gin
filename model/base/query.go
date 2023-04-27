package base

type PageQuery struct {
	PageNo   uint `json:"pageNo" form:"pageNo"`
	PageSize uint `json:"pageSize" form:"pageSize"`
	Offset   uint `json:"offset"`
	Limit    uint `json:"limit"`
}

func (q *PageQuery) SetFromLimit() *PageQuery {
	if q.PageSize == 0 {
		q.PageSize = 10
	}
	q.Limit = q.PageSize
	q.Offset = (q.PageNo - 1) * q.PageSize
	return q
}
