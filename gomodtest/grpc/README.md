## gRPC 流模式

![](https://zkept-1302605083.cos.ap-nanjing.myqcloud.com/md/gRPC_server_stream.png)



![](https://zkept-1302605083.cos.ap-nanjing.myqcloud.com/md/gRPC_client_stream.png)



![](https://zkept-1302605083.cos.ap-nanjing.myqcloud.com/md/gRPC_twoWay_stream.png)



### 1. 服务端流模式

#### 服务端

**Protobuf** 👇

```protobuf
rpc GetUserScoreByServerStream(UserScoreRequest) returns (stream UserScoreResponse);
```



**Interface**👇

接收两个参数，第一个参数是在 proto 文件中自定义的，第二个是 gRPC 自动生成的，用于返回数据

```go
GetUserScoreByServerStream(*UserScoreRequest, UserScoreService_GetUserScoreByServerStreamServer) error
```



**interface impl** 👇

```go
func (u *UserService) f(req *pb.UserScoreRequest,
   stream pb.UserScoreService_GetUserScoreByServerStreamServer) error
```



服务端通过 **Send 方法**，向客户端发送数据

**Send方法** 定义如下👇

```go
Send(*UserScoreResponse) error
```



使用👇

```go
rsp := new(pb.UserScoreResponse)
err := stream.Send(rsp)
```



#### 客户端

客户端使用  **Recv 方法** 接收服务端返回的数据

**Recv 方法** 定义如下

```go
Recv() (*UserScoreResponse, error)
```



使用👇

```go
stream, err := client.GetUserScoreByServerStream(context.Background(), req)
for {
		recv, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(recv.UserScore)
	}
```