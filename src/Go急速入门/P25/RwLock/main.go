package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁耗时     1.1572986s
// 读写互斥锁耗时  14.4021ms

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func read() {
	//lock.Lock()
	rwlock.RLock()
	time.Sleep(time.Microsecond)
	//lock.Unlock()
	rwlock.RUnlock()
	wg.Done()
}

func write() {
	rwlock.Lock()
	//lock.Lock()
	x++
	time.Sleep(time.Microsecond * 10)
	//lock.Unlock()
	rwlock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
