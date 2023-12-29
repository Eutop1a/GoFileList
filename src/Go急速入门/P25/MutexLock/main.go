package main

import (
	"fmt"
	"sync"
)

// 多个goroutine并发操作x

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 10000; i++ {
		lock.Lock() // 加锁
		x++
		lock.Unlock() // 释放锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
