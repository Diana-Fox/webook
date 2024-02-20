package dao

import (
	"context"
	"webook/internal/domian"
)

type UserDao interface {
	Insert(ctx context.Context, user domian.User) error
}
