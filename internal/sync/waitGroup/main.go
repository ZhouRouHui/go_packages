package main

import (
	"fmt"
	"sync"
)

/**
WaitGroup

WaitGroup 对象内部有一个计数器，最初从 0 开始
他有三个方法 Add(), Done(), Wait() 用来控制计数器的数量
Add 往计数器里加值，Done 把计数器减 1，Wait 会阻塞代码的运行，直到计数器为 0 为止
*/
var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Printf("i = %d\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
