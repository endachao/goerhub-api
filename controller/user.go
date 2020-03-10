package controller

import (
	"github.com/gin-gonic/gin"
	"goerhubApi/constraint"
	"goerhubApi/constraint/e"
	"goerhubApi/middleware/auth"
	"goerhubApi/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	Model model.UserModel
}

func (u *User) Login(c *gin.Context) {
	var loginRequest constraint.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		e.AbortError(c, 400, err)
		return
	}
	user, ok := u.Model.GetUserInfoByUserEmail(loginRequest.Email)
	if !ok {
		e.AbortError(c, 400, e.ErrFailedAuthentication)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		e.AbortError(c, 400, e.ErrFailedAuthentication)
		return
	}

	token, _ := auth.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "",
		"data": map[string]interface{}{
			"auth": token,
		},
	})
}

func (u *User) Register(c *gin.Context) {
	var registerRequest constraint.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		e.AbortError(c, 400, err)
		return
	}

	password := []byte(registerRequest.Password)
	password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		e.AbortError(c, 400, err)
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
		e.AbortError(c, 400, err)
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
	})
}

func (u *User) Profile(c *gin.Context) {
	userId, exist := auth.GetUserId(c)
	if !exist {
		e.AbortError(c, 400, e.ErrForbidden)
	}
	user := u.Model.GetUserInfoByUserId(userId)
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
