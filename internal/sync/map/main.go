package main

import (
	"fmt"
	"sync"
)

/**
sync.Map 线程安全的 map，go 语言原生的 map 是非线程安全的.
原生 map 在并发读写时，一般做法是加锁，但是加锁性能不高。

go 1.9版本中提供了 sync.Map，效率高，线程安全。
特性：
	* 无需初始化，直接申明即可
	* 不能像原生 map 一样进行取值和设置操作，而是使用该结构体的方法，Store 表示存储，Load 表示获取，Delete 表示删除
	* 使用结构体的 Range 方法配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续遍历时返回 true，终止迭代时返回 false。
*/

var m map[int]int = make(map[int]int, 10)

// go 原生 map 非线程安全示例
func goMapUnsafe() {
	go func() {
		for {
			m[0] = 199
		}
	}()

	go func() {
		for {
			_ = m[0]
		}
	}()

	for {
	}
}

func useSyncMap() {
	var scene sync.Map

	// 将键值对保存至 scene 中
	scene.Store("greece", 17)
	scene.Store("london", 100)
	scene.Store("eqypt", 200)

	// 从 sync.Map 中取值
	eqypt, ok := scene.Load("eqypt")
	if !ok {
		panic("load syncMap err")
	}
	fmt.Printf("scene 中 eqypt 的值为：%v\n", eqypt)

	scene.Delete("london")

	// 获取并删除
	greece, loaded := scene.LoadAndDelete("greece")
	fmt.Printf("greece = %v, loaded = %t\n", greece, loaded)

	// 获取或者存储
	loedan, loaded := scene.LoadOrStore("loedan", "zrh")
	fmt.Printf("loedan = %v, loaded = %t\n", loedan, loaded)

	// 遍历 sync.Map 里面的内容
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate: ", k, v)
		return true
	})
}

func main() {
	// go 原生 map 非线程安全示例
	//goMapUnsafe()

	useSyncMap()
}
