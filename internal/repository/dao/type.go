package dao

import (
	"context"
	"github.com/Diana-Fox/webook/internal/domian"
)

type UserDao interface {
	Insert(ctx context.Context, user domian.User) error
}
