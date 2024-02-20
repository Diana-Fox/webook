//go:build wireinject

package ioc

import (
	"github.com/Diana-Fox/webook/internal/repository"
	"github.com/Diana-Fox/webook/internal/repository/dao"
	"github.com/Diana-Fox/webook/internal/service"
	"github.com/Diana-Fox/webook/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitServer() *gin.Engine {
	wire.Build(
		InitDB,
		dao.NewUserDao,
		repository.NewUserRepository,
		service.NewUserService,
		web.NewUserHandler,
		InitGin,
		InitMiddlewares,
	)
	return new(gin.Engine)
}

//func InitUserHandler() web.UserHandler {
//	db := InitDB()
//	ud := dao.NewUserDao(db)
//	ur := repository.NewUserRepository(ud)
//	us := service.NewUserService(ur)
//	handler := web.NewUserHandler(us)
//	return handler
//}
