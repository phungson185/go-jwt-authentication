package repositories

import (
	"math"

	"jwt-authen/database"
	"jwt-authen/dtos"
	"jwt-authen/models"
)

func Pagination(pagination *dtos.Pagination) (RepositoryResult, int) {
	var items models.Item

	totalPages, fromRow, toRow := 0, 0, 0

	var totalRows int64 = 0

	offset := pagination.Page * pagination.Limit
	if err := database.Db.Model(&models.Item{}).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&items); err.Error != nil {
		return RepositoryResult{Error: err.Error}, totalPages
	}

	pagination.Rows = items

	if err := database.Db.Model(&models.Item{}).Count(&totalRows); err.Error != nil {
		return RepositoryResult{Error: err.Error}, totalPages
	}

	pagination.TotalRows = totalRows

	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if int64(toRow) > totalRows {
		// set to row with total rows
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
