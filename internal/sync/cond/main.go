package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func main() {
	for i := 0; i < 40; i++ {
		go func(x int) {
			cond.L.Lock()         // 获取锁
			defer cond.L.Unlock() // 释放锁
			cond.Wait()           // 等待通知，阻塞当前 goroutine
			fmt.Println(x)
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.Signal() // 下发一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second)
	cond.Signal() // 下发一个通知给已经获取锁的 goroutine
	time.Sleep(3 * time.Second)
	cond.Broadcast() // 下发广播给所有等待的 goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 60)
}
