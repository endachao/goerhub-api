package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goerhubApi/controller"
	"goerhubApi/middleware"
)

func InitRouter() *gin.Engine {
	gin.SetMode(viper.GetString("app.mode"))
	r := gin.Default()
	userController := controller.User{}
	authMiddleware, err := middleware.AuthMiddleware(userController.Login)
	if err != nil {
		panic(err)
	}
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello GoerHub",
		})
	})
	return r
}
