package main

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(utils.LoggerWrite())
	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
	router.Run(":8080")
}
