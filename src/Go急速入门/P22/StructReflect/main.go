package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score"ini:"s_score"`
}

func (s student) Study() string {
	msg := "学习"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "睡觉"
	fmt.Println(msg)
	return msg
}

// 根据反射获取结构体中的方法
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	// 因为下面需要拿到具体的方法调用，所以使用的是值信息
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{} // 指定参数
		v.Method(i).Call(args)       // 调用方法
	}
	// 通过方法名获取结构体的方法
	t.MethodByName("Sleep")              // (Method, bool)
	methodObj := v.MethodByName("Sleep") // Value
	fmt.Println(methodObj.Type())
}

func main() {
	stu1 := student{
		Name:  "xiaowangzi",
		Score: 90,
	}
	//// 通过反射获取结构体中的所有字段
	//t := reflect.TypeOf(stu1)
	//fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
	//// 遍历结构体变量的所有字段
	//for i := 0; i < t.NumField(); i++ {
	//	// i就是结构体字段的索引，根据i取字段
	//	filedObj := t.Field(i)
	//	fmt.Printf("name:%v \t type:%v \t tag:%v",
	//		filedObj.Name, filedObj.Type, filedObj.Tag)
	//	fmt.Println(filedObj.Tag.Get("json"),
	//		filedObj.Tag.Get("ini"))
	//}
	//// 根据名字取结构体中的字段
	//filedObj2, ok := t.FieldByName("Score")
	//if ok {
	//	fmt.Printf("name:%v \t type:%v \t tag:%v",
	//		filedObj2.Name, filedObj2.Type, filedObj2.Tag)
	//}
	printMethod(stu1)
}
