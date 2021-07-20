package repositories

import (
	"jwt-authen/database"
	"jwt-authen/dtos"
	"jwt-authen/models"

	"github.com/gin-gonic/gin"
)

type RevenueRepo struct{}

func (r *RevenueRepo) CalculateRevenue(c *gin.Context, t *dtos.Time) (interface{}, error) {

	var revenue float64

	if err := database.Db.Model(&models.Transaction{}).Select("sum(price)").Where("date_part('day', created_at) = ? AND date_part('month', created_at) = ? AND date_part('year', created_at) = ?", t.Day, t.Month, t.Year).Find(&revenue).Error; err != nil && t.Type_ == "particular" {
		return nil, err
	}

	if err := database.Db.Model(&models.Transaction{}).Select("sum(price)").Where("date(created_at) BETWEEN ? AND ?", t.From, t.To).Find(&revenue).Error; err != nil && t.Type_ == "range" {
		return nil, err
	}

	if err := database.Db.Model(&models.Transaction{}).Select("sum(price)").Where("date_part('week', created_at) = ?", t.Week).Find(&revenue).Error; err != nil && t.Type_ == "week" {
		return nil, err
	}

	return revenue, nil
}
