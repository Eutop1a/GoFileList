package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() { //开启了一个主goroutine执行hello这个函数
	for i := 0; i < 100; i++ {
		wg.Add(1)        // 计数器+1
		go func(i int) { // 开启了一个goroutine执行hello这个函数
			fmt.Println("Hello", i)
			wg.Done()
		}(i)
	}

	fmt.Println("World")
	wg.Wait()

}
