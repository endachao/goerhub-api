package e

import "github.com/gin-gonic/gin"

func AbortError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": err.Error(),
	})
	c.Abort()
}
