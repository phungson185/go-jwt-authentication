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

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var ItemRepo = new(repositories.ItemRepo)

func CreateItem(c *gin.Context) {
	email, _ := c.Get("User")

	var json dtos.CreateItem
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

	item.Metadata = fmt.Sprintf("localhost:8080/item/%d", item.ID)
	database.Db.Save(&item)

	c.JSON(http.StatusOK, dtos.Response(true, "Success", item))
}

func GetAllItem(c *gin.Context) {

	pagination := helpers.GeneratePaginationRequest(c)

	operationResult, totalPages, err := ItemRepo.Pagination(pagination)

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

func GetItemById(c *gin.Context) {
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

	c.JSON(http.StatusOK, dtos.Response(true, "Get Item Success", res))
}

func UpdateItemById(c *gin.Context) {

	email, _ := c.Get("User")

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

	if res.Owner != email {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "You aren't owner of item", nil))
		return
	}

	var input dtos.UpdateItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err = ItemRepo.Update(uint32(id), input)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Update Item Failed", nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", res))
}

func DeleteItemById(c *gin.Context) {

	email, _ := c.Get("User")

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

	if res.Owner != email {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "You aren't owner of item", nil))
		return
	}

	err = ItemRepo.Delete(uint32(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Delete Item Failed", nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", nil))
}

func BuyItem(c *gin.Context) {
	email, _ := c.Get("User")

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

	if res.Type == "Auction" {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Could not buy", nil))
		return
	}

	if res.Owner == email {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "You are owner of item", nil))
		return
	}

	hash := helpers.NewSHA1Hash()

	transaction := models.Transaction{
		ItemID: res.ID,
		TxHash: hash,
		Buyer:  fmt.Sprintf("%v", email),
		Seller: res.Owner,
		Price:  uint64(res.Price),
		Fee:    float64(res.Price) * 0.1,
	}

	if err := database.Db.Create(&transaction); err.Error != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response(false, "Transaction failed", nil))
		return
	}

	res.Owner = fmt.Sprintf("%v", email)
	res.Status = "Success"
	res.Type = "Fixed"

	if err := database.Db.Save(&res).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response(false, "Update item failed", nil))
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", transaction))
}

func ItemTransaction(c *gin.Context) {

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

	pagination := helpers.GeneratePaginationRequest(c)

	operationResult, totalPages, err := repositories.TransactionPagination(pagination, uint32(id))

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
