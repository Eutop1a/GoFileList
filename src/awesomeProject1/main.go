package main

import (
	"awesomeProject1/calc"
	"fmt"

	"rsc.io/quote"
)

//
//type Person interface {
//	getName() string
//}
//
//type Student struct {
//	name string
//	age  int
//}
//
//func (stu *Student) getName() string {
//	return stu.name
//}
//
//type Worker struct {
//	name   string
//	gender int
//}
//
//func (w *Worker) getName() string {
//	return w.name
//}
//
//var wg sync.WaitGroup
//
//func download(url string) {
//	fmt.Println("start to download", url)
//	time.Sleep(time.Second)
//	wg.Done()
//}
//
//var ch = make(chan string, 10)
//
//func download(url string) {
//	fmt.Println("start to download", url)
//	time.Sleep(time.Second)
//	ch <- url
//}

func main() {
	//var p Person = &Student{
	//	name: "Tom",
	//	age:  18,
	//}
	//
	//fmt.Println(p.getName())
	//for i := 0; i < 3; i++ {
	//	wg.Add(1)
	//	go download("a.com/" + string(i+'0'))
	//}
	//wg.Wait()
	//fmt.Println("Done!")
	//
	//for i := 0; i < 3; i++ {
	//	go download("a.com/" + string(i+'0'))
	//}
	//for i := 0; i < 3; i++ {
	//	msg := <-ch
	//	fmt.Println("finish", msg)
	//}
	//fmt.Println("done")
	fmt.Println(quote.Hello())
	fmt.Println(calc.Add(10, 3))

}
