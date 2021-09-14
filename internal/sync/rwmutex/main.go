package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/**
读写锁
性能高于互斥锁，因为互斥锁对于读的操作每一次都需要执行加锁解锁操作，
而读写锁使用读锁时，一个协程获取到读锁时其他协程也可以读取内容。

读写锁在读和写的不同操作时有不同的锁，写用写锁，读用读锁
*/

// 读写锁实例
var rwLock sync.RWMutex
var m map[int]int = map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}

func useRWMutex() {
	// 开启两个协程对 m 进行写操作
	for i := 0; i < 2; i++ {
		go func() {
			rand.Seed(time.Now().UnixNano())
			// 写数据写用写锁
			rwLock.Lock()
			m[0] = rand.Intn(100)
			rwLock.Unlock()
		}()
	}

	// 读数据用读锁
	rwLock.RLock()
	fmt.Println(m)
	rwLock.RUnlock()
}

func testRWMutexPerformance() {
	var count int32
	for i := 0; i < 100; i++ {
		go func() {
			for {
				rwLock.RLock()
				// 假设每次读取数据花费 1 毫秒
				time.Sleep(time.Millisecond)
				rwLock.RUnlock()

				atomic.AddInt32(&count, 1)
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(atomic.LoadInt32(&count)) // 241258，性能高出互斥锁 100 倍
}

func main() {
	// 读写锁的使用
	useRWMutex()
	fmt.Println()

	testRWMutexPerformance()
}
