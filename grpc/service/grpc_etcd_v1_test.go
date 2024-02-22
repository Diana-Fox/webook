package service

import (
	"context"
	"fmt"
	webookgrpc "github.com/Diana-Fox/webook/script/webook/grpc"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
	"time"
)

type EtcdTestSuiteV1 struct {
	suite.Suite
	client *etcdv3.Client
}

func (s *EtcdTestSuiteV1) SetupSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//这个是etcd
	client, err := etcdv3.New(etcdv3.Config{
		Endpoints: []string{"127.0.0.1:12379"},
		Context:   ctx,
	})
	require.NoError(s.T(), err)
	s.client = client //初始化客户端
}
func (s *EtcdTestSuiteV1) TestServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//过期时间
	var ttl int64 = 60 //单位是s
	leaseRes, err := s.client.Grant(ctx, ttl)
	if err != nil {
		return
	}
	//s.SetupSuite()
	//target是当前服务注册的服务名，一个服务一个manager
	manager, err := endpoints.NewManager(s.client, "service/user")
	require.NoError(s.T(), err)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//key是指实例的key,有instance id就是，没有就是ip
	addr := "127.0.0.1:8090"
	key := "service/user" + addr
	err = manager.AddEndpoint(ctx, key, endpoints.Endpoint{
		Addr: addr,
		//Metadata: UserServiceRPC{},//这个可以存放一些自定义的信息
	}, etcdv3.WithLease(leaseRes.ID))
	require.NoError(s.T(), err)
	go func() {
		//续约
		ctxKl, cancelKl := context.WithTimeout(context.Background(), time.Second)
		defer cancelKl()
		//默认1/3时间开始续约
		ch, errKl := s.client.KeepAlive(ctxKl, leaseRes.ID)
		require.NoError(s.T(), errKl)
		for kaResp := range ch {
			s.T().Log(kaResp.String())
		}
	}()
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
func (s *EtcdTestSuiteV1) TestEtcdClient() {
	db, err := resolver.NewBuilder(s.client)
	require.NoError(s.T(), err)
	clientConn, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(db),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(s.T(), err)
	client := webookgrpc.NewUserServiceClient(clientConn)
	timeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	login, err := client.Login(timeout, &webookgrpc.UserRequest{
		Id: 432,
	})
	fmt.Println(login)
}
func TestEtcdV1(t *testing.T) {
	suite.Run(t, new(EtcdTestSuiteV1))
}
