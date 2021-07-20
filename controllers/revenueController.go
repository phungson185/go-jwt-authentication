package controllers

import (
	"jwt-authen/dtos"
	"jwt-authen/helpers"
	"jwt-authen/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var RevenueRepo = new(repositories.RevenueRepo)

func Revenue(c *gin.Context) {

	time := helpers.GenerateTimeRequest(c)

	revenue, err := RevenueRepo.CalculateRevenue(c, time)

	if err != nil {
		c.JSON(http.StatusOK, dtos.Response(true, "Success", 0))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", revenue))
}
