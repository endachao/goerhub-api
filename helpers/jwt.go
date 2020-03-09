package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
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
