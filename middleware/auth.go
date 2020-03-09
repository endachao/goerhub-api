package middleware

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"goerhubApi/constraint/e"
	"goerhubApi/helpers"
	"log"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			e.AbortError(c, 401, errors.New("token is empty"))
			return
		}
		n := token[7:]
		log.Printf("%s\n", token)
		log.Printf("%s\n", n)
		claims, err := helpers.ParseToken(n)
		if err != nil {
			e.AbortError(c, 401, jwt.ErrInvalidSigningAlgorithm)
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			e.AbortError(c, 401, jwt.ErrExpiredToken)
			return
		}

		c.Next()
	}
}
