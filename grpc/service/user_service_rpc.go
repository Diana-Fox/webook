package service

import (
	"context"
	"fmt"
	"github.com/Diana-Fox/webook/script/webook/grpc"
)

type UserServiceRPC struct {
	webookgrpc.UnimplementedUserServiceServer
}

func NewUserServiceRPC() webookgrpc.UserServiceServer {
	return UserServiceRPC{}
}
func (u UserServiceRPC) Login(ctx context.Context, request *webookgrpc.UserRequest) (*webookgrpc.UserResponse, error) {
	fmt.Println(request)
	return &webookgrpc.UserResponse{
		Msg: "正常返回",
	}, nil
}
