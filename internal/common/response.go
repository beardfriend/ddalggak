package common

type Response struct {
	Message string `json:"message"`
}

type ResponseWithData struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type ResponseWithPagination struct {
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
	Pagination interface{} `json:"pagination"`
}
