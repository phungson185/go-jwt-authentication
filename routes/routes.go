package routes

import (
	"jwt-authen/controllers"
	"jwt-authen/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
		item.POST("/items", middleware.Authentication(), controllers.CreateItem)
		item.GET("/items", middleware.Authentication(), controllers.GetAllItem)
		item.GET("/items/:id", middleware.Authentication(), controllers.GetItemById)
		item.PUT("/items/:id", middleware.Authentication(), controllers.UpdateItemById)
		item.DELETE("/items/:id", middleware.Authentication(), controllers.DeleteItemById)
		item.POST("/items/:id/buy", middleware.Authentication(), controllers.BuyItem)
	}
}
