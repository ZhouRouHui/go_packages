package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type World struct {
}

// HelloWorld 供远程调用的方法
// 几个要求：
// 1）方法必须是导出的 —— 包外可见。 首字母大写。
// 2）方法必须有两个参数， 都是导出类型、內建类型。
// 3）方法的第二个参数必须是 “指针” （传出参数）
// 4）方法只有一个 error 接口类型的 返回值。
func (w *World) HelloWorld(name string, resp *string) error {
	*resp = "hello " + name
	return nil
}

func main() {
	// 1. 注册 rpc 服务，绑定对象方法
	// 		参1：服务名。字符串类型。
	//     	参2：对应 rpc 对象。 该对象绑定方法要满足如下条件：
	//         1）方法必须是导出的 —— 包外可见。 首字母大写。
	//         2）方法必须有两个参数， 都是导出类型、內建类型。
	//         3）方法的第二个参数必须是 “指针” （传出参数）
	//         4）方法只有一个 error 接口类型的 返回值。
	err := rpc.RegisterName("hello_service", new(World))
	if err != nil {
		panic(err)
	}

	// 2. 设置监听
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("开始监听")

	// 3. 建立链接
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("链接成功")

	// 4. 使用 jsonrpc 绑定服务
	jsonrpc.ServeConn(conn)
}
