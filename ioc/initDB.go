package ioc

import (
	"github.com/Diana-Fox/webook/internal/repository"
	"github.com/Diana-Fox/webook/internal/repository/dao"
	"github.com/Diana-Fox/webook/internal/service"
	"github.com/Diana-Fox/webook/internal/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	//viper.GetString("")
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/webook?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
func InitUserHandler() web.UserHandler {
	db := InitDB()
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	handler := web.NewUserHandler(us)
	return handler
}
