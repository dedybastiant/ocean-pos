package dto

type CommonResponse struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Pagination  *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}
