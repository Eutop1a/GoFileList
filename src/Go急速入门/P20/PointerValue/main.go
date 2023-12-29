package main

import "fmt"

// 使用值接收和指针接收的区别

type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

func (p *person) say() {
	fmt.Printf("%s狗叫\n", p.name)
}

// 使用值接收实现接口，类型的值和类型的指针都能保存到接口变量中
func (p person) move() {
	fmt.Printf("%s跑了\n", p.name)
}

// 使用指针接收实现接口，只有类型的指针能保存到接口变量中

//func (p *person) move() {
//	fmt.Printf("%s跑了\n", p.name)
//}

func main() {
	var m mover
	p1 := person{
		name: "nihao",
		age:  18,
	}
	p2 := &person{ // p2 是person类型的指针
		name: "no",
		age:  18,
	}
	m = p1 // 无法赋值，因为p1是person类型的值，没有实现mover接口
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer
	s = p2
	s.say()

	var a animal
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}
