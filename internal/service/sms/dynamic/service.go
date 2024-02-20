package dynamic

import (
	"context"
	"webook/internal/service/sms"
)

type Service struct {
	svc sms.SMS
	//切换器，判断是否切换
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
