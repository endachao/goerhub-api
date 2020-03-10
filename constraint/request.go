package constraint

import (
	"github.com/gin-gonic/gin"
)

type LoginHandleFunc func(c *gin.Context) (interface{}, error)
type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email           string `form:"email" json:"email" binding:"required,email,email_unique"`
	Username        string `form:"username" json:"username" binding:"required,username_unique"`
	Password        string `form:"password" json:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"eqfield=Password"`
}
