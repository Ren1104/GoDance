package main

import (
	"GoDance/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	controller.InitDB()

	r.Run()
}
