package service

import (
	"context"
	"github.com/Diana-Fox/webook/internal/domian"
	"github.com/Diana-Fox/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{ur: ur}
}

func (u *userService) SignUp(ctx context.Context, user domian.User) error {
	//context.WithTimeout(ctx,time.Minute)//可以设置超时时间
	//要加密
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)
	err = u.ur.Create(ctx, user)
	return err
}
