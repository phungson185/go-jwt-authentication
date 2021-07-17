package services

import (
	"fmt"

	"jwt-authen/dtos"
	"jwt-authen/repositories"

	"github.com/gin-gonic/gin"
)

func Pagination(context *gin.Context, pagination *dtos.Pagination) (*dtos.Pagination, error) {
	operationResult, totalPages, err := repositories.Pagination(pagination)
	if err != nil {
		return nil, err
	}

	var data = operationResult

	urlPath := context.Request.URL.Path

	searchQueryParams := ""

	for _, search := range pagination.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 1, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort) + searchQueryParams

	if data.Page > 0 {
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQueryParams
	}

	if data.Page >= totalPages {
		data.PreviousPage = ""
	}

	return data, nil
}
