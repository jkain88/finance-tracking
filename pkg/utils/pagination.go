package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaginationResult struct {
	HasNext     bool        `json:"has_next"`
	HasPrev     bool        `json:"has_prev"`
	NextPage    int         `json:"next_page"`
	PrevPage    int         `json:"prev_page"`
	TotalPages  int         `json:"total_pages"`
	PageSize    int         `json:"page_size"`
	CurrentPage int         `json:"current_page"`
	TotalCount  int64       `json:"total_count"`
	Rows        interface{} `json:"rows"`
}

func Paginate(c *gin.Context, db *gorm.DB, model interface{}, pagination *PaginationResult) func(db *gorm.DB) *gorm.DB {
	var count int64
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	db.Model(model).Count(&count)

	totalPages := int(math.Ceil(float64(count) / float64(pageSize)))
	hasNext := page < totalPages
	hasPrev := page > 1
	nextPage := page + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 1
	}

	pagination.HasNext = hasNext
	pagination.HasPrev = hasPrev
	pagination.NextPage = nextPage
	pagination.PrevPage = prevPage
	pagination.PageSize = pageSize
	pagination.CurrentPage = page
	pagination.TotalCount = count
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
