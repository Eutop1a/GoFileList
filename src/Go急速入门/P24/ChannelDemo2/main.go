package main

import "fmt"

/*
两个goroutine
	1、生成0~100的数字，发送到ch1
	2、从ch1中取出数据计算它的平方，把结果发送到ch2
*/

// ch只能写入
func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

// ch1只能发送，ch2只能写入
func f2(ch1 <-chan int, ch2 chan<- int) {
	// 从通道中取值的方式1
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)

}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)
	go f1(ch1)
	go f2(ch1, ch2)
	// 从通道中取值的方式2
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
