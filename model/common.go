package model

type Pagination struct {
	Page  uint64 `json:"page,omitempty" form:"page" binding:"required"`
	Limit uint64 `json:"limit,omitempty" form:"limit" binding:"required"`
	Total uint64 `json:"total,omitempty" form:"total"`
}
