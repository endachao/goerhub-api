package middleware

import (
	"github.com/gin-gonic/gin"
	"goerhubApi/constraint/e"
	"goerhubApi/helpers/auth"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			e.AbortError(c, 401, e.ErrEmptyAuthHeader)
			return
		}
		n := token[7:]
		claims, err := auth.ParseToken(n)
		if err != nil {
			e.AbortError(c, 401, e.ErrInvalidSigningAlgorithm)
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			e.AbortError(c, 401, e.ErrExpiredToken)
			return
		}
		c.Set("JWT-AUTH-USER", claims)
		c.Next()
	}
}
