package common

type ListRequest struct {
	PageNo         int    `json:"pageNo" validate:"required"`
	PageSize       int    `json:"pageSize" validate:"required, max=100"`
	IsDesc         bool   `json:"isDesc"`
	OrderFieldName string `json:"orderFieldName"`
}
