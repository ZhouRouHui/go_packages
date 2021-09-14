package main

import (
	"fmt"
	"sync"
)

/**
Once
作用：sync.Once 类型的变量只会执行一次 Do 方法，无论传给 Do 方法的参数里面做什么，
	常用的场景如数据初始化或者业务中捕获第一次出现的错误
*/

var once sync.Once

func onces() {
	fmt.Println("onces")
}

func onced() {
	fmt.Println("onced")
}

func main() {
	for i, v := range make([]int, 10) {
		once.Do(onces) // 只会在 i=0 的时候执行一次
		fmt.Println("count: ", v, " --- ", i)
	}

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onced) // 上面执行过了，这里就不会再执行了
			fmt.Println(123)
		}()
	}
}
