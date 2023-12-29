package main

import "fmt"

const (
	size = 10
)

func main() {
	a := new(int)
	b := make([]int, size)
	fmt.Printf("%#v\n, %#v\n", a, b)
}
