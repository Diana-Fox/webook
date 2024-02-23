package service

import (
	"context"
	"fmt"
	"github.com/Diana-Fox/webook/script/webook/grpc"
)

type UserServiceRPC struct {
	webookgrpc.UnimplementedUserServiceServer
	Name string
}

func NewUserServiceRPC(name string) webookgrpc.UserServiceServer {
	return UserServiceRPC{
		Name: name,
	}
}
func (u UserServiceRPC) Login(ctx context.Context, request *webookgrpc.UserRequest) (*webookgrpc.UserResponse, error) {
	fmt.Println(request)
	return &webookgrpc.UserResponse{
		Msg: "正常返回" + u.Name,
	}, nil
}
