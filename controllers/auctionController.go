package controllers

import (
	"jwt-authen/database"
	"jwt-authen/dtos"
	"jwt-authen/models"
	"jwt-authen/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateAuction(c *gin.Context) {
	var input dtos.CreateAuction

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Invalid ID", nil))
		return
	}

	_, err = repositories.FindById(uint32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "ID not found", nil))
		return
	}

	auction := models.Auction{
		ItemID:       uint32(id),
		InitialPrice: input.InitialPrice,
		FinalPrice:   input.FinalPrice,
		EndAt:        time.Unix(input.EndAt, 7),
	}

	if err := database.Db.Create(&auction); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create auction"})
		return
	}

	c.JSON(http.StatusOK, auction)
}
