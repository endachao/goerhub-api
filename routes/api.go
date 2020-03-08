package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
	"goerhubApi/controller"
	"goerhubApi/middleware"
	"goerhubApi/middleware/validators"
	"goerhubApi/model"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"time"
)

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func InitRouter() *gin.Engine {

	// gin router
	gin.SetMode(viper.GetString("app.mode"))
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("username_unique", validators.UserNameValidate)
		v.RegisterValidation("email_unique", validators.UserEmailValidate)
	}

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
