package main

import (
	"jwt-authen/database"
	"jwt-authen/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	database.Connect()

	r := gin.Default()
	routes.Setup(r)
	r.Run()
}
