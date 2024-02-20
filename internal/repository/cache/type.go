package cache

import (
	"context"
	"webook/internal/domian"
)

type UserCache interface {
	Get(ctx context.Context, id int64)
	Set(ctx context.Context, user domian.User)
}
