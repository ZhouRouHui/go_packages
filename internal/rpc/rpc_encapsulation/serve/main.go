package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type MyInterface interface {
	HelloWorld(string, *string) error
}

// RegisterService 注册服务
func RegisterService(name string, i MyInterface) (err error) {
	err = rpc.RegisterName(name, i)
	return
}

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
	//return nil
	return errors.New("test error")
}

func main() {
	err := RegisterService("hello_service", new(World))
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

	// 4. 绑定服务
	rpc.ServeConn(conn)
}
