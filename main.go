package main

import (
	"douyin/config"
	"douyin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	err := config.InitMySQL()
	if err != nil {
		println(err.Error())
		return
	}
	println("Successfully connect to the database")

	//删除Message部分
	//go service.RunMessageServer()

	r := gin.Default()

	routes.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
