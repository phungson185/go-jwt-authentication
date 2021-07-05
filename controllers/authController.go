package controllers

import (
	"jwt-authen/database"
	"jwt-authen/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Register(c *gin.Context) {
	var json map[string]string
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user := models.User{
		Email:       json["email"],
		Password:    json["password"],
		Phone:       json["phone"],
		UserAddress: json["userAddress"],
	}
	database.Db.Create(&user)
	c.JSON(http.StatusOK, json)
}

func VerifyEmail(c *gin.Context) {
	var json map[string]string
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var user models.User

	if err := database.Db.Model(&user).Where("email = ?", json["email"]).Update("status", true); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
