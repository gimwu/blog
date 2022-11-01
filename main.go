package main

import (
	"blog/bootstrap"
	"blog/router"
	_ "blog/router"
)

func main() {
	bootstrap.Init()
	rootRouter := router.Init()
	rootRouter.Run(":8080")
}
