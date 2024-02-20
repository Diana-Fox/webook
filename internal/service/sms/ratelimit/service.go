package ratelimit

import (
	"context"
	"webook/internal/service/sms"
	"webook/pkg/ginx/ratelimit"
)

type Service struct {
	svc sms.SMS
	r   ratelimit.Limiter //限流器
}

func NewService(svc sms.SMS) sms.SMS {
	return &Service{
		svc: svc,
	}
}

func (s Service) Send(ctx context.Context, tpl string, args []string, numbers ...string) {
	//TODO implement me
	panic("implement me")
}
