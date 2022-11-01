package api

import "github.com/gin-gonic/gin"

type TestApi struct {
}

func (t *TestApi) Test(c *gin.Context) {
	c.String(200, "hello,world")
}
