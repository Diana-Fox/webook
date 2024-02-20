package service

import (
	"context"
	"fmt"
	webookgrpc "github.com/Diana-Fox/webook/script/webook/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	//这是个连接池
	clientConn, err := grpc.Dial(":8090", grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	if err != nil {
		return
	}
	client := webookgrpc.NewUserServiceClient(clientConn)
	timeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	login, err := client.Login(timeout, &webookgrpc.UserRequest{
		Id: 1234,
	})
	if err != nil {
		return
	}
	fmt.Println(login.Msg)
}
