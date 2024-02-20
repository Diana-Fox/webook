package repository

import (
	"context"
	"github.com/Diana-Fox/webook/internal/domian"
	"github.com/Diana-Fox/webook/internal/repository/dao"
)

type userRepository struct {
	ud dao.UserDao
}

func (u *userRepository) LoginByPassword(ctx context.Context, user domian.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(ud dao.UserDao) UserRepository {
	return &userRepository{
		ud: ud,
	}
}
func (u *userRepository) Create(ctx context.Context, user domian.User) error {
	err := u.ud.Insert(ctx, user)
	return err
}
