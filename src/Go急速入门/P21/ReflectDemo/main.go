package main

import (
	"fmt"
	"reflect"
)

// ReflectType 获得x的类型
func ReflectType(x interface{}) {
	// 1、通过类型断言
	// 需要一个一个猜
	// 2、反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Kind())
	fmt.Printf("%T\n", obj)
}

type Cat struct {
}

type Dog struct {
}

func ReflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v %T\n", v, v)
	k := v.Kind() // 拿到值对应的类型种类
	fmt.Println(k)
	// 如何得到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		// 把反射取到的值转换成一个int32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v, %T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v, %T\n", ret, ret)
	}
}

func ReflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// Elem用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.14)
	}
}

func main() {
	//var a float32 = 3.14
	//ReflectType(a)
	//var b uint8 = 10
	//ReflectType(b)
	//// 结构体类型
	//var c Cat
	//ReflectType(c)
	//var d Dog
	//ReflectType(d)
	//// slice
	//var e []int
	//ReflectType(e)
	//var f []string
	//ReflectType(f)

	//var aa int32 = 1035
	//ReflectValue(aa)
	//var bb float32 = 3.1415926
	//ReflectValue(bb)

	var aaa int32 = 10
	ReflectSetValue(&aaa)
	fmt.Println(aaa)
}
