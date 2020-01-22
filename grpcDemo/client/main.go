package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"practicProject/grpcDemo/grpcDemo/config"
	hello "practicProject/grpcDemo/grpcDemo/proto"
)

func main() {
	// 客户带连接服务器
	conn, err := grpc.Dial(config.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端（获得client句柄）
	client := hello.NewHelloClient(conn)

	// 实例对象
	reqBody := &hello.HelloRequest{}
	reqBody.Name = "Jack"
	reqBody.Age = 23
	reqBody.Weight = []int32{120, 116, 126}
	// 通过句柄调用函数
	resp, err := client.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
}