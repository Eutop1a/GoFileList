package main

import "fmt"

func test(arr ...string) string {
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
	return ""
}

func main() {
	testArr := []string{"1234", "我测你的吗"}
	test(testArr...)
}
