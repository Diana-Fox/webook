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
	"strconv"
	"testing"
	"time"
	//要匿名引入
	_ "google.golang.org/grpc/balancer/weightedroundrobin"
	//要匿名引入自己自定义的类
	_ "github.com/Diana-Fox/webook/pkg/grpcx/wrr"
)

type BalancerTestSuite struct {
	suite.Suite
	client *etcdv3.Client
}

func (b *BalancerTestSuite) SetupSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//这个是etcd
	client, err := etcdv3.New(etcdv3.Config{
		Endpoints: []string{"127.0.0.1:12379"},
		Context:   ctx,
	})
	require.NoError(b.T(), err)
	b.client = client //初始化客户端
}
func (b *BalancerTestSuite) TestCustomRoundRobinClient() {
	//通过服务注册得到客户端
	db, err := resolver.NewBuilder(b.client)
	require.NoError(b.T(), err)
	//使用自定义的负载策略，"custom_wrr": {这里其实可以写配置文件}
	svcCfg := `{
			"loadBalancingConfig": [{
				"custom_wrr": {}
			}]
		}`
	clientConn, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(db),
		grpc.WithDefaultServiceConfig(svcCfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(b.T(), err)
	client := webookgrpc.NewUserServiceClient(clientConn)
	timeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for i := 0; i < 10; i++ {
		login, err := client.Login(timeout, &webookgrpc.UserRequest{
			Id: 432,
		})
		require.NoError(b.T(), err)
		fmt.Println(login)
	}
}

func (b *BalancerTestSuite) TestPickFirst() {
	//通过服务注册得到客户端
	db, err := resolver.NewBuilder(b.client)
	require.NoError(b.T(), err)
	//使用"weighted_round_robin"负载策略
	svcCfg := `{
			"loadBalancingConfig": [{
				"weighted_round_robin": {}
			}]
		}`
	clientConn, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(db),
		grpc.WithDefaultServiceConfig(svcCfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(b.T(), err)
	client := webookgrpc.NewUserServiceClient(clientConn)
	timeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for i := 0; i < 10; i++ {
		login, err := client.Login(timeout, &webookgrpc.UserRequest{
			Id: 432,
		})
		require.NoError(b.T(), err)
		fmt.Println(login)
	}

}
func (b *BalancerTestSuite) startServer(addr string, weight float64) {
	listen, err := net.Listen("tcp", addr)
	require.NoError(b.T(), err)
	em, err := endpoints.NewManager(b.client, "service/user")
	require.NoError(b.T(), err)
	addr = "127.0.0.1" + addr
	key := "service/user/" + addr
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Second)
	defer cancel()
	//注册端点
	err = em.AddEndpoint(ctx, key, endpoints.Endpoint{
		Addr:     addr,
		Metadata: `{"weight":` + strconv.Itoa(int(weight)) + `}`,
	})
	require.NoError(b.T(), err)
	server := grpc.NewServer() //创建服务器，用于注册信息
	userService := NewUserServiceRPC(addr)
	webookgrpc.RegisterUserServiceServer(server, userService)
	fmt.Println("启动成功了" + addr)
	err = server.Serve(listen)
	require.NoError(b.T(), err)
}
func (b *BalancerTestSuite) TestServer() {
	go func() {
		b.startServer(":8090", 90)
	}()
	b.startServer(":8091", 80)
}
func TestLoadBalance(t *testing.T) {
	suite.Run(t, new(BalancerTestSuite))
}
