package repository

import (
	"context"
	"github.com/Diana-Fox/webook/internal/domian"
)

type UserRepository interface {
	Create(ctx context.Context, user domian.User) error
	LoginByPassword(ctx context.Context, user domian.User) (bool, error)
}
