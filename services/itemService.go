package services

import (
	"fmt"

	"jwt-authen/dtos"
	"jwt-authen/repositories"

	"github.com/gin-gonic/gin"
)

func Pagination(context *gin.Context, pagination *dtos.Pagination) dtos.Response {
	operationResult, totalPages := repositories.Pagination(pagination)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*dtos.Pagination)

	urlPath := context.Request.URL.Path

	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort)
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort)

	if data.Page > 0 {
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort)
	}

	if data.Page < totalPages {
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort)
	}

	if data.Page > totalPages {
		data.PreviousPage = ""
	}

	return dtos.Response{Success: true, Data: data}
}
