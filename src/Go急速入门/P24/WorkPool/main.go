package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		result <- job * 2
		time.Sleep(time.Microsecond * 500)
		fmt.Printf("worker:%d stop job:%d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	// 开启3个goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, result)
	}
	// 发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	//输出结果
	for i := 0; i < 5; i++ {
		ret := <-result
		fmt.Println(ret)
	}
}
