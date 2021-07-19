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
		user.POST("/verifyEmail", controllers.VerifyEmail)
		user.POST("/login", controllers.Login)
		user.GET("/profile", controllers.Profile)
	}

	item := r.Group("/items")
	{
		item.POST("", middleware.Authentication(), controllers.CreateItem)
		item.GET("", middleware.Authentication(), controllers.GetAllItem)
		item.GET("/:id", middleware.Authentication(), controllers.GetItemById)
		item.PUT("/:id", middleware.Authentication(), controllers.UpdateItemById)
		item.DELETE("/:id", middleware.Authentication(), controllers.DeleteItemById)
		item.POST("/:id/buy", middleware.Authentication(), controllers.BuyItem)
		item.GET("/:id/transactions", middleware.Authentication(), controllers.ItemTransaction)
	}

	auction := r.Group("/auction/items")
	{
		auction.POST("/:id", middleware.Authentication(), controllers.CreateAuction)
	}
}
