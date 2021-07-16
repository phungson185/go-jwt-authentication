package database

import (
	"fmt"
	"jwt-authen/models"
	"log"
	"os"

	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	con, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%v user=%v password=%v port=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_PORT"))), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db = con
	con.AutoMigrate(&models.User{})
	con.AutoMigrate(&models.Item{})
	fmt.Println("Successfully connected!")
}
