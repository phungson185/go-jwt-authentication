package repositories

import (
	"fmt"
	"jwt-authen/database"
	"jwt-authen/dtos"
	"jwt-authen/models"
	"math"
	"strings"
	"time"
)

type AuctionRepo struct{}

func (a *AuctionRepo) FindById(id uint32) (*models.Auction, error) {
	var auction models.Auction

	if err := database.Db.Where(&models.Auction{ID: id}).Take(&auction).Error; err != nil {
		return nil, err
	}
	return &auction, nil
}

func (a *AuctionRepo) Update(id uint32, input dtos.UpdateAuction) (*models.Auction, error) {
	var auction models.Auction

	if err := database.Db.Where(&models.Auction{ID: id}).Updates(models.Auction{InitialPrice: input.InitialPrice, FinalPrice: input.FinalPrice, EndAt: time.Unix(input.EndAt, 7)}).Find(&auction).Error; err != nil {
		return nil, err
	}
	return &auction, nil
}

func (a *AuctionRepo) Delete(id uint32) error {
	var auction models.Auction
	if err := database.Db.Where(&models.Auction{ID: id}).Delete(&auction).Error; err != nil {
		return err
	}
	return nil
}

func (a *AuctionRepo) Pagination(pagination *dtos.Pagination) (*dtos.Pagination, int, error) {
	var auctions []models.Auction

	totalPages, fromRow, toRow := 0, 0, 0

	var totalRows int64 = 0

	offset := (pagination.Page - 1) * pagination.Limit

	find := database.Db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	searchs := pagination.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break
			}
		}
	}

	if err := find.Find(&auctions).Error; err != nil {
		return nil, totalPages, err
	}

	pagination.Rows = auctions

	if err := database.Db.Model(&models.Auction{}).Count(&totalRows).Error; err != nil {
		return nil, totalPages, err
	}

	pagination.TotalRows = totalRows

	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	if pagination.Page == 1 {
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if int64(toRow) > totalRows {
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return pagination, totalPages, nil
}
