package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	// 1. 用 jsonrpc 链接服务器
	conn, err := jsonrpc.Dial("tcp", ":9001")
	// 用 rpc 的方式去拨号创建链接，conn 就是一个 socket 套接字
	if err != nil {
		panic(err)
	}

	// 2. 调用远程函数
	var resp string
	// conn.Call("服务名.方法名", 传入参数, 传出参数)，传出参数就是远程函数返回的东西
	err = conn.Call("hello_service.HelloWorld", "zrh", &resp)
	if err != nil {
		panic(err)
	}

	fmt.Println("调用成功")
	fmt.Println(resp)
}
