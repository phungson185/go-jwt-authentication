package database

import (
	"fmt"
	"jwt-authen/models"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {

	dsn := "host=localhost user=postgres password=20184189 port=5432 sslmode=disable"
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db = con
	con.AutoMigrate(&models.User{})
	fmt.Println("Successfully connected!")
}
