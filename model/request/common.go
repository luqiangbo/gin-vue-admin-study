package request

type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type GetById struct {
	ID float64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids"`
}
