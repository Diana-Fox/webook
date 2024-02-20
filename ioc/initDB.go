package ioc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"webook/internal/repository"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/internal/web"
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
