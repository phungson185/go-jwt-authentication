package routes

import (
	"jwt-authen/controllers"
	"jwt-authen/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"jwt-authen/helpers"
	"jwt-authen/services"
)

func Setup(r *gin.Engine) {

	user := r.Group("/auth")
	{
		user.POST("/register", controllers.Register)
		user.POST("verifyEmail", controllers.VerifyEmail)
		user.POST("/login", controllers.Login)
		user.GET("/profile", controllers.Profile)
	}

	item := r.Group("/")
	{
		item.POST("/item", middleware.Authentication(), controllers.CreateItem)
	}

	r.GET("/pagination", func(context *gin.Context) {
		code := http.StatusOK

		pagination := helpers.GeneratePaginationRequest(context)

		response := services.Pagination(context, pagination)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
}
