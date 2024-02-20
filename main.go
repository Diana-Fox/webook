package main

import (
	"github.com/Diana-Fox/webook/ioc"
	"github.com/gin-gonic/gin"
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
