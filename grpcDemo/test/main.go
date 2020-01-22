package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	hello "practicProject/grpcDemo/grpcDemo/proto"
)

func main() {
	text := &hello.HelloRequest{
		Name: "Mike",
		Age: 23,
		Weight: []int32{120, 119, 111, 128},
	}
	fmt.Println(text)

	//编码
	data, err := proto.Marshal(text)
	if err != nil {
		fmt.Println("编码失败")
	}
	// 输出编码后的数据（二进制）
	fmt.Println(data)

	newText := &hello.HelloRequest{}
	// proto 解码
	err = proto.Unmarshal(data, newText)
	if err != nil {
		fmt.Println("解码失败")
	}
	// 输出解码后的数据
	fmt.Println(newText)

	fmt.Println("name:", newText.Name)
	fmt.Println("age:", newText.Age)
	fmt.Println("weight:", newText.Weight)
}
