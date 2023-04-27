package base

type Page struct {
	PageNo   uint        `json:"pageNo"`
	PageSize uint        `json:"pageSize"`
	Total    uint        `json:"total"`
	Data     interface{} `json:"data"`
}
