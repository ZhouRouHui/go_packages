package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/**
互斥锁
性能低于读写锁，因为互斥锁对于读的操作每一次都需要执行加锁解锁操作，
而读写锁使用读锁时，一个协程获取到读锁时其他协程也可以读取内容。
*/

// 互斥锁实例
var lock sync.Mutex
var m map[int]int = map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}

func useMutex() {
	// 开启两个协程对 m 进行写操作
	for i := 0; i < 2; i++ {
		go func() {
			rand.Seed(time.Now().UnixNano())
			lock.Lock()
			m[0] = rand.Intn(100)
			lock.Unlock()
		}()
	}

	// 当前协程进行读数据也需要加锁
	lock.Lock()
	fmt.Println(m)
	lock.Unlock()
}

// 开启 100 个协程，每个协程里面 for 循环进行加解锁，统计次数
// atomic 进行原子操作，保证线程安全，不会出现数据竞争，这里是为了测试方便
func testMutexPerformance() {
	var count int32
	for i := 0; i < 100; i++ {
		go func() {
			for {
				lock.Lock()
				// 假设每次读取数据花费 1 毫秒
				time.Sleep(time.Millisecond)
				lock.Unlock()

				atomic.AddInt32(&count, 1)
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(atomic.LoadInt32(&count)) // 2316，性能低出读写锁 100 倍
}

func main() {
	// 互斥锁的使用
	useMutex()
	fmt.Println()

	// 测试互斥锁的性能
	testMutexPerformance()
}
