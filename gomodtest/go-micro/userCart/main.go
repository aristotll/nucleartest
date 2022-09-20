package main

import (
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"go-micro/userCart/router"
)

func main() {
	r := router.GetRouter()
	// 获取在 127.0.0.1:2379 的 etcd 注册中心
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	// 使用 go-micro 创建一个服务
	service := web.NewService(web.Address(":8090"), // 端口号 8080
		web.Handler(r),        //  gin 的处理器
		web.Name("user"),      //  服务名
		web.Registry(etcdReg)) //   注册到的服务中心

	err := service.Run()
	if err != nil {
		fmt.Println(err)
	}
}
