package helpers

import (
	"jwt-authen/dtos"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationRequest(c *gin.Context) *dtos.Pagination {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	sort := c.DefaultQuery("sort", "created_at desc")

	return &dtos.Pagination{Limit: limit, Page: page, Sort: sort}
}
