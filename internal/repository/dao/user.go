package dao

import (
	"context"
	"github.com/Diana-Fox/webook/internal/domian"
	"gorm.io/gorm"
	"time"
)

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}
func (u *userDao) getTable() string {
	return "sys_user"
}

type User struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	EmailValidated int    `json:"emailValidated"`
	Phone          string `json:"phone"`
	PhoneValidated int    `json:"phoneValidated"`
	Status         int    `json:"status"`
	CreateTime     int64  `json:"createTime"`
}

func (u *userDao) Insert(ctx context.Context, user domian.User) error {
	var us = User{
		Name:           user.Name,
		Password:       user.Password,
		EmailValidated: 0,
		PhoneValidated: 0,
		Status:         1,
		CreateTime:     time.Now().UnixMilli(),
	}
	err := u.db.Table(u.getTable()).WithContext(ctx).Create(&us).Error
	return err
}
