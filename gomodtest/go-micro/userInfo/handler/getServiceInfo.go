package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"net/http"
)

// 随机获取一个注册中心中的服务信息
// 随机方式很不稳定，可能造成某一节点被频繁访问，导致压力大
// 或者某一节点很少被访问，压力小
func GetServiceByRandom(c *gin.Context) {
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	// 参数为服务名
	services, err := etcdReg.GetService("user")
	if err != nil {
		fmt.Println(err)
	}
	// 随机获取
	next := selector.Random(services)
	node, err := next()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"address": node.Address,
		"Id": node.Id,
		"metadata": node.Metadata,
	})
	fmt.Printf("address: %s  Id: %s  metadata: %s",
		node.Address, node.Id, node.Metadata)
}

// 轮询的方式获取服务
// 以相对平均的频率访问各个节点
func GetServiceByPoll(c *gin.Context) {
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	// 参数为服务名
	services, err := etcdReg.GetService("user")
	if err != nil {
		fmt.Println(err)
	}

	// 该 map 用于统计每个节点被访问的次数
	m := make(map[string]int)
	// 8080 出现的次数
	times8080 := 0
	times8090 := 0
	times8091 := 0
	for i := 0; i < 100; i++ {
		// 随机获取
		next := selector.RoundRobin(services)
		node, err := next()
		if err != nil {
			fmt.Println(err)
		}

		address := node.Address
		// 截取字符串，获取端口号
		port := address[14:]

		switch port {
		case "8080":
			times8080++
			m["8080"] = times8080
		case "8090":
			times8090++
			m["8090"] = times8090
		case "8091":
			times8091++
			m["8091"] = times8091
		}

		fmt.Println(m)

		//fmt.Printf("address: %s  Id: %s  metadata: %s \n",
		//	node.Address, node.Id, node.Metadata)
		//time.Sleep(time.Second * 2)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "请看控制台",
		"8080 被访问的次数 ": m["8080"],
		"8090 被访问的次数 ": m["8090"],
		"8091 被访问的次数 ": m["8091"],
	})

}
