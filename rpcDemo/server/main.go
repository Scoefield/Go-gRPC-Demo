package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func helloText(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "http request golang rpc test")
	if err != nil {
		fmt.Println("request error", err)
	}
}

/*
- 方法是导出的
- 方法有两个参数，都是导出类型或内建类型
- 方法的第二个参数是指针
- 方法只有一个error接口类型的返回值
 */
type HelloType int
func (h *HelloType)GetHello (argType int, replyType *int) error {
	fmt.Println("print client send content:", argType)

	// 修改返回内容值
	*replyType = argType + 12306
	return nil
}

type Student struct {
	Name string
}

func (s *Student)GetStuScore (args *Student, reply *float64) error{
	if args.Name == "Jack" {
		*reply = 98.6
	} else if args.Name == "Mike" {
		*reply = 89.9
	} else {
		*reply = 60.0
	}
	return nil
}

func main() {
	http.HandleFunc("/hello", helloText)

	// 创建一个对象
	//hello := new(HelloType)
	stu := Student{}
	// 服务端注册一个对象
	//err := rpc.Register(hello)
	err := rpc.Register(&stu)
	if err != nil {
		fmt.Println("register server error:", err)
	}
	// 连接网络
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":10086")
	if err != nil{
		fmt.Println("network error:", err)
	}
	fmt.Println("http server start listening...")
	_ = http.Serve(listen, nil)
}
