package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		if header == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"missing token",
			})

			c.Abort()

			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token)(interface{}, error){

				return []byte(os.Getenv("JWT_SECRET")), nil

			},
		)

		if err != nil || !token.Valid {

			c.JSON(401, gin.H{
				"message":"invalid token",
			})

			c.Abort()

			return
		}

		c.Next()

	}
}