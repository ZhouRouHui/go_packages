package main

import (
	"fmt"
	"os"
	"os/signal"
)

/**
os/signal 包实现对信用的处理

golang 中对信号的处理主要使用 os/signal 包中的两个方法
	Notify() 监听收到的信号
	Stop() 	取消监听
*/

// Notify 方法使用
func testNotify() {
	c := make(chan os.Signal, 0)
	signal.Notify(c)

	s := <-c
	fmt.Println("Got signal: ", s)

	/**
	编译并执行后，通过 ps aux | grep ./main 查看进程号
	kill 进程号
	便会得到打印结果：Got signal:  terminated
	*/
}

func testStop() {
	c := make(chan os.Signal, 0)
	signal.Notify(c)

	signal.Stop(c) // 不允许在 c 中存入内容

	s := <-c // c 无内容，此次阻塞，所以不会执行下面的语句，也就没有输出
	fmt.Println("Got signal: ", s)
}

func main() {
	//testNotify()

	//testStop()
}
