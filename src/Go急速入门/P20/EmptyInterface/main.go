package main

import "fmt"

// 空接口
// 没有定义任何方法时
// 任意类型都实现了空接口->空接口可以存储任意值

// 一般不需要提前定义
type xxx interface {
}

// 应用
// 1、作为函数参数
// 2、可以作为map的value
func main() {
	var x interface{}
	x = "hello world"
	fmt.Println(x)
	var m = make(map[string]interface{}, 16)
	m["name"] = "na"
	m["age"] = 18
	m["hobby"] = []string{"base", "bool", "double"}
	fmt.Println(m)

	// 类型断言
	ret, ok := x.(string)
	if !ok {
		fmt.Println("不是bool")
	} else {
		fmt.Println(ret)
	}
	switch t := x.(type) {
	case string:
		fmt.Println("is string", t)
	case bool:
		fmt.Println("is bool", t)
	case int:
		fmt.Println("is int", t)
	default:
		fmt.Println("猜不到")
	}
}
