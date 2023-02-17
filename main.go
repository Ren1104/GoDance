package main

import (
	"GoDance/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	model.InitDB()

	model.InitTable()

	r.Run()
}
