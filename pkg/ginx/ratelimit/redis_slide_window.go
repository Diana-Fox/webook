package ratelimit

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"time"
)

//go:embed slide_window.lua
var luaSlideWindow string

type RateLimit struct {
	cmd redis.Cmdable
	//窗口大小
	interval time.Duration
	// 阈值
	rate int
	//interval内允许rate个请求
}

func (r RateLimit) Limit(ctx context.Context, key string) (bool, error) {
	return r.cmd.Eval(ctx, luaSlideWindow, []string{key},
		r.interval.Milliseconds(), r.rate, time.Now().UnixMilli()).Bool()
}

func NewRateLimit(cmd redis.Cmdable, interval time.Duration, rate int) Limiter {
	return &RateLimit{
		cmd:      cmd,
		interval: interval,
		rate:     rate,
	}
}