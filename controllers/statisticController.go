package controllers

import (
	"jwt-authen/dtos"
	"jwt-authen/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var StatisticRepo = new(repositories.StatisticRepo)

func TotalUser(c *gin.Context) {

	res, err := StatisticRepo.UserRegisterInADay()

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", res))
}

func NewestItem(c *gin.Context) {

	res, err := StatisticRepo.ListNewestItem()

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", res))
}

func NewestAuction(c *gin.Context) {

	res, err := StatisticRepo.ListNewestAuction()

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", res))
}
