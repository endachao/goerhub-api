package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
	"goerhubApi/constraint/requestValidate"
	"goerhubApi/controller"
	"goerhubApi/middleware"
	"goerhubApi/model"
	"gopkg.in/go-playground/validator.v9"
)

func InitRouter() *gin.Engine {

	// gin router
	gin.SetMode(viper.GetString("app.mode"))
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("username_unique", requestValidate.UserNameValidate)
		v.RegisterValidation("email_unique", requestValidate.UserEmailValidate)
	}

	user := r.Group("/user")
	{
		userController := controller.User{Model: model.UserModel{}}
		user.POST("/login", userController.Login)
		user.POST("/register", userController.Register)

		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/profile", userController.Profile)
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello GoerHub",
		})
	})
	return r
}
