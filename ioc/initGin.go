package ioc

import (
	"github.com/gin-gonic/gin"
	"webook/internal/web"
)

func InitGin(mdls []gin.HandlerFunc, hdl web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	hdl.RegisterRoutes(server)
	return server
}
func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
