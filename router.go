package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func initRouter() *gin.Engine {
	gin.SetMode(viper.GetString("app.mode"))
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello GoerHub",
		})
	})
	return r
}
