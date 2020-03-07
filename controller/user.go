package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"goerhubApi/constraint"
	"goerhubApi/model"
)

type User struct {
}

func (u *User) Login(c *gin.Context) (interface{}, error) {
	var loginRequest constraint.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		return nil, err
	}
	user, ok := model.UserModel{}.GetUserInfoByUserEmail(loginRequest.Email)
	if !ok {
		return nil, jwt.ErrFailedAuthentication
	}

	return &user, nil
}
