package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	myhttp "github.com/micro/go-plugins/client/http"
	"go-micro/userInfo/model"
	"io/ioutil"
	"log"
	"net/http"
)

// 原生的方法请求
// addr, path, method string
// http://localhost:8080/call/basic?method=GET&addr=localhost:8090&path=/user/cart
func BasicCall(c *gin.Context)  {
	method := c.Query("method")
	addr := c.Query("addr")
	path := c.Query("path")
	log.Println("http://" + addr + path)
	request, err := http.NewRequest(method, "http://" + addr + path, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.DefaultClient

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	readAll, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	rsp := string(readAll)
	log.Printf(rsp)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": rsp,
	})
}

// 该方法有时成功有时失败，还未找到原因
// 报错: invalid character 'p' after top-level value
func call2(s selector.Selector) {
	newClient := myhttp.NewClient(client.Selector(s),
		client.ContentType("application/json"))
	// 请求
	// warning: 该方法默认以 post 方法请求，对应的请求地址需要改为 post
	// 模拟发送携带请求参数的请求
	req := newClient.
		NewRequest("user", "/user/cart/json",
			map[string]interface{}{"Id": 11, "productName": "iphone1"})
	log.Println("this is req!!!!", req)

	// 用 map 接收响应数据
	var rsp map[string]interface{}
	err := newClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println(rsp)
}

func call3(s selector.Selector) {
	newClient := myhttp.NewClient(client.Selector(s),
		client.ContentType("application/json"))
	// 请求
	// warning: 该方法默认以 post 方法请求，对应的请求地址需要改为 post
	// 模拟发送携带请求参数的请求
	req := newClient.
		NewRequest("user", "/user/cart/json",
			model.CartRequest{
				Id:          1111,
				ProductName: "iphone312312",
			})
	log.Println("this is req!!!!", req)


	var rsp model.CartResponse
	err := newClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println(rsp)
}

func CallByPlugins(c *gin.Context) {
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	// 快速创建一个选择器，并制定注册中心和选择器方式
	newSelector := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin))
	call2(newSelector)
}

func CallBygRPC(c *gin.Context) {
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	// 快速创建一个选择器，并制定注册中心和选择器方式
	newSelector := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin))
	call3(newSelector)
}