package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"goerhubApi/constraint"
	"goerhubApi/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model model.UserModel
}

func (u *User) Login(c *gin.Context) (interface{}, error) {
	var loginRequest constraint.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		return nil, err
	}
	user, ok := u.Model.GetUserInfoByUserEmail(loginRequest.Email)
	if !ok {
		return nil, jwt.ErrFailedAuthentication
	}

	return &user, nil
}

func (u *User) Register(c *gin.Context) {
	var registerRequest constraint.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		responseError(c, 400, err.Error())
		return
	}

	password := []byte(registerRequest.Password)
	password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		responseError(c, 400, err.Error())
		return
	}

	user := model.Users{
		Nickname: registerRequest.Username,
		Username: registerRequest.Username,
		Password: string(password),
		Email:    registerRequest.Email,
	}

	err = u.Model.CreateUser(user)
	if err != nil {
		responseError(c, 400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
	})
}
