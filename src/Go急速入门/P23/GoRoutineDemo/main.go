package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func hello(i int) {
	fmt.Println("Hello", i)
	wg.Done() // 通知wg把计数器减一
}

func main() { //开启了一个主goroutine执行hello这个函数
	for i := 0; i < 10000; i++ {
		wg.Add(1)   // 计数器+1
		go hello(i) // 开启了一个goroutine执行hello这个函数
	}

	fmt.Println("World")
	// time.Sleep(time.Second)
	wg.Wait()

}
