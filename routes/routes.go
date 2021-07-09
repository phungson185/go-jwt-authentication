package routes

import (
	"jwt-authen/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Setup(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/verifyEmail", controllers.VerifyEmail)
	r.POST("/auth/login", controllers.Login)
	r.GET("/auth/profile", controllers.Profile)
}
