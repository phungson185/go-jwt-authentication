package controllers

import (
	"jwt-authen/database"
	"jwt-authen/dtos"
	"jwt-authen/helpers"
	"jwt-authen/models"
	"jwt-authen/repositories"
	"jwt-authen/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var AuctionRepo = new(repositories.AuctionRepo)

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

	_, err = ItemRepo.FindById(uint32(id))

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

func GetAllAuction(c *gin.Context) {

	pagination := helpers.GeneratePaginationRequest(c)

	operationResult, totalPages, err := AuctionRepo.Pagination(pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, err.Error(), nil))
	}

	response, err := services.Pagination(c, operationResult, totalPages)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", response))
}

func GetAuctionById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Invalid ID", nil))
		return
	}
	res, err := AuctionRepo.FindById(uint32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "ID not found", nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Get Auction Success", res))
}

func UpdateAuctionById(c *gin.Context) {

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Invalid ID", nil))
		return
	}

	res, err := AuctionRepo.FindById(uint32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "ID not found", nil))
		return
	}

	var input dtos.UpdateAuction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err = AuctionRepo.Update(uint32(id), input)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Update Auction Failed", nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", res))
}

func DeleteAuctionById(c *gin.Context) {

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Invalid ID", nil))
		return
	}

	_, err = AuctionRepo.FindById(uint32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "ID not found", nil))
		return
	}

	err = AuctionRepo.Delete(uint32(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Delete Auction Failed", nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", nil))
}
