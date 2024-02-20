package main

import (
	"github.com/gin-gonic/gin"
	"webook/ioc"
)

func main() {
	server := gin.Default()
	//uh := web.NewUserHandler()
	//uh.RegisterRoutes(server)
	//uh := ioc.InitUserHandler()
	//uh.RegisterRoutes(server)
	server = ioc.InitServer()
	err := server.Run(":18080")
	if err != nil {
		panic(err)
	}
}
