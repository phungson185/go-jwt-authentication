package middleware

import (
	"jwt-authen/database"
	"jwt-authen/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const SecretKey = "secret"

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")

		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthenticated"})
			c.Abort()
		}

		claims := token.Claims.(*jwt.StandardClaims)

		var user models.User

		if err := database.Db.Where("email = ?", claims.Issuer).First(&user); err.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			c.Abort()
		}

		c.Set("User", user.Email)

		c.Next()
	}
}
