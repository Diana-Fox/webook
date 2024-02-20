package cache

import (
	"context"
	"github.com/Diana-Fox/webook/internal/domian"
)

type UserCache interface {
	Get(ctx context.Context, id int64)
	Set(ctx context.Context, user domian.User)
}
