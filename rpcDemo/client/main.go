package main

import (
	"fmt"
	"net/rpc"
)

type Student struct {
	Name string
}

func main() {
	// 建立网络连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("client rpc dial http error:", err)
	}

	//var HelloType int
	//err = client.Call("HelloType.GetHello", 2020, &HelloType)
	args := &Student{"Mike"}
	var score float64
	err = client.Call("Student.GetStuScore", args, &score)
	if err != nil {
		fmt.Println("client call error:", err)
	}

	//fmt.Println("response result:", HelloType)
	fmt.Println("response result:", score)
}
