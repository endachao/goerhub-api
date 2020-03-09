package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"goerhubApi/constraint"
	"goerhubApi/model"
	"time"
)

const identityKey = "email"

func AuthMiddleware(loginFunc constraint.LoginHandleFunc, loginResponse func(*gin.Context, int, string, time.Time)) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "goer hub",
		Key:         []byte("dd8Ub1JJkes7EJZawpFEknCnykW6s7Co"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.Users); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.Users{
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: loginFunc,
		LoginResponse: loginResponse,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.Users); ok && v.Username == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
