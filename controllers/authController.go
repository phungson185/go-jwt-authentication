package controllers

import (
	"fmt"
	"jwt-authen/database"
	"jwt-authen/dtos"
	"jwt-authen/models"
	"jwt-authen/services"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"math/rand"
)

const SecretKey = "secret"

func Register(c *gin.Context) {

	var json dtos.CreateUser

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(999999) + 100000

	password, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user := models.User{
		Email:       json.Email,
		Password:    string(password),
		Phone:       json.Phone,
		UserAddress: json.UserAddress,
		VerifyCode:  strconv.Itoa(randNum),
	}

	if err := database.Db.Create(&user); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not register"})
		return
	}

	content := "Your authentication code is: " + strconv.Itoa(randNum)
	services.SendMail(json.Email, "Verify Email", content)

	c.JSON(http.StatusOK, dtos.Response(true, "Success", &user))

}

func VerifyEmail(c *gin.Context) {
	var json dtos.VerifyEmail
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User

	if err := database.Db.Model(&user).Where("email = ? AND verify_code = ?", json.Email, json.VerifyCode).Update("status", true); err.Error != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "Could not verify", nil))
		return
	}

	c.JSON(http.StatusOK, dtos.Response(true, "Success", nil))
}

func Login(c *gin.Context) {
	var json dtos.Login

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.Db.Model(&user).Where("email = ? AND status = true", json.Email).First(&user); err.Error != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "User not found", nil))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password)); err != nil {
		c.JSON(http.StatusNotFound, dtos.Response(false, "Incorrect password", nil))
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.Response(false, "Could not login", nil))
		return
	}

	c.SetCookie("jwt", token, int(time.Now().Add(time.Hour*24).Unix()), "/", "localhost", false, true)

	c.JSON(http.StatusOK, dtos.Response(true, "Success", token))
}

func Profile(c *gin.Context) {
	email, _ := c.Get("User")
	var user models.User

	database.Db.Where("email = ?", fmt.Sprintf("%v", email)).First(&user)

	c.JSON(http.StatusOK, dtos.Response(true, "Success", user))
}
