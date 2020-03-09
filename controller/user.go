package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"goerhubApi/constraint"
	"goerhubApi/model"
	"golang.org/x/crypto/bcrypt"
	"time"
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

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
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

func (u *User) Profile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user := u.Model.GetUserInfoByUserId(int(claims["pk"].(float64)))
	c.JSON(200, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"username":    user.Username,
			"nickname":    user.Nickname,
			"email":       user.Email,
			"gold_number": user.GoldNumber,
		},
	})
}

func (u *User) LoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(code, gin.H{
		"code": code,
		"data": map[string]string{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		},
	})
}
