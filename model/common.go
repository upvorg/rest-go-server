package model

type Pagination struct {
	Page  uint64 `json:"page" form:"page,default=1" binding:"min=1"`
	Limit uint64 `json:"limit" form:"limit,default=15" binding:"min=1,max=30"`
	Total uint64 `json:"total"` // for response
}

type PaginationData struct {
	Pagination
	Data interface{} `json:"data"`
}
