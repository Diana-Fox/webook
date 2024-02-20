package ioc

import (
	"github.com/Diana-Fox/webook/internal/web"
	"github.com/gin-gonic/gin"
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
