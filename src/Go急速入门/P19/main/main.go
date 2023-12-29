package main

// 当代码都放在&GOPATH目录下的话
// 导入包的时候要从&GOPATH/src开始写起

// 可以起别名：类似于 aaa "P19/calc"

import (
	"Package/calc"
	"fmt"
)

func init() {
	fmt.Println("main init()")
}

func main() {
	fmt.Println("hello world")
	ret := calc.Add(10, 20)
	fmt.Println(calc.Name)
	fmt.Println(ret)
}
