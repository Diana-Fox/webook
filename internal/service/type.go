package service

import (
	"context"
	"github.com/Diana-Fox/webook/internal/domian"
)

type UserService interface {
	SignUp(ctx context.Context, user domian.User) error
}
type CodeService interface {
	SendCodeByEmail(ctx context.Context) error
}
