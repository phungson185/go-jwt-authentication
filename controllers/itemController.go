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

func CreateItem(c *gin.Context) {
	email, _ := c.Get("User")
	// if email == nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthenticated"})
	// 	return
	// }
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

	if err := database.Db.Model(&models.Item{}).Where("id = ?", item.ID).Update("metadata", fmt.Sprintf("localhost:8080/item/%d", item.ID)); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update metadata"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func GetAllItem(c *gin.Context) {

	pagination := helpers.GeneratePaginationRequest(c)

	response, err := services.Pagination(c, pagination)

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
	res, err := repositories.FindById(uint32(id))

	if err != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "ID not found", nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Get Item success", res))
}

func UpdateItemById(c *gin.Context) {

	email, _ := c.Get("User")

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Invalid ID", nil))
		return
	}

	res, err := repositories.FindById(uint32(id))

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

	res, err = repositories.Update(uint32(id), input)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, "Update Item Failed", nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", res))
}
