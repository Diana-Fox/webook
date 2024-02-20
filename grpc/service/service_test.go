package service

import (
	webookgrpc "github.com/Diana-Fox/webook/script/webook/grpc"
	"google.golang.org/grpc"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	server := grpc.NewServer()
	defer func() {
		server.GracefulStop() //优雅退出
	}()
	userService := NewUserServiceRPC()
	//注册
	webookgrpc.RegisterUserServiceServer(server, userService)
	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		return
	}
}
