package main

import "fmt"

// 接口:是一种类型，是一种抽象的类型

// 为什么需要接口

type dog struct{}

func (d dog) say() {
	fmt.Println("汪")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("苗")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("啊!")
}

// 接口不管是什么类型，只管你需要实现什么方法
// 定义一个抽象类型，只要实现了say这个方法的类型都可以成为sayer类型
type sayer interface {
	say()
}

func fight(arg sayer) {
	arg.say() // 传进来是啥都得打，然后叫
}

func main() {
	//c1 := cat{}
	//fight(c1)
	//d1 := dog{}
	//fight(d1)
	//p1 := person{
	//	name: "你好",
	//}
	//fight(p1)
	var s sayer
	c2 := cat{}
	s = c2
	p2 := person{name: "nihao "}
	s = p2
	fmt.Println(s)
}
