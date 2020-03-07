package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goerhubApi/controller"
	"goerhubApi/middleware"
	"goerhubApi/model"
)

func InitRouter() *gin.Engine {

	// gin router
	gin.SetMode(viper.GetString("app.mode"))
	r := gin.Default()

	user := r.Group("/user")
	{
		// controller and middleware
		userController := controller.User{Model: model.UserModel{}}
		authMiddleware, err := middleware.AuthMiddleware(userController.Login)
		if err != nil {
			panic(err)
		}
		user.POST("/login", authMiddleware.LoginHandler)
		user.POST("/register", userController.Register)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello GoerHub",
		})
	})
	return r
}
