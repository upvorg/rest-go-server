package model

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Page  uint64 `json:"page" form:"page,default=1" binding:"min=1"`
	Limit uint64 `json:"limit" form:"limit,default=15" binding:"min=1,max=30"`
	Total uint64 `json:"total"` // for response
}

type PaginationData struct {
	Pagination
	Data interface{} `json:"data"`
}

func Paginate(r *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(r.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
