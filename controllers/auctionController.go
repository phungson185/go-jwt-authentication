package controllers

import (
	"fmt"
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

	res, err := ItemRepo.FindById(uint32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "ID not found", nil))
		return
	}

	if res.Status == "Success" {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Item is already on marketplace", nil))
		return
	}

	res.Status = "Success"
	res.Type = "Auction"

	if err := database.Db.Save(&res).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response(false, "Update Item Failed", nil))
		return
	}

	auction := models.Auction{
		ItemID:       uint32(id),
		InitialPrice: float64(res.Price),
		FinalPrice:   input.FinalPrice,
		EndAt:        time.Unix(input.EndAt, 7),
	}

	if err := database.Db.Create(&auction); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create auction"})
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", auction))
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

	_, err = AuctionRepo.FindById(uint32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "ID not found", nil))
		return
	}

	var input dtos.UpdateAuction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := AuctionRepo.Update(uint32(id), input)

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

func Bid(c *gin.Context) {

	email, _ := c.Get("User")

	var json dtos.CreateBid
	var highestBid float64

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	if err := database.Db.Model(&models.Bid{}).Select("price").Where(&models.Bid{AuctionID: uint32(id)}).Order("created_at desc").Limit(1).Find(&highestBid).Error; err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, err.Error(), nil))
		return
	}

	if json.Price <= float64(highestBid) && (res.Status == "Success") {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Invalid bid price", nil))
		return
	} else if json.Price <= float64(res.InitialPrice) && (res.Status == "Pending") {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Invalid bid price", nil))
		return
	}

	hash := helpers.NewSHA1Hash()

	bid := models.Bid{
		AuctionID: uint32(id),
		Bidder:    fmt.Sprintf("%v", email),
		Price:     json.Price,
		TxHash:    hash,
		Fee:       float64(json.Price) * 0.1,
	}

	if err := database.Db.Create(&bid); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create bid"})
		return
	}

	res.Status = "Success"

	if err := database.Db.Save(&res); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update auction"})
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", bid))
}
