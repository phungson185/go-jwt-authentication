package controllers

import (
	"fmt"
	"jwt-authen/database"
	"jwt-authen/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type CreateItemInput struct {
	Name        string `json:"name" binding:"required`
	Description string `json:"description"`
	Price       int64  `json:"price,string" binding:"required`
	Currency    string `json:"currency" binding:"required`
	Owner       string `json:"owner" binding:"required`
	Creator     string `json:"creator" binding:"required`
	Metadata    string `json:"metadata" binding:"required`
}

type ItemModel struct{}

func CreateItem(c *gin.Context) {
	email, _ := c.Get("User")
	if email == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthenticated"})
		return
	}
	var json CreateItemInput
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	item := models.Item{
		Name:        json.Name,
		Description: json.Description,
		Price:       json.Price,
		Currency:    json.Currency,
		Owner:       fmt.Sprintf("%v", email),
		Creator:     fmt.Sprintf("%v", email),
	}

	if err := database.Db.Create(&item); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create item"})
		return
	}

	if err := database.Db.Model(&models.Item{}).Where("id = ?", item.ID).Update("metadata", fmt.Sprintf("localhost:8080/item/%d", item.ID)); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update metadata"})
		return
	}

	c.JSON(http.StatusOK, item)
}




