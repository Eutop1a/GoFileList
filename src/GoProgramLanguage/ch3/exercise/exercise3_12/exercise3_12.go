package main

import (
	"fmt"
)

func anagrams(a, b string) bool {
	map1, map2 := map[string]int{}, map[string]int{}
	for _, v := range a {
		map1[string(v)]++
	}

	for _, v := range b {
		map2[string(v)]++
	}
	for s, i := range map1 {
		if map2[s] != i {
			return false
		}
	}
	return true
	//return reflect.DeepEqual(map1, map2)两种方法都行
}

func main() {
	a, b := "Hello", "olleH"
	fmt.Println(anagrams(a, b))
}
