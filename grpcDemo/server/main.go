package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"practicProject/grpcDemo/grpcDemo/config"
	hello "practicProject/grpcDemo/grpcDemo/proto"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService ...
var HelloService = helloService{}

// 实现SayHello服务
func (h helloService) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	//resp := new(hello.HelloResponse)
	resp := &hello.HelloResponse{}
	resp.Msg = "request success"
	resp.Code = "200"
	resp.Data = fmt.Sprintf("{name: %s, age: %d, weight: %d}", in.Name, in.Age, in.Weight)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	src := grpc.NewServer()

	// 注册HelloService
	hello.RegisterHelloServer(src, &HelloService)

	fmt.Println("Listen on " + config.Address)
	// 等待网络连接
	src.Serve(listen)
}