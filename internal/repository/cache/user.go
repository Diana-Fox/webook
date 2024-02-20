package cache

import (
	"context"
	"github.com/Diana-Fox/webook/internal/domian"
	"github.com/redis/go-redis/v9"
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
