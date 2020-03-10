package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goerhubApi/constraint/e"
	"time"
)

type Claims struct {
	UserId int
	jwt.StandardClaims
}

type TokenMap struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

var jwtSecret = []byte(viper.GetString("jwt.secret"))

func GenerateToken(userId int) (TokenMap, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(72 * time.Hour).Unix()

	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    viper.GetString("app.name"),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return TokenMap{
		Token:     token,
		ExpiresAt: expireTime,
	}, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GetUserId(c *gin.Context) (value int, exists bool) {
	claims, exist := c.Get("JWT-AUTH-USER")
	if !exist {
		return
	}
	return claims.(*Claims).UserId, true
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			e.AbortError(c, 401, e.ErrEmptyAuthHeader)
			return
		}
		n := token[7:]
		claims, err := ParseToken(n)
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
