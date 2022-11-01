package router

import (
	"blog/api"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(utils.LoggerWrite())

	testGroup := router.Group("/test")
	testApi := &api.TestApi{}
	testGroup.GET("/", testApi.Test)
	return router
}
