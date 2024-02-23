package wrr

import (
	_ "container/heap"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

const name = "custom_wrr"

func init() {
	//注册自定义
	balancer.Register(
		base.NewBalancerBuilder(name,
			&PickerBuilder{},
			//config是是否做监控检测
			base.Config{HealthCheck: true}))
}

// PickerBuilder 自定义负载均衡算法
type PickerBuilder struct {
}

// Build 选中策略
func (p *PickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	//去构造节点
	nodes := make([]*node, len(info.ReadySCs))
	for conn, connInfo := range info.ReadySCs {
		nd := &node{
			cc: conn,
			//weight: ,
		}

		md, ok := connInfo.Address.Metadata.(map[string]any)
		if ok {
			weightVal := md["weight"]
			weight, _ := weightVal.(int)
			nd.weight = weight
		}
		if nd.weight == 0 {
			nd.weight = 80
		}
		nd.currentWeight = nd.weight //初始权重
		nodes = append(nodes, nd)
	}
	return &Picker{
		nodes: nodes,
	}
}

type Picker struct {
	//真正执行负载均衡
	nodes []*node
}

// 实现真正的负载均衡算法
func (p *Picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	//TODO implement me
	var total int
	for _, n := range p.nodes {
		total += n.weight //总权重
		//计算节点当前权重
		n.currentWeight = n.currentWeight + n.weight
	}
	//感觉节点可以选用最大堆存储，然后直接返回最大堆的节点

	return balancer.PickResult{
		SubConn: p.nodes[0].cc,
		Done: func(info balancer.DoneInfo) {
			//完成后可执行的方法，调整权重
		},
	}, nil
}

// 节点
type node struct {
	//heap.HeapNode                  //本来组合一下我的堆节点，
	//方便使用自定义的大顶堆，结果发现不咋好弄
	weight        int              //权重
	currentWeight int              //当前权重
	cc            balancer.SubConn //真正的节点
}
