//go:build wireinject

package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"webook/internal/repository"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/internal/web"
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
