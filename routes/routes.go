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

	auction := r.Group("/auctions")
	{
		auction.POST("/items/:id", middleware.Authentication(), controllers.CreateAuction)
		auction.GET("", middleware.Authentication(), controllers.GetAllAuction)
		auction.GET("/:id", middleware.Authentication(), controllers.GetAuctionById)
		auction.PUT("/:id", middleware.Authentication(), controllers.UpdateAuctionById)
		auction.DELETE("/:id", middleware.Authentication(), controllers.DeleteAuctionById)
	}

	r.GET("revenue", middleware.Authentication(), controllers.Revenue)

	statistic := r.Group("/statistic")
	{
		statistic.GET("/users", middleware.Authentication(), controllers.TotalUser)
		statistic.GET("/newestitem", middleware.Authentication(), controllers.NewestItem)
		statistic.GET("/newestauction", middleware.Authentication(), controllers.NewestAuction)
		statistic.GET("/sellestitem", middleware.Authentication(), controllers.BestSellingItem)
	}
}
