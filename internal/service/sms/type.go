package sms

import "context"

type SMS interface {
	Send(ctx context.Context, tpl string,
		args []string, numbers ...string)
}
