package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"webook/internal/domian"
)

type userCache struct {
	client redis.Cmdable
}

func NewUserCache(client redis.Cmdable) UserCache {
	return &userCache{
		client: client,
	}
}
func (u *userCache) Get(ctx context.Context, id int64) {
	//TODO implement me
	panic("implement me")
}

func (u *userCache) Set(ctx context.Context, user domian.User) {
	//TODO implement me
	panic("implement me")
}
